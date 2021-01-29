package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func find(C int, star [][]int, what int) (int, int) {
	var preComp [][]int
	var prev int
	preComp = make([][]int, len(star))

	if C == what {
		return C, C
	}
	for i, ray := range star { // prefixové součty pro všechny paprsky od jádra
		preComp[i] = make([]int, len(star[i]))
		prev = 0
		for j, f := range ray {
			preComp[i][j] = f + prev
			prev += f
		}
	}
	for i, ray := range preComp { // hledá mezi kombinacemi v paprsích
		for j := range ray {
			for k := range ray[j:] {
				if j == 0 { // od začátku
					if ray[k] == what { // bez jádra
						return star[i][0], star[i][k]
					}
					if ray[k]+C == what { // s jádrem
						return C, star[i][k]
					}
				} else { // od prostřed
					if ray[k+1]-ray[j-1] == what {
						return star[i][k+1], star[i][j]
					}
				}
			}
		}
	}
	for i, start := range preComp[:len(preComp)-1] { // přidá kombinace přes jádro, tj. 2 paprsky
		for ii, end := range preComp[i+1:] {
			for j, s := range start {
				for k, e := range end {
					if s+C+e == what {
						return star[i][j], star[ii+1][k]
					}
				}
			}
		}
	}
	return -1, -1
}

func main() {
	var C, R int // jádro, počet paprsků
	var star [][]int
	var count, what int
	reader := bufio.NewReader(os.Stdin)

	fmt.Scan(&R, &C)
	star = make([][]int, R)
	for i := 0; i < R; i++ { // přečte jednotlivé řádky a rozdělí je podle mezer
		var tmpS string
		tmpB, _ := reader.ReadBytes('\n')
		tmpS = string(tmpB[:len(tmpB)-1])
		split := strings.Split(tmpS, " ")
		star[i] = make([]int, len(split))
		for j := range split {
			fmt.Sscan(split[j], &star[i][j])
		}
	}
	fmt.Fscan(reader, &count) // načte hledané frekvence
	for i := 0; i < count; i++ {
		fmt.Fscan(reader, &what)
		a, b := find(C, star, what)
		fmt.Println(a, b)
	}
}
