package Estruturas

import "math"

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

func Area(retangulo Retangulo) float64 {
	return retangulo.Largura * retangulo.Altura
}
