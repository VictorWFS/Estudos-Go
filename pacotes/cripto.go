//cripto são criptografados
//usar o sha-1 para criptografia

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	cripSha := sha1.New()
	cripSha.Write([]byte("Codigo em crypto sha1 com go"))
	valorSha := cripSha.Sum([]byte{})
	fmt.Println(valorSha)

}
