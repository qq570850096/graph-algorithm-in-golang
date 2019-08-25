package main

import (
	"Graph_algo/Adj"
	"Graph_algo/search"
	"fmt"
	"io"
)

func main() {
	var (
		mar *Adj.Hash
		//cc *DFS.CC
		cycleDetection *search.Cycle
		mar2           *Adj.Hash
	)
	mar = &Adj.Hash{}
	mar2 = &Adj.Hash{}
	if err := mar.ReadFromFile("g2_noCycle.txt"); err != nil && err != io.EOF {
		panic(err)
	}
	if err := mar2.ReadFromFile("g2.txt"); err != nil && err != io.EOF {
		panic(err)
	}
	//single = new(search.SingleSource)
	//single.Init(mar,0)
	//fmt.Println(single.Pre())
	//fmt.Println("0 -> 6 : ",single.Path(6))
	//path = &search.Path{
	//	T:5,
	//	SingleSource:single,
	//}
	//fmt.Println(path.Path())
	//fmt.Println(path.Visited())
	cycleDetection = new(search.Cycle)
	cycleDetection.Init(mar)
	fmt.Println(cycleDetection.HasCycle())
	cycleDetection.Init(mar2)
	fmt.Println(cycleDetection.HasCycle())
}
