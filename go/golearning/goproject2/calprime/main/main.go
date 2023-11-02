package main

import "math"

func putNum(initChan chan uint64, max uint64) {
	for i := uint64(2); i <= max; i++ {
		initChan <- i
	}
	close(initChan)
}

// from to exitFlag
func primeNum(initChan chan uint64, primeChan chan uint64, exitChan chan bool) {
	for {
		num, ok := <-initChan
		if !ok {
			break
		}
		var i uint64
		for i = 2; i < num; i++ {
			if num%i == 0 {
				break
			}
		}
		if i == num {
			primeChan <- num
		}

	}
	exitChan <- true
}

// 只计算到sqrt(num)
func primeNum1(initChan chan uint64, primeChan chan uint64, exitChan chan bool) {
	for {
		num, ok := <-initChan
		if !ok {
			break
		}
		i, sqrt := uint64(2), uint64(math.Sqrt(float64(num)))
		for ; i <= sqrt; i++ {
			if num%i == 0 {
				break
			}
		}
		if i > sqrt {
			primeChan <- num
		}

	}
	exitChan <- true
}

// for range test
func primeNum2(initChan chan uint64, primeChan chan uint64, exitChan chan bool) {
	for num, ok := <-initChan; ok; num, ok = <-initChan {
		i, sqrt := uint64(2), uint64(math.Sqrt(float64(num)))
		for ; i <= sqrt; i++ {
			if num%i == 0 {
				break
			}
		}
		if i > sqrt {
			primeChan <- num
		}
	}
	exitChan <- true
}

func main() {
	initChan := make(chan uint64, 1000)
	primeChan := make(chan uint64, 2000)
	exitChan := make(chan bool, 4)

	go putNum(initChan, 8000)

	// 开启若干写成计算素数
	for i := 0; i < 4; i++ {
		go primeNum(initChan, primeChan, exitChan)
	}

	// for i := 0; i < 4; i++ {
	// 	<-exitChan
	// }
	// close(primeChan)
	// for num := range primeChan {
	// 	println(num)
	// }

	// 不要让主线程等待
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	// for num, ok := <-primeChan; ok; num, ok = <-primeChan {
	// 	println(num)
	// }

	for {
		num, ok := <-primeChan
		if !ok {
			break
		}
		println(num)
	}

}
