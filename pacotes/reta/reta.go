package main

import "math"

//Iniciando com letra maiúscula, será público(visível dentro e fora do pacote)
//Iniciando com letra minúscula, será privado(em todo o pacote)

type Ponto struct {
	//struct pública
	//Ponto representa uam coordenada no plano cartesiano
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
