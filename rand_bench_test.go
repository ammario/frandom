package frandom

import (
	"io"
	"testing"
)

func BenchmarkRandom(b *testing.B) {

	wr := &countedWriter{}

	wr.Count = 0
	wr.Limit = uint64(b.N)

	b.ResetTimer()
	if _, err := WriteTo(wr); err != io.EOF {
		panic(err)
	}
	b.SetBytes(1)

}
