package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tendermint/tendermint/crypto/tmhash"
	ethermint "github.com/tharsis/ethermint/types"
)

// NewTokenPair returns an instance of TokenPair
func NewTokenPair(erc20Address common.Address, denom string, enabled bool) TokenPair {
	return TokenPair{
		Erc20Address: erc20Address.String(),
		Denom:        denom,
		Enabled:      true,
	}
}

// GetID returns the SHA256 hash of the ERC20 address and denomination
func (b TokenPair) GetID() []byte {
	id := b.Erc20Address + "|" + b.Denom
	return tmhash.Sum([]byte(id))
}

// GetErc20Contract casts the hex string address of the ERC20 to common.Address
func (b TokenPair) GetERC20Contract() common.Address {
	return common.HexToAddress(b.Erc20Address)
}

// Validate performs a stateless validation of a TokenPair
func (b TokenPair) Validate() error {
	if err := sdk.ValidateDenom(b.Denom); err != nil {
		return err
	}

	if err := ethermint.ValidateAddress(b.Erc20Address); err != nil {
		return err
	}

	return nil
}
