package avs

import (
	"fmt"
	"math/big"

	"github.com/pkg/errors"

	geth "github.com/ethereum/go-ethereum/common"

	elCliUtils "github.com/Layr-Labs/eigenlayer-cli/pkg/utils"
)

func avsAddressOrDefault(avsAddr string, chainID *big.Int) (geth.Address, error) {
	var resp geth.Address
	if avsAddr != "" {
		if !geth.IsHexAddress(avsAddr) {
			return geth.Address{}, errors.New(fmt.Sprintf("invalid avs address: %s", avsAddr))
		}
		resp = geth.HexToAddress(avsAddr)
	} else if addr, ok := skateAvsAddressOnChain(chainID); ok {
		resp = addr
	} else {
		return geth.Address{}, errors.New(fmt.Sprintf("avs address not provided and no default for chain found: %d", chainID.Uint64()))
	}

	return resp, nil
}

// NOTE: update with our AVS address when deployment changes
func skateAvsAddressOnChain(chainID *big.Int) (geth.Address, bool) {
	switch chainID.Int64() {
	case elCliUtils.HoleskyChainId:
		return geth.HexToAddress("0x2a0D46ED3D9D13F6a9b5B0D3274675143c803071"), true
	case elCliUtils.MainnetChainId:
		return geth.Address{}, false
	default:
		return geth.Address{}, false
	}
}
