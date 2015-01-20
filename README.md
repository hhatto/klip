# klip [![Build Status](https://travis-ci.org/hhatto/klip.png?branch=master)](https://travis-ci.org/hhatto/klip)
Go Parser for Amazon Kindle's clippings.txt file.

***this module ALPHA version***


## Installation
```sh
go get github.com/hhatto/klip
```

## Usage
```go
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
```

## License
MIT
