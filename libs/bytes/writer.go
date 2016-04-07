package bytes

type Writer struct {
	n   int
	buf []byte
}

// NewWriterSize returns a size n writer.
func NewWriterSize(n int) *Writer {
	return &Writer{buf: make([]byte, n)}
}

// Size returns writer's size.
func (w *Writer) Size() int {
	return len(w.buf)
}

// Reset reset writer's n.
func (w *Writer) Reset() {
	w.n = 0
}

// Buffer returns writer's current byte array.
func (w *Writer) Buffer() []byte {
	return w.buf[:w.n]
}

// Peek set writer's cursor to next n bytes and returns skipped bytes.
func (w *Writer) Peek(n int) []byte {
	var buf []byte
	w.grow(n)
	buf = w.buf[w.n : w.n+n]
	w.n += n
	return buf
}

// Write append p to writer's buf.
func (w *Writer) Write(p []byte) {
	w.grow(len(p))
	w.n += copy(w.buf[w.n:], p)
}

// grow extent writer's buf size to 2*cur_size + n if needed.
func (w *Writer) grow(n int) {
	var buf []byte
	if w.n+n < len(w.buf) {
		return
	}
	buf = make([]byte, 2*len(w.buf)+n)
	copy(buf, w.buf[:w.n])
	w.buf = buf
	return
}
