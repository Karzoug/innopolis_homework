package main

import (
	"fmt"
	"math/big"
)

func splitPrimesAndNotPrimes(arr []int) (prime, notPrime <-chan int) {
	primeCh := make(chan int)
	notPrimeCh := make(chan int)

	go func() {
		for _, v := range arr {
			if big.NewInt(int64(v)).ProbablyPrime(0) {
				primeCh <- v
			} else {
				notPrimeCh <- v
			}
		}
		close(primeCh)
		close(notPrimeCh)
	}()

	return primeCh, notPrimeCh
}

func main() {
	done := make(chan struct{}, 2)
	primeCh, notPrimeCh := splitPrimesAndNotPrimes([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	primeNumbers := make([]int, 0)
	notPrimeNumbers := make([]int, 0)
	go func() {
		for v := range primeCh {
			primeNumbers = append(primeNumbers, v)
		}
		done <- struct{}{}
	}()

	go func() {
		for v := range notPrimeCh {
			notPrimeNumbers = append(notPrimeNumbers, v)
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(primeNumbers)
	fmt.Println(notPrimeNumbers)
}
