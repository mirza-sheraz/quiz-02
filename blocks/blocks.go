// blocks.go

package blocks

import (
	"fmt"
)

// Block represents a data structure for a block
type Block struct {
	ID   int
	Data string
}

// DisplayAllBlocks prints all blocks in a slice
func DisplayAllBlocks(blockSlice []*Block) {
	fmt.Println("Blocks:")
	for _, block := range blockSlice {
		fmt.Printf("ID: %d, Data: %s\n", block.ID, block.Data)
	}
}

// NewBlock creates a new block with the given ID and data
func NewBlock(id int, data string) *Block {
	return &Block{
		ID:   id,
		Data: data,
	}
}

// ModifyBlock modifies the data of an existing block
func ModifyBlock(block *Block, newData string) {
	block.Data = newData
}
