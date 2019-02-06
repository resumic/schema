package main

import (
	"os"

	"github.com/resumic/schema/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
