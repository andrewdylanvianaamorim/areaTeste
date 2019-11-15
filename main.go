package areaTeste

//Pi é um proporção numérica
const Pi = 3.1416

//CalculaAreaCircunferencia pede o raio e retorna a área calculada
func CalculaAreaCircunferencia(raio float64) float64 {
	return Pi * raio * raio
}

//CalculaAreaRetangulo pede a base e altura do retângulo e retorna a área claculada
func CalculaAreaRetangulo(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
