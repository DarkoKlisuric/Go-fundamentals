package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int)  {
	for {
		select {
		case v := <-e:
			fmt.Println("from the eve channel: ", v)

		case v := <-o:
			fmt.Println("from the odd channel: ", v)

		case v := <-q:
			fmt.Println("from the quit channel: ", v)
			return
		}
	}
}

/*
from the eve channel:  0
from the odd channel:  1
from the eve channel:  2
from the odd channel:  3
from the eve channel:  4
from the odd channel:  5
from the eve channel:  6
from the odd channel:  7
from the eve channel:  8
from the odd channel:  9
from the eve channel:  10
from the odd channel:  11
from the eve channel:  12
from the odd channel:  13
from the eve channel:  14
from the odd channel:  15
from the eve channel:  16
from the odd channel:  17
from the eve channel:  18
from the odd channel:  19
from the eve channel:  20
from the odd channel:  21
from the eve channel:  22
from the odd channel:  23
from the eve channel:  24
from the odd channel:  25
from the eve channel:  26
from the odd channel:  27
from the eve channel:  28
from the odd channel:  29
from the eve channel:  30
from the odd channel:  31
from the eve channel:  32
from the odd channel:  33
from the eve channel:  34
from the odd channel:  35
from the eve channel:  36
from the odd channel:  37
from the eve channel:  38
from the odd channel:  39
from the eve channel:  40
from the odd channel:  41
from the eve channel:  42
from the odd channel:  43
from the eve channel:  44
from the odd channel:  45
from the eve channel:  46
from the odd channel:  47
from the eve channel:  48
from the odd channel:  49
from the eve channel:  50
from the odd channel:  51
from the eve channel:  52
from the odd channel:  53
from the eve channel:  54
from the odd channel:  55
from the eve channel:  56
from the odd channel:  57
from the eve channel:  58
from the odd channel:  59
from the eve channel:  60
from the odd channel:  61
from the eve channel:  62
from the odd channel:  63
from the eve channel:  64
from the odd channel:  65
from the eve channel:  66
from the odd channel:  67
from the eve channel:  68
from the odd channel:  69
from the eve channel:  70
from the odd channel:  71
from the eve channel:  72
from the odd channel:  73
from the eve channel:  74
from the odd channel:  75
from the eve channel:  76
from the odd channel:  77
from the eve channel:  78
from the odd channel:  79
from the eve channel:  80
from the odd channel:  81
from the eve channel:  82
from the odd channel:  83
from the eve channel:  84
from the odd channel:  85
from the eve channel:  86
from the odd channel:  87
from the eve channel:  88
from the odd channel:  89
from the eve channel:  90
from the odd channel:  91
from the eve channel:  92
from the odd channel:  93
from the eve channel:  94
from the odd channel:  95
from the eve channel:  96
from the odd channel:  97
from the eve channel:  98
from the odd channel:  99
from the quit channel:  0
about to exit

 */