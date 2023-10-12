package main

import (
	"flag"

	"github.com/wojiaowenzhong233/graph_order_homework/sorter"
)

func main() {

	filepath := flag.String("input", "", "Graph file to sort.")
	output := flag.String("output", "", "Output file name.")

	flag.Parse()

	if *filepath == "" {
		flag.Usage()
		return
	}

	g_sorter := sorter.CreateSorter("degree")
	G := sorter.CreateGraph(*filepath)

	sorter.SaveGraph(*output, g_sorter.Sort(G))
}
