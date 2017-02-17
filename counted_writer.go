package frandom

import "io"

type countedWriter struct {
	Count uint64
	Limit uint64
}

func (w *countedWriter) Write(b []byte) (n int, err error) {
	if w.Count+uint64(len(b)) > w.Limit {
		return 0, io.EOF
	}
	w.Count += uint64(len(b))
	return
}
