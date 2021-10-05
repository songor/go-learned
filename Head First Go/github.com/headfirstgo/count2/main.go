package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
	"sort"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	var names []string
	for name := range counts {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("Votes for %s: %d\n", name, counts[name])
	}
}
