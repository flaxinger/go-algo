package main

import(
	"fmt"
	"slices"
)



func main() {
	
	var N, M, S int 

	fmt.Scan(&N)
	fmt.Scan(&M)
	fmt.Scan(&S)

	board := initBoard(N, M, S)
	vsD, vsB := make([]bool, N+1), make([]bool, N+1)
	
	dfs(S, board, vsD)
	fmt.Println()
	bfs(S, board, vsB)
}

func initBoard(n, m, s int) [][]int {
	board := make([][]int, n+1)

	for i := 0; i < m; i++ {
		var a, b int

		fmt.Scan(&a)
		fmt.Scan(&b)

		board[a] = append(board[a], b)
		board[b] = append(board[b], a)
	}

	for i:= range board {
		slices.Sort(board[i])
	}
	return board
}

func dfs(s int, board [][]int, vs []bool) {
	vs[s] = true
	fmt.Printf("%d ", s)

	for _, node := range board[s] {
		if !vs[node] {
			dfs(node, board, vs)
		}
	}
}

func bfs(s int, board [][]int, vs []bool) {	

	q := make([]int, 0)
	q = append(q, s)
	vs[s] = true

	for len(q)> 0{
		cur := q[0]
		q = q[1:]

		fmt.Printf("%d ", cur)
		for _, v := range board[cur] {
			if !vs[v] {
				vs[v] = true
				q = append(q, v)
			}
		}

	}

}