package bytes

import (
	"testing"
)

func TestBuffer(t *testing.T) {
	p := NewPool(2, 10)
	if p.totalMax != 20 || p.freeMax != 20 || p.totalNum != 2 || p.freeNum != 2 {
		t.FailNow()
	}

	b := p.Get()
	if b.Bytes() == nil || len(b.buf) != 10 {
		t.FailNow()
	}
	if p.totalMax != 20 || p.freeMax != 10 || p.totalNum != 2 || p.freeNum != 1 {
		t.FailNow()
	}

	b = p.Get()
	if b.Bytes() == nil || len(b.Bytes()) != 10 {
		t.FailNow()
	}
	if p.totalMax != 20 || p.freeMax != 0 || p.totalNum != 2 || p.freeNum != 0 {
		t.FailNow()
	}

	b = p.Get()
	if b.Bytes() == nil || len(b.Bytes()) != 10 {
		t.FailNow()
	}

	p.Put(b)
	if p.totalMax != 40 || p.freeMax != 20 || p.totalNum != 4 || p.freeNum != 2 {
		t.FailNow()
	}
}
