package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v : length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])

	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs1", dwarfs1) // length 5, capacity 5
	dwarfs2 := append(dwarfs1, "Orcus")
	dump("dwarfs2", dwarfs2) // length 6, capacity 10
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs3", dwarfs3) // legnth 9, capacity 10
}
