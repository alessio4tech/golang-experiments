package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 4, 6, 8, 10}

	c := readNumbers(a...)
	d := readNumbers(b...)

	e := make(<-chan int)

	e = merge(c, d)

	for stream := range e {

		fmt.Println(stream)
	}

}

func merge(a <-chan int, b <-chan int) <-chan int {

	e := make(chan int)

	go func() {

		defer close(e)

		for a != nil && b != nil {

			select {

			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				e <- v
				fmt.Println("branch a ", v)
			case v, ok := <-b:

				if !ok {

					b = nil
					continue
				}
				e <- v
				fmt.Println("branch b ", v)
			}

		}
	}()

	return e

}

func readNumbers(vs ...int) <-chan int {

	output := make(chan int)

	go func() {
		defer close(output)
		for _, v := range vs {
			output <- v
			//	time.Sleep(time.Duration(time.Millisecond * 1000))
		}
	}()

	return output

}
