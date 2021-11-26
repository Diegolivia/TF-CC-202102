package blockchain

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}
