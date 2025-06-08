package pqueue

import (
	"testing"

	"untref-ayp2/guia-monticulos-binarios/persona"

	"github.com/stretchr/testify/assert"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	assert.NotNil(t, pq)
	assert.True(t, pq.IsEmpty())
}

func TestIsEmptyAfterEnqueue(t *testing.T) {
	pq := NewPriorityQueue()
	p := persona.NewPesona("Alice", 25)
	pq.Enqueue(p)
	assert.False(t, pq.IsEmpty())
	_, _ = pq.Dequeue()
	assert.True(t, pq.IsEmpty())
}

func TestDequeueEmpty(t *testing.T) {
	pq := NewPriorityQueue()
	p, err := pq.Dequeue()
	assert.Error(t, err)
	assert.Nil(t, p)
}

func TestEnqueueDequeueOrder(t *testing.T) {
	pq := NewPriorityQueue()
	p1 := persona.NewPesona("Bob", 30)
	p2 := persona.NewPesona("Carol", 20)
	p3 := persona.NewPesona("Dave", 40)
	pq.Enqueue(p1)
	pq.Enqueue(p2)
	pq.Enqueue(p3)

	first, err := pq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, p3, first)

	second, err := pq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, p1, second)

	third, err := pq.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, p2, third)
}
