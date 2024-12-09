package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// EuclideanGCD calculates the greatest common divisor of two integers.
func EuclideanGCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Worker estimates the count of coprime pairs in a subset of iterations.
func Worker(iterations int, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < iterations; i++ {
		x := randGen.Int63n(1e18) + 1 // Generate a random number in the range [1, 10^18]
		y := randGen.Int63n(1e18) + 1
		if EuclideanGCD(x, y) == 1 {
			count++
		}
	}
	results <- count
}

// EstimatePi estimates the value of Pi using a Monte Carlo method with concurrency.
func EstimatePi(totalIterations, workers int) float64 {
	iterationsPerWorker := totalIterations / workers
	results := make(chan int, workers)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go Worker(iterationsPerWorker, results, &wg)
	}

	// Wait for all workers to complete and close the channel.
	go func() {
		wg.Wait()
		close(results)
	}()

	totalCoprimeCount := 0
	for count := range results {
		totalCoprimeCount += count
	}

	probability := float64(totalCoprimeCount) / float64(totalIterations)
	if probability > 0 {
		return math.Sqrt(6 / probability)
	}
	return math.Inf(1)
}

func main() {
	var sampleSize, workers int
	fmt.Print("Enter the number of operations: ")
	fmt.Scan(&sampleSize)
	fmt.Print("Enter the number of workers: ")
	fmt.Scan(&workers)

	if workers <= 0 {
		fmt.Println("Number of workers must be greater than 0.")
		return
	}

	estimatedPi := EstimatePi(sampleSize, workers)

	fmt.Printf("Estimated value of π: %.10f\n", estimatedPi)
	fmt.Printf("Difference: %.10f\n", math.Abs(estimatedPi-math.Pi))
}

// Enter the number of operations: 1000000000
// Enter the number of workers: 12
// Estimated value of π: 3.1416122494
// Difference: 0.0000195958
// Aproksymacja liczby π w Twoim algorytmie nie powinna być stosowana
// w kontekście kryptografii z powodu kilku istotnych ograniczeń matematycznych i statystycznych:
//  - Brak deterministyczności
//  - Zależność od generatora liczb losowych
//  - Statystyczna natura metody Monte Carlo
//  - Zależność od wielkości próby
