package main

import (
	"os"

	"github.com/ammario/frandom"
)

func main() {
	_, err := frandom.New().WriteTo(os.Stdout)
	panic(err)
}
