package main

import (
	"fmt"
	"os"

	"github.com/Ayobami0/phoenix/internal/cli"
)

func main() {
	if err := cli.New().Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
