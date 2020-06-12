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
		for j := range arr[i] {
			if random {
				arr[i][j] = rand.Intn(100)
			} else {
				arr[i][j] = val
				val++
			}
		}
	}

	return arr
}
func isSquareMatrix(matr [][]int) bool {
	init_cap := cap(matr)
	for _, arr := range matr {
		if init_cap != cap(arr) {
			return false
		}
	}
	return true
}

func getTornadoMatrix(arr [][]int, cap int) []int {
	resArr := make([]int, 0, cap*cap)

	for index := range getSequence(cap) {
		resArr = append(resArr, arr[index[0]][index[1]])
	}
	return resArr
}

func main() {
	arr := generateArr(5, false)
	fmt.Println(arr)


	if isSquareMatrix(arr) {
		resArr:= getTornadoMatrix(arr, cap(arr))
		fmt.Println(resArr)
	} else {
		fmt.Println("Invalid matrix format")
	}
}
