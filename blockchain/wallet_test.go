package blockchain

import (
	"encoding/hex"

	check "gopkg.in/check.v1"
)

type WalletSuite struct{}

var _ = check.Suite(&WalletSuite{})

func (s *WalletSuite) TestWalletCreate(c *check.C) {
	w, err := NewWallet("hello bar glass tower pastor")
	c.Assert(err, check.IsNil)
	c.Assert(string(w.GetAddress()), check.Equals, "12SZgbXcJ1W6jCKXnSCj2k8UmZFDDTRM3g")
}

func (s *WalletSuite) TestSign(c *check.C) {
	w, err := NewWallet("")
	c.Assert(err, check.IsNil)

	sig, err := w.Sign("hello world")
	c.Assert(err, check.IsNil)
	c.Check(hex.EncodeToString(sig), check.Equals, "07e2b7f9229be748740479a2a276992b5408cffc91ef5e5c7192cdd3c714db82f81f5a913d2fac75e98b080f2be00ce5c589a89e520e36698ddd2d816ae77dc1")
	c.Check(w.Verify(sig), check.Equals, true)
}
