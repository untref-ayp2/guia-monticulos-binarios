package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHeap(t *testing.T) {
	h := NewHeap(func(a, b int) int { return a - b })
	assert.NotNil(t, h)
	assert.Equal(t, 0, h.Size())
}

func TestSizeInsert(t *testing.T) {
	h := NewHeap(func(a, b int) int { return a - b })
	assert.Equal(t, 0, h.Size())
	h.Insert(10)
	assert.Equal(t, 1, h.Size())
	h.Insert(20)
	assert.Equal(t, 2, h.Size())
}

func TestTopEmpty(t *testing.T) {
	h := NewHeap(func(a, b int) int { return a - b })
	_, err := h.Top()
	assert.Error(t, err)
}

func TestTop(t *testing.T) {
	h := NewHeap(func(a, b int) int { return a - b })
	elems := []int{5, 3, 8, 1, 9}
	for _, e := range elems {
		h.Insert(e)
	}
	top, err := h.Top()
	assert.NoError(t, err)
	assert.Equal(t, 1, top)
}

func TestRemoveEmpty(t *testing.T) {
	h := NewHeap(func(a, b int) int { return a - b })
	val, err := h.Remove()
	assert.Error(t, err)
	var zero int
	assert.Equal(t, zero, val)
}

func TestRemoveSequence(t *testing.T) {
	h := NewHeap(func(a, b int) int { return a - b })
	elems := []int{4, 2, 7, 1, 3}
	for _, e := range elems {
		h.Insert(e)
	}
	sorted := []int{1, 2, 3, 4, 7}
	for _, exp := range sorted {
		val, err := h.Remove()
		assert.NoError(t, err)
		assert.Equal(t, exp, val)
	}
	_, err := h.Remove()
	assert.Error(t, err)
}

func TestMaxHeapBehavior(t *testing.T) {
	h := NewHeap(func(a, b int) int { return b - a })
	elems := []int{4, 2, 7, 1, 3}
	for _, e := range elems {
		h.Insert(e)
	}
	sortedDesc := []int{7, 4, 3, 2, 1}
	for _, exp := range sortedDesc {
		val, err := h.Remove()
		assert.NoError(t, err)
		assert.Equal(t, exp, val)
	}
}
