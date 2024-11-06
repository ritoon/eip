package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1. Fonction worker qui effectue une opération simple sur une tâche
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d traite la tâche %d\n", id, task)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simule un traitement
		result := task * 2
		results <- result
	}
}

// 2. Fonction startWorkerPool qui initialise et démarre le pool de workers
func startWorkerPool(numWorkers int, tasks <-chan int, results chan<- int) {
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
}

// 3. Fonction distributeTasks qui envoie des tâches dans le canal de tâches
func distributeTasks(numTasks int, tasks chan<- int) {
	for i := 1; i <= numTasks; i++ {
		tasks <- i
		fmt.Printf("Tâche %d envoyée\n", i)
	}
	close(tasks)
}

func main() {
	// Initialisation du générateur de nombres aléatoires
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Création des canaux de tâches et de résultats
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// Initialisation et démarrage du pool de workers
	numWorkers := 10
	numTasks := 100
	go startWorkerPool(numWorkers, tasks, results)

	// Distribution des tâches
	go distributeTasks(numTasks, tasks)

	// Récupération des résultats
	for result := range results {
		fmt.Printf("Résultat reçu: %d\n", result)
	}
}
