package frandom

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"runtime"
	"sync/atomic"
	"time"
)

//BufferSize is a suitable default buffer size
const BufferSize = 1024

//SeedLimit specifies the number of bytes Rand can generate before reseeding
const SeedLimit = 1024 * 1024

var (
	counter uint64

	//mutex is set to 1 if a thread is currently holding it
	mutex uint64

	buf    = make([]byte, BufferSize)
	stream cipher.Stream
)

func init() {
	seed()
}

//seed seeds the rng
func seed() {
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

	stream = cipher.NewOFB(block, iv)

	//start buf off nice and random
	//pretty sure this doesn't affect security but oh well.
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
}

func read(b []byte) (n int, err error) {
	if counter > SeedLimit {
		seed()
		counter = 0
	}

	counter += uint64(len(b))
	stream.XORKeyStream(b, b)
	return len(b), nil
}

//Read reads into b.
//Read is thread safe.
//It never returns an error.
func Read(b []byte) (n int, err error) {
	//makeshift block until can get state
	for !atomic.CompareAndSwapUint64(&mutex, 0, 1) {
		runtime.Gosched()
	}
	n, err = read(b)
	atomic.StoreUint64(&mutex, 0)
	return
}

//WriteTo writes to a writer
func WriteTo(wr io.Writer) (written int64, err error) {
	for !atomic.CompareAndSwapUint64(&mutex, 0, 1) {
		time.Sleep(time.Microsecond)
	}
	for {
		read(buf)
		n, err := wr.Write(buf)
		written += int64(n)
		if err != nil {
			atomic.StoreUint64(&mutex, 0)
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
