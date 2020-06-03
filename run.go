package main

import (
	"fmt"
	"math/rand"
)

var step = map[int][2]int{
	0: {1, 0},
	1: {0, -1},
	2: {-1, 0},
	3: {0, 1},
}

func addIndex(ind1 [2]int, ind2 [2]int) [2]int {
	return [2]int{ind1[0] + ind2[0], ind1[1] + ind2[1]}
}

func getS(cnt, stepIndex int, currIndex [2]int, c chan [2]int) {
	if cnt <= 0 {
		close(c)
		return
	}
	inc := step[stepIndex]

	for indexCnt := cnt; indexCnt > 0; indexCnt-- {
		currIndex = addIndex(inc, currIndex)
		c <- currIndex
	}

	if stepIndex%2 == 1 {
		cnt--
	}
	getS(cnt, (stepIndex+1)%4, currIndex, c)
}

func getSequence(n int) chan [2]int {
	c := make(chan [2]int)
	go getS(n, 3, [2]int{0, -1}, c)
	return c
}

func generateArr(n int, random bool) [][]int {
	val := 0
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
		for j := range arr[i]{
			if random{
				arr[i][j] = rand.Intn(100)
			}else {
				arr[i][j] = val
				val++
			}
		}
	}

	return arr
}

func main() {
	n := 5
	resArr:= make([]int, 0, n*n)
	arr := generateArr(n, false)
	fmt.Println(arr)

	for index := range getSequence(n){
		resArr = append(resArr, arr[index[0]][index[1]])
	}
	fmt.Println(resArr)
}
