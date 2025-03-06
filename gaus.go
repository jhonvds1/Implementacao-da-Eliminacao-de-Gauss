package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAXN = 2000

var N int
var A [MAXN][MAXN]float64
var B [MAXN]float64
var X [MAXN]float64

func initializeInputs(n int) {
	N = n
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			A[i][j] = rand.Float64()
		}
		B[i] = rand.Float64()
		X[i] = 0.0
	}
}

func printMatrix() {
	if N > 10 {
		return
	}
	fmt.Println("A =")
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%5.2f ", A[i][j])
		}
		fmt.Println()
	}
	fmt.Println("\nB =")
	for i := 0; i < N; i++ {
		fmt.Printf("%5.2f ", B[i])
	}
	fmt.Println()
}

func gauss() {
	for norm := 0; norm < N-1; norm++ {
		for row := norm + 1; row < N; row++ {
			multiplier := A[row][norm] / A[norm][norm]
			for col := norm; col < N; col++ {
				A[row][col] -= A[norm][col] * multiplier
			}
			B[row] -= B[norm] * multiplier
		}
	}

	// Substituição regressiva
	for row := N - 1; row >= 0; row-- {
		X[row] = B[row]
		for col := N - 1; col > row; col-- {
			X[row] -= A[row][col] * X[col]
		}
		X[row] /= A[row][row]
	}
}

func printResult() {
	if N > 100 {
		return
	}
	fmt.Println("\nX =")
	for i := 0; i < N; i++ {
		fmt.Printf("%5.2f ", X[i])
	}
	fmt.Println()
}

func main() {
	N = 5 // Defina o tamanho desejado
	initializeInputs(N)
	printMatrix()

	start := time.Now()
	gauss()
	elapsed := time.Since(start)

	printResult()
	fmt.Printf("\nElapsed time = %v ms\n", elapsed.Milliseconds())
}
