package blockchain

import (
	"math/rand"

	hdwallet "github.com/WeMeetAgain/go-hdwallet"
)

type Address string

// Generate a private key with given seed
func GenerateKeys(seed []byte) *hdwallet.HDWallet {
	return hdwallet.MasterKey(seed)
}

// Generate a public address with given private key. Optionally choose if you
// want random public address or consistent
func GetPublicKey(wallet *hdwallet.HDWallet, useRand bool) (Address, error) {
	i := uint32(0)
	if useRand {
		i = uint32(rand.Intn(2 ^ 31))
	}
	pub, err := wallet.Child(i)
	return Address(pub.Address()), err
}
