package main

import "fmt"

func main() {
	var slice []int
	var valor int
	var suma = 0

	for {
		fmt.Println("Ingrese los valores (-1 para terminar):")
		fmt.Scan(&valor)
		if valor == -1 {
			break
		}
		slice = append(slice, valor)
		suma = suma + valor
	}

	promedio := suma / len(slice)
	mayProm := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > promedio {
			mayProm++
		}
	}

	fmt.Println("EL slice completo es:")
	fmt.Println(slice)
	fmt.Println("Promedio:", promedio)
	fmt.Println("La cantidad de productos ingresados mayores al promedio es:", mayProm)
}
