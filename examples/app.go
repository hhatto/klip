package main

import (
	"fmt"
	"os"

	"github.com/hhatto/klip"
)

func main() {
	clips, _ := klip.Load(os.Args[1])
	for i := range clips {
		if clips[i].Meta.Type != klip.Highlight {
			continue
		}
		fmt.Printf("title  : %s\n", clips[i].Title)
		fmt.Printf("author : %v\n", clips[i].Author)
		fmt.Printf("content: %v\n", clips[i].Content)
		fmt.Println("===")
	}
}
