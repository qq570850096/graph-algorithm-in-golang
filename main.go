package main

import (
	"Graph_algo/Adj"
	"Graph_algo/DFS"
	"fmt"
	"io"
)

func main()  {
	var (
		mar *Adj.Hash
		cc *DFS.CC
	)
	mar = &Adj.Hash{}
	if err:=mar.ReadFromFile("g2.txt");err!=nil && err!=io.EOF{
		panic(err)
	}
	cc = new(DFS.CC)
	cc.Init(mar)
	fmt.Println(cc.Cccount())
	fmt.Println(mar)
	fmt.Println(mar.LinkedVertex(1))
	fmt.Println(mar.Degree(1))
}
