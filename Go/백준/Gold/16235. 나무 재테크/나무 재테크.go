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
	dx = []int{1, 1, 1, 0, 0, -1, -1, -1}
	dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}
	N int
)

type land struct {
	nutrition        int
	nutritionPerIter int
	trees            []int
}

func (l *land) addTree(i int) {
	if len(l.trees) > 0 {
		l.trees = append(l.trees[:1], l.trees[:]...)
		l.trees[0] = i
	} else {
		l.trees = append(l.trees, i)
	}
}

func (l *land) killTrees(idx int) {
	for i := idx; i < len(l.trees); i++ {
		l.nutrition += l.trees[i]/2
	}
	l.trees = l.trees[:idx]
}

func (l *land) reduceNutrition(n int){
	l.nutrition -= n
}


func (l *land) addNutrition(){
	l.nutrition += l.nutritionPerIter
}

func main() {
	var M, K, x, y, z int 
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N, &M, &K)
	board := make([][]land, N+1)

	for i := 1; i <= N; i++ {
		board[i] = make([]land, N+1)
		for j := 1; j <= N; j++ {
			board[i][j] = land{
				trees: make([]int, 0),
				nutrition: 5,
			}
			fmt.Fscan(in, &board[i][j].nutritionPerIter)
		}
	}

	for i := 0; i < M; i++ {
		fmt.Fscan(in, &x, &y, &z)
		board[x][y].addTree(z)
	}

	fmt.Println(solve(board, K))
}


func solve(board [][]land, k int) int {
	// printBoard(board)
	for i := 0; i < k; i++ {
		for x := 1; x <= N; x++{
			for y := 1; y <= N; y++ {
				cur := &board[x][y]
				for idx, tree := range cur.trees {
					if tree <= cur.nutrition {
						cur.reduceNutrition(tree)
						cur.trees[idx]++
					} else{
						cur.killTrees(idx)
						break
					}
				}
				cur.addNutrition()
			}
		}
		for x := 1; x <= N; x++{
			for y := 1; y <= N; y++ {
				cur := &board[x][y]
				for _, tree := range cur.trees {
					if tree % 5 == 0 {
						for k := 0; k < 8; k++ {
							nx := dx[k] + x
							ny := dy[k] + y
							if nx > 0 && nx <= N && ny > 0 && ny <= N {
								board[nx][ny].addTree(1)
							}
						}
					}
				}
			}
		}
		// printBoard(board)
	}	
	ans := 0
	for x := 1; x <= N; x++ {
		for y := 1; y <= N; y++ {
			ans += len(board[x][y].trees)
		}
	}
	return ans
}

func printBoard(board [][]land){
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			cur := board[i][j]
			fmt.Printf("(%d, %d)\t", cur.nutrition, len(cur.trees))
		}
		fmt.Println()
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			cur := board[i][j]
			if len(cur.trees) > 0 {
				fmt.Printf("(%d, %d) ",i, j)
			}
			for _, tree := range cur.trees {
				fmt.Printf("%d ", tree)
			}
			if len(cur.trees) > 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println("-------------")

}