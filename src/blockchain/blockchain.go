package blockchain

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce int
}

type Chain struct {
	Blocks []*Block
}

func InitChain() *Chain {
	return &Chain{[]*Block{CreateBlock("Genesis Block", []byte{})}}
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Work()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (chain *Chain) AddBlock(data string) {
	previousBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}