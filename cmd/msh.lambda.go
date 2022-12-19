package main

import (
	"flag"
	L "github.com/ryanjarv/msh/pkg/logger"
	"github.com/ryanjarv/msh/pkg/providers/lambda"
	"strings"
)

func main() {
	flag.Parse()
	L.Debug.Println("received args:", strings.Join(flag.Args(), " "))
	err := lambda.Run(flag.Args())
	if err != nil {
		L.Error.Fatalln(err)
	}
}
