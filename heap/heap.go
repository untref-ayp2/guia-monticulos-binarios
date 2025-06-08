package heap

type Heap[T any] struct {
	elements []T
	// compare es una función de comparación de elementos.
	// Para un heap de mínimo, devuelve un número negativo si `a < b`, un número
	// positivo si `a > b` o cero si son iguales.
	compare func(a, b T) int
}

// NewHeap crea un nuevo heap binario con una función de comparación
// personalizada.
//
// Uso:
//
//	heap := heap.NewHeap[int](func(a int, b int) int {
//		return a - b
//	})
//
// Parámetros:
//   - `compare` función de comparación personalizada.
//
// Retorna:
//   - un puntero a un heap binario con una función de comparación personalizada.
func NewHeap[T any](compare func(a, b T) int) *Heap[T] {
	// Implementar
	return nil
}

// Size retorna la cantidad de elementos en el heap.
//
// Uso:
//
//	size := heap.Size()
//
// Retorna:
// - la cantidad de elementos en el heap.
func (m *Heap[T]) Size() int {
	// Implementar
	return -1
}

// Insert agrega un elemento al heap.
//
// Uso:
//
//	heap := heap.NewHeap[int](func(a int, b int) int {
//		return a - b
//	})
//	heap.Insert(5)
//
// Parámetros:
//   - `element`: elemento a agregar al heap.
func (m *Heap[T]) Insert(element T) {
	// Implementar
}

// Remove elimina y retorna el elemento en la cima del heap.
//
// Uso:
//
//	heap := heap.NewHeap[int](func(a int, b int) int {
//		return a - b
//	})
//	heap.Insert(5)
//	element, _ := heap.Remove()
//
// Retorna:
//   - el elemento en la cima del heap.
//   - un error si el heap está vacío.
func (m *Heap[T]) Remove() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}

// Top retorna el elemento en la cima del heap.
//
// Uso:
//
//	heap := heap.NewHeap[int](func(a int, b int) int {
//		return a - b
//	})
//	heap.Insert(5)
//	element, _ := heap.Top()
//
// Retorna:
//   - el elemento en la cima del heap.
//   - un error si el heap está vacío.
func (m *Heap[T]) Top() (T, error) {
	// Implementar
	var zero T
	return zero, nil
}
