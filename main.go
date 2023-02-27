package main

import (
	"fmt"
	"time"

)


 func main() {
	// variável que captura o impute do usuário do tipo tempo
	var tempoTrabalho time.Duration
	var tempoDescanso time.Duration
	var tempoMaior time.Duration
	
	// impute do usuário
	fmt.Println("Digite o tempo em minutos para estar trabalhando/focado")
	fmt.Scan(&(tempoTrabalho))

	fmt.Println("Digite o tempo de descanso")
	fmt.Scan(&(tempoDescanso))

	fmt.Println("Digite o tempo maior de descanso")
	fmt.Scan(&(tempoMaior))
	
	var tempoFocado = tempoTrabalho * time.Minute
	var tempoDeDescanso = tempoDescanso * time.Minute
	var tempoDescansoMaior = tempoMaior * time.Minute

	for {
		// tempo trabalhando/focado
		fmt.Println("Comece a trabalhar/focar")
		// para a execução do programa durante o período de tempo da var tempoTrabalho
		time.Sleep(tempoFocado) 
		tempoFocado++

		if tempoFocado >= 4 {
			time.Sleep(tempoDescansoMaior)
			tempoFocado = 0
		} else {
			// tempo de descanso
			fmt.Println("Começo do descanso")
			time.Sleep(tempoDeDescanso)
		}
		
            }
 }