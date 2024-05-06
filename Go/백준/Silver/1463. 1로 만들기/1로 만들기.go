package main

import "fmt"


var Max = 100001

func main() {
	
	var input int 
	fmt.Scan(&input)
	cache := make([]int, input+1, input+1)

	for i:= range cache {
		cache[i] = Max 
	}

	fmt.Println(reduce(input, cache))
}

func reduce(i int, cache []int) int {

	if cache[i] != Max {
		return cache[i]
	}
	if i == 1 {
		return 0
	}

	if i % 2 == 0 {
		cache[i] = min(cache[i], reduce(i/2, cache)+1)
	}
	if i % 3 == 0 {
		cache[i] = min(cache[i], reduce(i/3, cache)+1)
	}
	cache[i] = min(cache[i], reduce(i-1, cache)+1)
	return cache[i]
}