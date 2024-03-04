package geometria

type Retangulo struct {
    Largura float64
    Altura  float64
}

func (r Retangulo) Area() float64 {
    return r.Largura * r.Altura
}

func (r Retangulo) Perimetro() float64 {
    return 2*r.Largura + 2*r.Altura
}