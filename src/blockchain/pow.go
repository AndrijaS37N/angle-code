package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const (
	Difficulty = 13
)

type POW struct {
	Block  *Block
	Target *big.Int
}

func NewProof(block *Block) *POW {
	target := big.NewInt(1)
	// Left shift - Lsh().
	target.Lsh(target, uint(256-Difficulty))
	return &POW{block, target}
}

func (pow *POW) InitData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			pow.Block.PreviousHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
}

func (pow *POW) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}

func ToHex(number int64) []byte {
	buffer := new(bytes.Buffer)
	error := binary.Write(buffer, binary.BigEndian, number)
	if error != nil {
		log.Panic(error)
	}
	return buffer.Bytes()
}
func (pow *POW) Work() (int, []byte) {
	var intHash big.Int
	var hash [32] byte
	nonce := 0
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}
