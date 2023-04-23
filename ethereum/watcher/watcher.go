// Copyright 2023 The AthanorLabs/atomic-swap Authors
// SPDX-License-Identifier: LGPL-3.0-only

// Package watcher provides tools to track events emitted from ethereum contracts.
package watcher

import (
	"context"
	"errors"
	"math/big"
	"time"

	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	logging "github.com/ipfs/go-log"
)

const (
	checkForBlocksTimeout = time.Second
)

var (
	log = logging.Logger("ethereum/watcher")
)

// EventFilter filters the chain for specific events (logs).
// When it finds a desired log, it puts it into its outbound channel.
type EventFilter struct {
	ctx         context.Context
	cancel      context.CancelFunc
	ec          *ethclient.Client
	topic       ethcommon.Hash
	filterQuery eth.FilterQuery
	logCh       chan<- ethtypes.Log
}

// NewEventFilter returns a new *EventFilter.
func NewEventFilter(
	ctx context.Context,
	ec *ethclient.Client,
	contract ethcommon.Address,
	fromBlock *big.Int,
	topic ethcommon.Hash,
	logCh chan<- ethtypes.Log,
) *EventFilter {
	filterQuery := eth.FilterQuery{
		FromBlock: fromBlock,
		Addresses: []ethcommon.Address{contract},
	}

	ctx, cancel := context.WithCancel(ctx)
	return &EventFilter{
		ctx:         ctx,
		cancel:      cancel,
		ec:          ec,
		topic:       topic,
		filterQuery: filterQuery,
		logCh:       logCh,
	}
}

// Start starts the EventFilter. It watches the chain for logs.
func (f *EventFilter) Start() error {
	go func() {
		for {
			select {
			case <-f.ctx.Done():
				return
			case <-time.After(checkForBlocksTimeout):
			}

			currHeader, err := f.ec.HeaderByNumber(f.ctx, nil)
			if err != nil {
				log.Errorf("failed to get header in event watcher: %s", err)
				if errors.Is(err, ethrpc.ErrClientQuit) {
					return // non-recoverable error
				}
				continue
			}

			if currHeader.Number.Cmp(f.filterQuery.FromBlock) <= 0 {
				// no new blocks, don't do anything
				continue
			}

			// let's see if we have logs
			logs, err := f.ec.FilterLogs(f.ctx, f.filterQuery)
			if err != nil {
				log.Errorf("failed to filter logs for topic %s: %s", f.topic, err)
				continue
			}

			// If you think we are missing log events, uncomment to debug:
			// log.Debugf("filtered for logs from block %s to block %s",
			// 	f.filterQuery.FromBlock, currHeader.Number)

			for _, l := range logs {
				if l.Topics[0] != f.topic {
					continue
				}

				if l.Removed {
					log.Debugf("found removed log: tx hash %s", l.TxHash)
					continue
				}

				log.Debugf("watcher for topic %s found log in block %d", f.topic, l.BlockNumber)
				f.logCh <- l
			}

			// the filter query is inclusive of the `ToBlock`, so we already checked `ToBlock`
			// and we don't want to check it again, so increment by 1
			f.filterQuery.FromBlock = big.NewInt(0).Add(currHeader.Number, big.NewInt(1))
		}
	}()

	return nil
}

// Stop stops the EventFilter.
func (f *EventFilter) Stop() {
	f.cancel()
}
