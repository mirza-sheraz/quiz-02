// main.go

package main

import (
	"quiz-02/blocks"
)

func main() {
	// Create some blocks
	block1 := blocks.NewBlock(1, "Block 1")
	block2 := blocks.NewBlock(2, "Block 2")

	// Display all blocks
	blocks.DisplayAllBlocks([]*blocks.Block{block1, block2})

	// Modify block 1
	blocks.ModifyBlock(block1, "Modified Block 1")

	// Display all blocks again
	blocks.DisplayAllBlocks([]*blocks.Block{block1, block2})
}
