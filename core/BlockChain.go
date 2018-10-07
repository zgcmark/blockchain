package core

type BlockChain struct {
	blockchain []*Block
}

func (bc *BlockChain) init() {
	//初始化区块链(创世区块)
	godblock := NewGodblock()
	bc.blockchain = append(bc.blockchain, godblock)

}

func (bc *BlockChain) AddBlock(data string) {
	//取出链表中上一个区块的hash
	preblock := bc.blockchain[len(bc.blockchain)-1]
	newblock := NewBlock(data, preblock.Hash)
	bc.blockchain = append(bc.blockchain, newblock)

}
