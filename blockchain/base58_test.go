package blockchain

import (
	"encoding/hex"
	"strings"

	check "gopkg.in/check.v1"
)

type Base58Suite struct{}

var _ = check.Suite(&Base58Suite{})

func (s *Base58Suite) TestWalletCreate(c *check.C) {
	rawHash := "00010966776006953D5567439E5E39F86A0D273BEED61967F6"
	hash, err := hex.DecodeString(rawHash)
	c.Assert(err, check.IsNil)

	encoded := Base58Encode(hash)
	c.Check("16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM", check.Equals, string(encoded))

	decoded := Base58Decode([]byte("16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM"))
	c.Check(strings.ToLower("00010966776006953D5567439E5E39F86A0D273BEED61967F6"), check.Equals, hex.EncodeToString(decoded))
}
