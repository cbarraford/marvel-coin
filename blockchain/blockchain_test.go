package blockchain

import (
	"io/ioutil"
	"log"

	check "gopkg.in/check.v1"
)

type BlockchainSuite struct {
	tmpDir string
}

var _ = check.Suite(&BlockchainSuite{})

func (s *BlockchainSuite) SetUpSuite(c *check.C) {
	s.tmpDir, _ = ioutil.TempDir("", "marvel-tmp")
	log.Printf("TMP: %s", s.tmpDir)
}

func (s *BlockchainSuite) TearDownSuite(c *check.C) {
	//os.RemoveAll(s.tmpDir)
}

func (s *BlockchainSuite) TestInit(c *check.C) {
	c.Assert(InitBlockchain(s.tmpDir), check.IsNil)

	init, err := LoadBlock(Hash("init"))
	c.Assert(err, check.IsNil)
	c.Check(init.Hash, check.Equals, Hash("init"))
	c.Check(init.PreviousHash, check.Equals, Hash(""))
}
