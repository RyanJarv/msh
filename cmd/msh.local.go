package main

import (
	"flag"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/local"
	"strings"
)

func main() {
	flag.Parse()
	L.Debug.Println("received args:", strings.Join(flag.Args(), " "))
	err := local.Run(flag.Args())
	if err != nil {
		L.Error.Fatalln(err)
	}

	L.Debug.Println("ran successfully:", strings.Join(flag.Args(), " "))
}
