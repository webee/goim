package bytes

import (
	"sync"
)

// Buffer is a byte array.
type Buffer struct {
	buf  []byte
	next *Buffer // next free buffer
}

// Bytes return buffer's underly byte array.
func (b *Buffer) Bytes() []byte {
	return b.buf
}

// Pool is a buffer pool.
type Pool struct {
	lock     sync.Mutex
	free     *Buffer
	max      int //当前设置内存大小
	num      int //当前设置buffer数
	size     int //当前设置buffer size
	totalMax int
	freeMax  int
	totalNum int
	freeNum  int
}

// NewPool new a memory buffer pool struct.
func NewPool(num, size int) (p *Pool) {
	p = new(Pool)
	p.init(num, size)
	return
}

// Init init the memory buffer.
func (p *Pool) Init(num, size int) {
	p.init(num, size)
	return
}

// init init or new allocate the memory buffer.
func (p *Pool) init(num, size int) {
	p.lock.Lock()

	//更新当前设置
	p.num = num
	if p.size == 0 {
		//只能设置一次size
		p.size = size
	}
	p.max = num * size

	//new buffers
	p.grow()
	p.lock.Unlock()
}

// grow grow the memory buffer size, and update free pointer.
func (p *Pool) grow() {
	var (
		b   *Buffer
		bs  []Buffer
		buf []byte
	)
	buf = make([]byte, p.max)
	bs = make([]Buffer, p.num)
	//p.free = nil
	for i := p.num - 1; i >= 0; i-- {
		b = &bs[i]
		b.buf = buf[i*p.size : (i+1)*p.size]
		b.next = p.free
		p.free = b
	}

	//更新统计
	p.totalMax += p.max
	p.freeMax += p.max
	p.totalNum += p.num
	p.freeNum += p.num

	return
}

// Get get a free memory buffer.
func (p *Pool) Get() (b *Buffer) {
	p.lock.Lock()
	if b = p.free; b == nil {
		p.grow()
		b = p.free
	}
	p.free = b.next
	//更新统计
	p.freeMax -= p.size
	p.freeNum--

	p.lock.Unlock()
	return
}

// Put put back a memory buffer to free.
func (p *Pool) Put(b *Buffer) {
	p.lock.Lock()
	b.next = p.free
	p.free = b
	//更新统计
	p.freeMax += p.size
	p.freeNum++
	p.lock.Unlock()
	return
}
