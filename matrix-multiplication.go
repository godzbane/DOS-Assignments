package main

import "time"
import "math/rand"
import "runtime"
import "fmt"

var matrixSize = 1


func main() {

    r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		runtime.GOMAXPROCS(runtime.NumCPU())
		
		matrix1 := make([][]int, matrixSize)
		matrix2 := make([][]int, matrixSize)
		result  := make([][]int, matrixSize)
		
		
		for i := 0; i < matrixSize; i++ {
		
				matrix1[i] = make([]int, matrixSize)
				matrix2[i] = make([]int, matrixSize)
				result[i] = make([]int, matrixSize)
		
				for j := 0; j < matrixSize; j++ {
		
					matrix1[i][j] = r1.Intn(5)
					matrix2[i][j] = r1.Intn(5)
				
				}
		
		}
		
		println ("Matrix 1:")
		
		for i := 0; i < matrixSize; i++ {
			for j := 0; j < matrixSize; j++ {
				fmt.Printf ("%d ", matrix1[i][j])
				}
			println("")	
			}
			
			println ("\nMatrix 2:")
		
		for i := 0; i < matrixSize; i++ {
			for j := 0; j < matrixSize; j++ {
				fmt.Printf ("%d ", matrix2[i][j])
				}
			println("")	
			}
			
			println ("\nStart with empty result matrix:")
		for i := 0; i < matrixSize; i++ {
			for j := 0; j < matrixSize; j++ {
				fmt.Printf ("%d ", result[i][j])
				}
			println("")	
			}
		 
		
		for i := 0; i < matrixSize; i++ {
		
						go func () {
					
							ans := 0
				
							for j := 0; j < matrixSize; j++ {
					
							for k := 0; k < matrixSize; k++ {
								ans += matrix1[i][k]*matrix2[k][j]
								}
							result[i][j] = ans
							}
							
							
							
						} ()
					
		}
		
		println ("\nResult:")
		for i := 0; i < matrixSize; i++ {
			for j := 0; j < matrixSize; j++ {
				fmt.Printf ("%d ", result[i][j])
				}
			println("")	
			}
			
}


