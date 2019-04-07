package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Point []float64

func main() {
	fileName := "entrada.txt"
	var vecPoint []Point
	var dim int
	vecPoint, dim = storePoints(fileName)
	fmt.Println(dim)
	printSlice(vecPoint)
}

func storePoints (fileName string) ([]Point, int) {
	file, err := os.Open(fileName) //Abrindo arquivo
	if err != nil {
		fmt.Println("DEU RUIM")
	}

	var i int
	var myfloat float64
	var dim int
	var p Point
	var vecPoint []Point
	scanner := bufio.NewScanner(file) //Criando um Scanner
	for scanner.Scan() {
		if i == 0 {
			firstline := scanner.Text()
			fmt.Println(firstline)
			firstreader := strings.NewReader(firstline)
			for {
				_, err := fmt.Fscan(firstreader, &myfloat)
				if err != nil {
					break;
				}
				p = append(p, myfloat)
				dim++ 
			}
			vecPoint = append(vecPoint, p)
			fmt.Println(dim)
			fmt.Println("Passou a primeira linha")
		}else {
			fmt.Println("Próxima linha")
			line := scanner.Text()
			fmt.Println(line)
			reader := strings.NewReader(line)
			p = nil
			for j := 0; j < dim; j++ {
				fmt.Fscan(reader, &myfloat)
				p = append(p,myfloat)
			}
			vecPoint = append(vecPoint,p)
			printSlice(vecPoint)
		}
		i++
	}
	return vecPoint, dim
}

func printSlice(s []Point) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}