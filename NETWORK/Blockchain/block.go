package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type Block struct {
	Timestamp   int64
	Hash        []byte
	Transaction []*Transaction
	PrevHash    []byte
	Nonce       int
}

type Transaction struct {
	ID          []byte
	CreateAdd   []byte
	Product     []byte
	Destination []byte
	Status      []byte
}

func (tx Transaction) SerializeTx() []byte {
	var encod bytes.Buffer
	err := gob.NewEncoder(&encod).Encode(tx)
	Handle(err)
	return encod.Bytes()
}

func DeserializeTx(data []byte) Transaction {
	var tx Transaction
	err := gob.NewDecoder(bytes.NewBuffer((data))).Decode(&tx)
	Handle(err)
	return tx
}

func (b *Block) SerializeBlock() []byte {
	var res bytes.Buffer
	err := gob.NewEncoder(&res).Encode(b)
	Handle(err)
	return res.Bytes()
}

func (tx *Transaction) Hash() []byte {
	var hash [32]byte
	txCopy := *tx
	txCopy.ID = []byte{}
	hash = sha256.Sum256(txCopy.SerializeTx())
	return hash[:]
}

func createBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte{}, txs, prevHash, 0}
	return block
}

func Genesis(Nullity *Transaction) *Block {
	return createBlock([]*Transaction{Nullity}, []byte{})
}
