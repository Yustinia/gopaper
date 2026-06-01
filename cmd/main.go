package main

import (
	"fmt"

	"github.com/Yustinia/gopaper/internal"
)

func main() {
	client := internal.NewClient("ILXDm2WUnAwvbEkwr7WwulmT6kqwbvP7")

	result, err := client.Search()
	if err != nil {
		panic(err)
	}

	for i, wall := range result.Wall {
		fmt.Printf("[%d] %s\n", i, wall.Path)
	}
}
