package main

func main() {
	block := Block{
	 	Height: 1,
	 	Time: "the second",
	 	Txs_n: 2,
	 	Inner_txs_n: 3,
	 	Txs: "bowie",
	}
	
	insert(block);
	read()
}
