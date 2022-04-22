package main

import "fmt"

func main() {
    ch := make(chan int, 2)
	//work well
	ch <- 1
	ch <- 2   
	ch <- 3
  
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//dead lock
	// ch <- 1
	// ch <- 2   
	// // ch <- 3
  
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch) //<-here
	//dead lock
	// ch <- 1
	// ch <- 2   
	// ch <- 3 //<-here
  
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// // fmt.Println(<-ch)

}