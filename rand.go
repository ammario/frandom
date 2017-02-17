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
	buf    []byte
	stream cipher.Stream
}

//New returns a new randomness generator
func New() *Rand {
	key := make([]byte, 16)
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

	rand := &Rand{
		buf:    make([]byte, BufferSize),
		stream: cipher.NewCTR(block, iv),
	}
	//start buf off nice and random
	//pretty sure this doesn't affect security but oh well.
	if _, err := rand.Read(rand.buf); err != nil {
		panic(err)
	}
	return rand
}

//Read reads into b.
//It never returns an error.
func (r *Rand) Read(b []byte) (n int, err error) {
	r.stream.XORKeyStream(b, b)
	return len(b), nil
}

//WriteTo writes to a writer
func (r *Rand) WriteTo(wr io.Writer) (written int64, err error) {
	for {
		r.stream.XORKeyStream(r.buf, r.buf)
		n, err := wr.Write(r.buf)
		written += int64(n)
		if err != nil {
			return written, err
		}
	}
}

// //UnsafeWriteTo writes to a wr without checking errors.
// func (r *Rand) UnsafeWriteTo(wr io.Writer) (written int64, err error) {
// 	for {
// 		r.stream.XORKeyStream(r.buf, r.buf)
// 		wr.Write(r.buf)
// 	}
// }
