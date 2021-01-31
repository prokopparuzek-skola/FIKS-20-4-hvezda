// psáno v jazyce golang
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func compute(C int, star [][]int) (out []int) {
	var preComp [][]int
	var prev int
	out = make([]int, 0)
	preComp = make([][]int, len(star))

	for i, ray := range star { // prefixové součty pro všechny paprsky od jádra
		preComp[i] = make([]int, len(star[i]))
		prev = 0
		for j, f := range ray {
			preComp[i][j] = f + prev
			prev += f
		}
	}
	for _, ray := range preComp { // přidá všechny kombinace v paprsích
		for j := range ray {
			for k := range ray[j:] {
				if j == 0 { // od začátku
					out = append(out, ray[k])   // bez jádra
					out = append(out, ray[k]+C) // s jádrem
				} else { // od prostřed
					out = append(out, ray[k+1]-ray[j-1])
				}
			}
		}
	}
	for i, start := range preComp[:len(preComp)-1] { // přidá kombinace přes jádro, tj. 2 paprsky
		for _, end := range preComp[i+1:] {
			for _, s := range start {
				for _, e := range end {
					out = append(out, s+C+e)
				}
			}
		}
	}
	out = append(out, C) // přidá jádro
	return
}

func main() {
	var C, R int // jádro, počet paprsků
	var star [][]int
	var frequencies []int
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
	frequencies = compute(C, star)
	sort.Slice(frequencies, func(i, j int) bool { return frequencies[i] < frequencies[j] })
	out := fmt.Sprint(frequencies)
	fmt.Println(out[1 : len(out)-1])
}
