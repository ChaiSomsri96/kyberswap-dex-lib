package metavault

import (
	"math/big"

	"github.com/KyberNetwork/kyberswap-dex-lib/pkg/util/bignumber"
)

// VaultUtils
// https://github.com/gmx-io/gmx-contracts/blob/master/contracts/core/VaultUtils.sol
type VaultUtils struct {
	vault *Vault `msgpack:"-"`
}

func NewVaultUtils(vault *Vault) *VaultUtils {
	return &VaultUtils{
		vault: vault,
	}
}

func (u *VaultUtils) GetSwapFeeBasisPoints(tokenIn string, tokenOut string, usdmAmount *big.Int) *big.Int {
	isStableSwap := u.vault.StableTokens[tokenIn] && u.vault.StableTokens[tokenOut]

	var baseBps *big.Int
	if isStableSwap {
		baseBps = u.vault.StableSwapFeeBasisPoints
	} else {
		baseBps = u.vault.SwapFeeBasisPoints
	}

	var taxBps *big.Int
	if isStableSwap {
		taxBps = u.vault.StableTaxBasisPoints
	} else {
		taxBps = u.vault.TaxBasisPoints
	}

	feeBasisPoints0 := u.GetFeeBasisPoints(tokenIn, usdmAmount, baseBps, taxBps, true)
	feeBasisPoints1 := u.GetFeeBasisPoints(tokenOut, usdmAmount, baseBps, taxBps, false)

	if feeBasisPoints0.Cmp(feeBasisPoints1) > 0 {
		return feeBasisPoints0
	} else {
		return feeBasisPoints1
	}
}

func (u *VaultUtils) GetFeeBasisPoints(token string, usdmDelta *big.Int, feeBasisPoints *big.Int, taxBasisPoints *big.Int, increment bool) *big.Int {
	if !u.vault.HasDynamicFees {
		return feeBasisPoints
	}

	initialAmount := u.vault.USDMAmounts[token]
	nextAmount := new(big.Int).Add(initialAmount, usdmDelta)

	if !increment {
		if usdmDelta.Cmp(initialAmount) > 0 {
			nextAmount = bignumber.ZeroBI
		} else {
			nextAmount = new(big.Int).Sub(initialAmount, usdmDelta)
		}
	}

	targetAmount := u.vault.GetTargetUSDMAmount(token)

	if targetAmount.Cmp(bignumber.ZeroBI) == 0 {
		return feeBasisPoints
	}

	var initialDiff *big.Int
	if initialAmount.Cmp(targetAmount) > 0 {
		initialDiff = new(big.Int).Sub(initialAmount, targetAmount)
	} else {
		initialDiff = new(big.Int).Sub(targetAmount, initialAmount)
	}

	var nextDiff *big.Int
	if nextAmount.Cmp(targetAmount) > 0 {
		nextDiff = new(big.Int).Sub(nextAmount, targetAmount)
	} else {
		nextDiff = new(big.Int).Sub(targetAmount, nextAmount)
	}

	if nextDiff.Cmp(initialDiff) < 0 {
		rebateBps := new(big.Int).Div(new(big.Int).Mul(taxBasisPoints, initialDiff), targetAmount)

		if rebateBps.Cmp(feeBasisPoints) > 0 {
			return bignumber.ZeroBI
		} else {
			return new(big.Int).Sub(feeBasisPoints, rebateBps)
		}
	}

	averageDiff := new(big.Int).Div(new(big.Int).Add(initialDiff, nextDiff), bignumber.Two)

	if averageDiff.Cmp(targetAmount) > 0 {
		averageDiff = targetAmount
	}

	taxBps := new(big.Int).Div(new(big.Int).Mul(taxBasisPoints, averageDiff), targetAmount)

	return new(big.Int).Add(feeBasisPoints, taxBps)
}
