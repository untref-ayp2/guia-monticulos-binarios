package pqueue

import (
	"untref-ayp2/guia-monticulos-binarios/heap"
	"untref-ayp2/guia-monticulos-binarios/persona"
)

type PriorityQueue struct {
	items *heap.Heap[*persona.Persona]
}

func NewPriorityQueue() *PriorityQueue {
	// Implementar
	return nil
}

func (pq *PriorityQueue) IsEmpty() bool {
	// Implementar
	return false
}

func (pq *PriorityQueue) Enqueue(p *persona.Persona) {
	// Implementar
}

func (pq *PriorityQueue) Dequeue() (*persona.Persona, error) {
	// Implementar
	return nil, nil
}
