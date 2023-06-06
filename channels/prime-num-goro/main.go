package main

import "fmt"

func generate(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		fmt.Printf("GENERATING %d to SRC Ch\n", i)
		ch <- i
	}
	close(ch)
}

func filter(src <-chan int, dest chan<- int, prime int) {
	for i := range src {
		//		fmt.Println("processing num -", i)
		fmt.Printf("%d -- is validated by goRounting-filter=%d\n", i, prime)
		if i%prime != 0 {
			fmt.Printf("published -%d to dest channel by goro-%d\n", i, prime)
			dest <- i
		}
	}
	close(dest)
}

func sieve(limit int) {
	ch := make(chan int)
	fmt.Printf("Type of `c`: %T\n", ch)
	fmt.Printf("Value of `c` is %v\n", ch)
	//fmt.Println(ch)
	go generate(limit, ch)
	for {
		fmt.Println("going to consume from <-ch")
		prime, ok := <-ch
		fmt.Println(prime, "--", ok)
		if !ok {
			break
		}
		ch1 := make(chan int)
		if prime != 2 {
			go filter(ch, ch1, prime)
			ch = ch1 //not a blocking statement
		}
		fmt.Println("going to print prime num=", prime)
		//fmt.Print()
	}
	fmt.Println()
}

func main() {
	sieve(10)
}
