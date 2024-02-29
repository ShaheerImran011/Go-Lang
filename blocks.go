

package blocks

import (
    "fmt"
)


type Block struct {
    ID   int
    Data string
}


func DisplayAllBlocks(blocks []Block) {
    fmt.Println("Blocks in the system:")
    for _, block := range blocks {
        fmt.Printf("ID: %d, Data: %s\n", block.ID, block.Data)
    }
}


func NewBlock(id int, data string) Block {
    return Block{ID: id, Data: data}
}


func ModifyBlock(blocks []Block, id int, newData string) {
    for i := range blocks {
        if blocks[i].ID == id {
            blocks[i].Data = newData
            fmt.Printf("Block with ID %d modified successfully.\n", id)
            return
        }
    }
    fmt.Printf("Block with ID %d not found.\n", id)
}
