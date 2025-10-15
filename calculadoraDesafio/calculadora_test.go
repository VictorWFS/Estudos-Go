package main

import "testing"

func shouldSumCorrect(t *testing.T) {
	teste := Soma(1, 5, 10)
	resultado := 16
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func shouldMultCorrect(t *testing.T) {
	teste := Multiplicar(10, 5)
	resultado := 50
	if teste != resultado {
		t.Error("Valor esperado: 50", resultado, "Valor retornado: ", teste)

	}
}

func shouldSubtractCorrect(t *testing.T) {
	teste := Subtrair(10, 5)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)

	}
}

func shouldDivisionCorrect(t *testing.T) {
	teste, err := Dividir(10, 4)
	resultado := 2.5
	if err != nil {
		t.Fatal("A função Dividir retornou um erro inesperado: ", err)
	}
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}
