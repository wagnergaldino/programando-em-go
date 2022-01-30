package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)
	
func main() {
	palavras := os.Args[1:]
	estatisticas := colherEstatisticas(palavras)
	imprimir(estatisticas)
}

func colherEstatisticas(palavras []string) map[string]int {
	estatisticas := make(map[string]int)	
	for _, palavra := range palavras {
	
		for i := 0; i < utf8.RuneCountInString(palavra); i++ {
		
			inicial := strings.ToUpper(string(palavra[i]))
			contador, encontrado := estatisticas[inicial]		
			if encontrado {
				estatisticas[inicial] = contador + 1
			}else {
				estatisticas[inicial] = 1
			}
		
		}
			
	}
	return estatisticas
}

func imprimir(estatisticas map[string]int) {
	fmt.Printf("Contagem de cada letra:\n")	
	for inicial, contador := range estatisticas {
		fmt.Printf("%s = %d\n", inicial, contador)
	}
}
