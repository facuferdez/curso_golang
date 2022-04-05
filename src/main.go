package main

import "fmt"

func main() {
	/*var mat [3][4]int

	for i := 0; i < 3; i++ {
		for f := 0; f < 4; f++ {
			fmt.Println("Ingrese los valores de la fila:")
			fmt.Scan(&mat[i][f])

		}

	}
	fmt.Println("Matriz completa")
	fmt.Println(mat)
	may := mat[0][0]

	for i := 0; i < 3; i++ {
		for f := 0; f < 4; f++ {
			if mat[i][f] > may {
				may = mat[i][f]
			}

		}

	}
	fmt.Println("El elemento mayor de la matriz es:", may)*/

	var mat [2][5]int

	for i := 0; i < 2; i++ {
		for f := 0; f < 5; f++ {
			fmt.Println("Ingresa los valores:")
			fmt.Scan(&mat[i][f])
		}

	}
	fmt.Println("Matriz completa")
	fmt.Println(mat)
	var cambio int

	for f := 0; f < 5; f++ {

		cambio = mat[0][f]
		mat[0][f] = mat[1][f]
		mat[1][f] = cambio
	}

	fmt.Println("Matriz completa cambiada")
	fmt.Println(mat)
	fmt.Println(mat)
	fmt.Println("llllll")

}
