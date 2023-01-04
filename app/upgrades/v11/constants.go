package v11

const (
	// UpgradeName is the shared upgrade plan name for mainnet
	UpgradeName = "v11.0.0"
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/evmos/evmos/releases/download/v11.0.0/evmos_11.0.0_Darwin_arm64.tar.gz","darwin/amd64":"https://github.com/evmos/evmos/releases/download/v11.0.0/evmos_11.0.0_Darwin_amd64.tar.gz","linux/arm64":"https://github.com/evmos/evmos/releases/download/v11.0.0/evmos_11.0.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/evmos/evmos/releases/download/v11.0.0/evmos_11.0.0_Linux_amd64.tar.gz","windows/x86_64":"https://github.com/evmos/evmos/releases/download/v11.0.0/evmos_11.0.0_Windows_x86_64.zip"}}'`

	// at the time of this migration, on mainnet, channels 0 to 36 were open
	// so this migration covers those channels only
	openChannels = 36

	// FundingAccount is the account which holds the rewards for the incentivized testnet.
	// It contains ~7.4M tokens, of which ~5.6M are to be distributed.
	FundingAccount = "evmos1f7vxxvmd544dkkmyxan76t76d39k7j3gr8d45y"

	// The remaining tokens are sent to the Community Pool.
	CommunityPoolAccount = "evmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8974jnh"
)
