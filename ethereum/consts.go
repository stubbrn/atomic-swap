package contracts

// Gas prices for our operations. Most of these are set by the highest value we
// ever see in a test, so you would need to adjust upwards a little to use as a
// gas limit. We use these values to estimate minimum required balances.
const (
	MaxNewSwapETHGas   = 50639
	MaxNewSwapTokenGas = 86218
	MaxSetReadyGas     = 32054
	MaxClaimETHGas     = 43349
	MaxClaimTokenGas   = 47522
	MaxRefundETHGas    = 43132
	MaxRefundTokenGas  = 47294
	MaxTokenApproveGas = 47000 // 46223 with our contract
)

// constants that are interesting to track, but not used by swaps
const (
	maxSwapCreatorDeployGas = 1094089
	maxTestERC20DeployGas   = 798286 // using long token names or symbols will increase this
)