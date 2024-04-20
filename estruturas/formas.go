package main

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return 0
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return 0
}
