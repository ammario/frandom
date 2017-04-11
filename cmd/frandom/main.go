package main

import (
	"os"

	"github.com/ammario/frandom"

	"flag"
)

func main() {
	threads := flag.Int("t", 1, "allows for multiple generator threads. Should only be greater than 1 for testing entropy.")
	flag.Parse()
	for i := 0; i < *threads; i++ {
		go func() {
			_, err := frandom.WriteTo(os.Stdout)
			panic(err)
		}()
	}

	<-make(chan struct{})
}
