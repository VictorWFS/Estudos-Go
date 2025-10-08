//aceita um conjunto de dados e o reduz a um tamanho fixo menor
//hashes são usadas em TUDO em programação, principalmente para buscar dados e detectar
//em go são divididas em criptografadas e não criptografadas
//lista das não cript: adler32, crc32, crc64
//nesse código vamos criar uma HASH e escrever nossos dados nele.

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	//criação do hash
	hash := crc32.NewIEEE()
	//escrever os dados no hash
	hash.Write([]byte("Cnodigo com pacote hash"))
	//calculo do hash
	valor := hash.Sum32()
	fmt.Println(valor)
}
