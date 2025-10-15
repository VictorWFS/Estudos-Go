package main

import "testing"

//padrão triple A
//A - Arrange (Preparar)
//A - Act (Rodar)
//A - Assert (Verificar as asserções)

func ShouldSumCorrect(t *testing.T) {
	//arrange
	teste := soma(3, 2, 1)
	//act
	resultado := 6
	//Assert
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado:", teste)
	}
}

func ShouldSumIncorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func ShouldMultCorrect(t *testing.T) {
	teste := Multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}
