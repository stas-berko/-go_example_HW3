package main

import "fmt"

var step = map[int][2]int{
	0: {1, 0},
	1: {0, -1},
	2: {-1, 0},
	3: {0, 1},
}

func addIndex(ind1 [2]int, ind2 [2]int) [2]int {
	return [2]int{ind1[0] + ind2[0], ind1[1] + ind2[1]}
}

func getS(cnt, stepIndex int, currIndex [2]int, stor [][2]int) [][2]int {
	if cnt <= 0 {
		return stor
	}
	inc := step[stepIndex]

	for indexCnt := cnt; indexCnt > 0; indexCnt-- {
		currIndex = addIndex(inc, currIndex)
		stor = append(stor, currIndex)
	}

	if stepIndex%2 == 1 {
		cnt--
	}
	return getS(cnt, (stepIndex+1)%4, currIndex, stor)
}

func main() {
	n := 4
	sequence := make([][2]int, 0, n*n)
	fmt.Println(getS(n, 3, [2]int{0, -1}, sequence))

}
