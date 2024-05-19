package main

import(
	"fmt"
	"bufio"
	"os"
)

type element struct {
	val int
	id int
}

type printer struct {
	queue []element
}

func (p *printer) add(e element){
	p.queue = append(p.queue, e)
}

func (p *printer) print() element {
	toPrint := p.queue[0]
	idxOfToPrint := 0

	for i := 0; i < len(p.queue); i++ {
		if p.queue[i].val > toPrint.val {
			idxOfToPrint = i
			toPrint = p.queue[i]
		}
	}

	if len(p.queue) != 1{
		p.queue = append(p.queue[idxOfToPrint+1:], p.queue[:idxOfToPrint]...)
	} else {
		p.queue = make([]element,0)
	}
	return toPrint
}

func main() {

	var T, N, Q, num int 
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &T)

	for i := 0; i < T; i++ {
		fmt.Fscan(in, &N, &Q)
		p := printer{
			queue: make([]element, 0, N),
		}
		for j := 0; j < N; j++ {
			fmt.Fscan(in, &num)
			p.add(
				element{
					id:j,
					val:num,
				},
			)
		}
		counter := 0
		for {
			counter++
			res := p.print()
			if res.id == Q {
				fmt.Printf("%d\n", counter)
				break
			}
		}
	}
}