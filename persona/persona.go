package persona

type Persona struct {
	nombre string
	edad   int
}

func NewPesona(nombre string, edad int) *Persona {
	return &Persona{nombre: nombre, edad: edad}
}

func (p *Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) GetEdad() int {
	return p.edad
}

func ComparadorDeMenorAMayorEdad(a, b Persona) int {
	return a.edad - b.edad
}

func ComparadorDeMayorAMenorEdad(a, b Persona) int {
	return b.edad - a.edad
}
