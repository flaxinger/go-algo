package main

import(
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)



func main() {
	
	var N int 
	in := bufio.NewReader(os.Stdin)

	fmt.Scan(&N)
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)


	arrS := strings.Split(line, " ")

	arr := make([]int, N)
	for i, s := range arrS {
		intFromString, err := strconv.Atoi(s)
		if err != nil {
			panic("something wrong")
		}
		arr[i] = intFromString

	}

	ans := minMovesForLine(
		arr, 
	)
	fmt.Printf("%d", ans)
}

func minMovesForLine(arr []int) int {
	N := len(arr)
	best := max(N-2, 0)
	for i := 0; i < N; i++ {
		for j := i+1; j < N; j++ {

			slope := (arr[j]- arr[i]) / (j - i)
			slopeCanBeUsed := (arr[j] - arr[i]) % (j - i) == 0
			if !slopeCanBeUsed {
				continue
			}
			start := arr[i] - slope * i
			totalMoves := 0
			for k, v := range arr {
				if start + slope * k != v {
					totalMoves++
					if totalMoves > best {
						break
					}
				}
			}
			best = min(best, totalMoves)
		}
	}
	return best
}