package main

import (
	"os"

	"github.com/ammario/frandom"
)

func main() {
	rng := frandom.New()
	_, err := rng.WriteTo(os.Stdout)
	panic(err)
}
