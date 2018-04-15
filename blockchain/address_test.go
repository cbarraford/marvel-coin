package blockchain

import (
	"testing"

	check "gopkg.in/check.v1"
)

func TestPackage(t *testing.T) { check.TestingT(t) }

type AddressSuite struct{}

var _ = check.Suite(&AddressSuite{})

func (s *AddressSuite) TestCreateKeys(c *check.C) {
	priv := GenerateKeys([]byte("hello world"))
	c1, err := priv.Child(1)
	c.Assert(err, check.IsNil)
	c2, err := priv.Child(2)
	c.Assert(err, check.IsNil)
	c.Check(c1.Address(), check.Equals, "1HFTWLEV2NZDPUoEYn4TxzBiCf838NGjXC")
	c.Check(c2.Address(), check.Equals, "168uP7SGKPT1EYDEpN2gnfXZ6vUZimN9gE")
}

func (s *AddressSuite) TestPubKey(c *check.C) {
	priv := GenerateKeys([]byte("hello world"))

	// test that we create the first public key ok
	addy, err := GetPublicKey(priv, false)
	c.Assert(err, check.IsNil)
	c.Check(addy, check.Equals, Address("1AWCuyy2yQqEHJvnAecQUEaFhutA6k1psX"))

	// check that we can create random public keys for our private key
	addy, err = GetPublicKey(priv, true)
	c.Assert(err, check.IsNil)
	c.Check(addy, check.Not(check.Equals), Address("1AWCuyy2yQqEHJvnAecQUEaFhutA6k1psX"))
}
