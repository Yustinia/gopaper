package main

import (
	"fmt"

	"github.com/Yustinia/gopaper/internal"
)

func main() {
	client := internal.NewClient("ILXDm2WUnAwvbEkwr7WwulmT6kqwbvP7")
	params := internal.NewSearch()
	params.Page = 1
	params.Purity = "111"

	result, err := client.Search(params)
	if err != nil {
		panic(err)
	}

	for i, wall := range result.Wall {
		fmt.Printf("[%d] %s\n", i, wall.Path)
	}
}
