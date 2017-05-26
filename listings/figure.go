type Figure struct {
	ID, Quant                int
	Matrix                   [][]float64
	Width, Height, AngleStep float64
	Primitives               []Primitive
	MassCenter               Point
}
