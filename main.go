package main

import (
	"os"

	"go.resumic.org/schema/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
