package main

import (
	"bufio"
	"fmt"
	"os"
)

type Date struct {
	day, month, year int
}

func (d Date) Less(other Date) bool {
	if d.year != other.year {
		return d.year < other.year
	}
	if d.month != other.month {
		return d.month < other.month
	}
	return d.day < other.day
}

func main() {
	var N int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &N)

	minDate := Date{1, 1, 2011}
	maxDate := Date{1, 1, 1980}
	var minName, maxName string

	for i := 0; i < N; i++ {
		var name string
		var day, month, year int
		fmt.Fscan(in, &name, &day, &month, &year)

		cur := Date{day, month, year}
		if cur.Less(minDate) {
			minDate = cur
			minName = name
		}
		if maxDate.Less(cur) {
			maxDate = cur
			maxName = name
		}
	}
	fmt.Println(maxName)
	fmt.Println(minName)
}
