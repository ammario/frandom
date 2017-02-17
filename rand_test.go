package frandom

import (
	"io"
	"testing"
)

func TestRandom(t *testing.T) {
	t.Run("Read", func(t *testing.T) {
		buf := make([]byte, 256000)
		rand := New()
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}
		var nullCount int
		for _, byt := range buf {
			if byt == 0x00 {
				nullCount++
			}
		}
		if nullCount > 1500 {
			t.Fatalf("%v null bytes", nullCount)
		}
	})
}

func BenchmarkRandom(b *testing.B) {

	wr := &countedWriter{}

	rand := New()

	wr.Count = 0
	wr.Limit = uint64(b.N)
	b.ResetTimer()
	if _, err := rand.WriteTo(wr); err != io.EOF {
		panic(err)
	}
	b.SetBytes(1)

}
