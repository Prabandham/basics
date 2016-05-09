package main

import "fmt"
import "sync"
import "time"

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for j := range jobs {
		defer wg.Done()
		// fmt.Println("Wroker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)    //We assume the lenght of the jobs channel is going to be 100.
	results := make(chan int, 100) //To accomodate the results of a 100 jobs we assign the capacity of the results channel to be 100 as well.
	wg := sync.WaitGroup{}

	for i := 0; i <= 9; i++ {
		go worker(i, jobs, results, &wg)
	}

	start_time := time.Now()
	for j := 0; j <= 20; j++ {
		wg.Add(1)
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	fmt.Println("Done processing all jobs in", time.Since(start_time))
	for result := range results {
		fmt.Println("Done with job and result is", result)
	}
}
