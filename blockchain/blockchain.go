package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var dataDir string
var lock sync.Mutex

const defaultVersion = 1

type Hash string

type Transaction struct {
	ID      []byte
	Inputs  []IO `json:"inputs"`
	Outputs []IO `json:"outputs"`
}

type IO struct {
	Amount      float64 `json:"amount"`
	FromAddress Address `json:"from_address"`
	ToAddress   Address `json:"to_address"`
	Signature   string  `json:"signagure"`
}

type Block struct {
	Version      int           `json:"version"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash Hash          `json:"previous_hash"`
	Hash         Hash          `json:"hash"`
	Nonce        string        `json:"nonce"`
}

func InitBlockchain(dir string) error {
	dataDir = dir
	initBlock := Block{
		Version: defaultVersion,
		Hash:    Hash("init"),
	}
	return SaveBlock(initBlock)
}

func SaveBlock(block Block) error {
	lock.Lock()
	defer lock.Unlock()

	r, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(blockLocation(block.Hash), r, 0644)
}

func LoadBlock(hash Hash) (Block, error) {
	lock.Lock()
	defer lock.Unlock()

	block := Block{}

	data, err := ioutil.ReadFile(blockLocation(hash))
	err = json.Unmarshal(data, &block)

	return block, err
}

func blockLocation(hash Hash) string {
	return filepath.Join(dataDir, fmt.Sprintf("%s.json", string(hash)))
}
