package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"math"
)

type Point []float64

type Group []Point

func main() {
	fileName1 := "distancia.txt"
	fileName2 := "entrada.txt"
	var limit float64
	var vecPoint []Point
	var dim int

	limit = getLimit(fileName1)
	vecPoint, dim = storePoints(fileName2)
	//fmt.Println(dim)
	//printSlice(vecPoint)
	//fmt.Println(vecPoint[0])
	groups := lider(vecPoint, limit, dim)
	sse := calcSSE(groups, dim)
	fmt.Println(sse)
}

func lider(vecPoint []Point, limit float64, dim int) ([][]Point){
	var groups [][]Point
	var l Point
	// pt := vecPoint[0]
	// group0 := makeGroup(pt)
	// groups = append(groups, group0)
	for _, pt := range vecPoint {
		lider := true
		for j := 0; j < len(groups); j++ {
			l = groups[j][0]
			dist := getDistance(pt, l, dim)
			if(dist <= limit) {
				groups[j] = append(groups[j], pt)
				lider = false
				break
			}
		}
		if (lider){
			newgroup := makeGroup(pt)
			groups = append(groups, newgroup)
		}
	}

	//fmt.Println("DENTRO DA LÍDER")
	//printGroups(groups, len(groups))
	return groups
}

func calcSSE(groups [][]Point, dim int) float64 {
	var sse float64
	for i := 0; i < len(groups); i++ {
		c := calcCentroid(groups[i], dim)
		for _, pt := range groups[i] {
			x := getDistance(c, pt, dim)
			//fmt.Println(x)
			sse = sse + x*x
		}
	}
	return sse	
}

func calcCentroid(group []Point, dim int) Point {
	n := len(group)
	var c Point
	for j := 0; j < dim; j++ {
		c = append(c,group[0][j])
	}
	for i := 1; i < n; i++ {
		for j := 0; j < dim; j++ {
			c[j] = c[j] + group[i][j]
		}
	}
	for j := 0; j < dim; j++ {
		c[j] = c[j] / float64(n)
	}
	//fmt.Println(c)
	return c
}

func makeGroup(lider Point) ([]Point) {
	return []Point{lider}
}

func getDistance(pt1 Point, pt2 Point, dim int) float64 {
	var soma float64
	for i := 0; i < dim; i++ {
		soma = soma + ((pt1[i] - pt2[i])*(pt1[i] - pt2[i]))
	}
	distancia := math.Sqrt(soma)
	return distancia
}

func printGroups(groups [][]Point, tam int){
	for i := 0; i < tam; i++ {
		printSlice(groups[i])
	} 
}

func getLimit(fileName string) float64 {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("DEU RUIM")
	}
	var limit float64
	n, err := fmt.Fscanln(file, &limit)
	if n == 0 || err != nil {
		fmt.Println("ERROR: Não foi possível fazer a leitura do limite")
    }
	return limit
}

func storePoints(fileName string) ([]Point, int) {
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
			//fmt.Println(firstline)
			firstreader := strings.NewReader(firstline)
			for {
				_, err := fmt.Fscan(firstreader, &myfloat)
				if err != nil {
					break;
				}
				p = append(p, myfloat)
				dim++ 
			}
			p = append(p, float64(i+1)) // O último elemento representa a linha em que o ponto foi lido
			vecPoint = append(vecPoint, p)
			//fmt.Println(dim)
			//fmt.Println("Passou a primeira linha")
		}else {
			//fmt.Println("Próxima linha")
			line := scanner.Text()
			//fmt.Println(line)
			reader := strings.NewReader(line)
			p = nil
			for j := 0; j < dim; j++ {
				fmt.Fscan(reader, &myfloat)
				p = append(p,myfloat)
			}
			p = append(p, float64(i+1)) // O último elemento representa a linha em que o ponto foi lido
			vecPoint = append(vecPoint,p)
			//printSlice(vecPoint)
		}
		i++
	}
	return vecPoint, dim
}

func printSlice(s []Point) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}