package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var _ Signer = callbackSigner{}

type callbackSigner struct {
	chainID        *big.Int
	relayerAddress common.Address
}

func NewCallbackSigner(chainID *big.Int, relayerAddress common.Address) Signer {
	return callbackSigner{chainID: chainID, relayerAddress: relayerAddress}
}

func (c callbackSigner) Sender(tx *Transaction) (common.Address, error) {
	if c.relayerAddress == (common.Address{}) {
		return common.Address{}, fmt.Errorf("relayer address not set")
	}
	return c.relayerAddress, nil
}

func (c callbackSigner) SignatureValues(tx *Transaction, sig []byte) (r, s, v *big.Int, err error) {
	// TODO: Return dummy values for now
	return new(big.Int), new(big.Int), new(big.Int), nil
}

func (c callbackSigner) ChainID() *big.Int {
	return c.chainID
}

func (c callbackSigner) Hash(tx *Transaction) common.Hash {
	// Assume a simple RLP hash of the transaction is sufficient
	return rlpHash([]interface{}{tx.Nonce(), tx.GasPrice(), tx.Gas(), tx.To(), tx.Value(), tx.Data()})
}

func (c callbackSigner) Equal(signer Signer) bool {
	_, ok := signer.(callbackSigner)
	return ok
}
