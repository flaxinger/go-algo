package main

import(
	"fmt"
	"bufio"
	"os"
)

const (
	Cheese = 1
	Air = -1
	InnerAir = 0
)

var (
	dx = []int{1, 0, -1, 0}
	dy = []int{0, 1, 0, -1}
)

type point struct {
	X int
	Y int
}

func main() {
	var N, M int 
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &M)

	board := make([][]int, N)

	for i := 0; i < N; i++ {
		board[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscan(in, &board[i][j])
		}
	}
	fmt.Println(solve(board, N, M))
}

func solve(board [][]int, n, m int) int {
	dayCount := 0
	q := make([]point, 0, n*m)
	dfs(board, 0, 0)

	for {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if board[i][j] == Cheese {
					count := 0
					for k := 0; k < 4; k++ {
						nx := i + dx[k]
						ny := j + dy[k]
						if inValidRange(nx, ny, n, m) && board[nx][ny] == Air {
							count++
						}
					}
					if count >= 2 {
						board[i][j] = InnerAir
						q = append(q, point{
							X: i,
							Y: j,
						})
					}
				}
			}
		}

		if len(q) == 0{
			break
		}
		
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			dfs(board, cur.X, cur.Y)
		}
		dayCount++
	}
	return dayCount
}

func inValidRange(x, y, xMax, yMax int) bool{
	if x >= 0 && y >= 0 && x < xMax && y < yMax {
		return true
	} else {
		return false
	}
}

func dfs(board [][]int, x, y int) {
	if board[x][y] == InnerAir {
		board[x][y] = Air
	} else {
		return
	}

	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		if inValidRange(nx, ny, len(board), len(board[0])){
			dfs(board, nx, ny)
		}
	}
}