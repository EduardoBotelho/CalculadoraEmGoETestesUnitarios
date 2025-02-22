package main

import "testing"

func ShouldSumCorrect(t *testing.T) { //convenção de nomes ShouldSum/correct(assinatura do método)
	//arrange
	teste := soma(2, 3)
	//act
	resultado := 5

	//assert
	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func ShouldSumIncorrect(t *testing.T) {
	teste := soma(2, 3)

	resultado := 6
	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func ShouldSubtCorrect(t *testing.T) {
	teste := subtracao(3, 2)
	resultado := 1
	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func ShouldSubtIncorrect(t *testing.T) {
	teste := subtracao(3, 2)
	resultado := 6
	if teste != resultado {
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func ShouldDiviCorrect(t *testing.T) {
	teste := divisao(8, 2)
	resultado := 4
	if teste != float64(resultado) {//conversao para mesmo tipo float64
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
func ShouldDiviIncorrect(t *testing.T) {
	teste := divisao(8, 2)
	resultado := 5
	if teste != float64(resultado) {//conversao para mesmo tipo float64
		t.Error("Valor Esperado: ", resultado, "Valor retornado: ", teste)
	}
}
