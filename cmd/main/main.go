package main

import (
	"github.com/johnmcdnl/johncoin"
	"github.com/johnmcdnl/johncoin/cmd/cli"
)

func main() {
	bc := johncoin.NewBlockchain()
	//defer bc.db.Close()

	c := cli.CLI{BlockChain: bc}
	c.Run()
}
