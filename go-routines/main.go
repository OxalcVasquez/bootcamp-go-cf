package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simula una tarea pesada
func process(id int, job int) {
	delay := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(delay))
	fmt.Printf("Worker %d procesó trabajo %d (tardó %dms)\n", id, job, delay)
}

func worker(id int, jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		process(id, job)
	}
	done <- true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	done := make(chan bool)

	// Lanzar workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, done)
	}

	// Enviar trabajos
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Ya no habrá más trabajos

	// Esperar que todos terminen
	for w := 1; w <= numWorkers; w++ {
		<-done
	}

	fmt.Println("🎉 Todos los trabajos han sido procesados.")
}
