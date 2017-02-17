package frandom

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

//BufferSize is a suitable default buffer size
const BufferSize = 1024

//Rand generates randomness
type Rand struct {
	buf []byte
}

//New returns a new random generator
func New() *Rand {
	return &Rand{
		buf: make([]byte, BufferSize),
	}
}

//Read reads into b
func (r *Rand) Read(b []byte) (n int, err error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//panic(block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	for n < len(b) {
		toWrite := len(r.buf) % len(b)
		stream.XORKeyStream(b[n:n+toWrite], r.buf[:toWrite])
		n += toWrite
	}

	return
}

//WriteTo writes to a writer
func (r *Rand) WriteTo(wr io.Writer) (written int64, err error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//panic(block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	for {
		stream.XORKeyStream(r.buf[:], r.buf[:])
		n, err := wr.Write(r.buf[:])
		written += int64(n)
		if err != nil {
			return written, err
		}
	}
}
