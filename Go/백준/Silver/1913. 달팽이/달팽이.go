package main

import(
	"fmt"
	"bufio"
	"os"
)


var (
	dx = []int{0, 1, 0, -1}
	dy = []int{1, 0, -1, 0}
)

func main() {
	var N, M, X, Y int
	in := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &N, &M)

	board := make([][]int, N, N)
	x := N/2
	y := N/2
	X = x
	Y = y
	n := 1
	for i := 0; i < N; i++ {
		board[i] = make([]int, N, N)
	}
	board[x][y] = n
	for i := 1; i <= N/2; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < i*2; k++ {
				n++
				if k == 0 && j == 0 {
					x--
					board[x][y] = n
				} else {
					x +=dx[j]
					y +=dy[j]
					board[x][y] = n
				}
				if n == M {
					X = x
					Y = y
				}
			}
		}
	}	

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j == N-1 {
				fmt.Fprintf(w, "%d\n", board[i][j])
			} else {
				fmt.Fprintf(w, "%d ", board[i][j])
			}
		}
	}

	fmt.Fprintf(w, "%d %d", X+1, Y+1)
	w.Flush()
}