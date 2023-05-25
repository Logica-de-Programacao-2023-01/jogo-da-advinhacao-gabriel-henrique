package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var JogarMaisUmaVez string
	TentativasTotais := []int{}
	var NumeroCorreto int

	fmt.Println("Bem vindo ao adivinhe o número. ")

	for {

		NumeroAleatorio := rand.Intn(100) + 1
		tentativas := 0
		fmt.Print("Adivinhe o número aleatório. Digite um número int de 1 a 100: ")
		fmt.Scan(&NumeroCorreto)
		for {
			if NumeroAleatorio > NumeroCorreto {
				fmt.Println("O número é maior que", NumeroCorreto)
				fmt.Print("Tente novamente: ")
				fmt.Scan(&NumeroCorreto)
				tentativas++

			} else if NumeroAleatorio < NumeroCorreto {
				fmt.Println("O número é menor que", NumeroCorreto)
				fmt.Print("Tente novamente: ")
				fmt.Scan(&NumeroCorreto)
				tentativas++

			} else if NumeroAleatorio == NumeroCorreto {
				fmt.Println("Muito bem você acertou!!!")
				tentativas++
				break
			}

		}
		TentativasTotais = append(TentativasTotais, tentativas)
		fmt.Println("Você precisou de %d tentativas.", tentativas)

		fmt.Println(" Deseja jogar mais uma vez?(sim/não)")
		fmt.Scan(&JogarMaisUmaVez)
		if JogarMaisUmaVez == "n" {
			for i := 0; i < len(TentativasTotais); i++ {

				fmt.Printf("você fez %d . ", TentativasTotais[i], i+1)
			}
			break
		}
	}

}
