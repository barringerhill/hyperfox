package main

type Block struct {
	Id int64
	Height int64  `xorm:"unique"`
	Time string   `xorm:"unique"`
	Txs_n int64
	Inner_txs_n int64
	Txs string
}

type Tx struct {
	Id int64
	Tx_id string   `xorm: "unique"`
	Height int64   `xorm: "unique"`
	Content string 
}
