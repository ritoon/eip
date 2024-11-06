package main

import (
	"math/rand"
	"testing"
	"time"
)

// Test aléatoire pour le pool de workers
func TestWorkerPoolWithRandomData(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	// Génération aléatoire de tâches
	numTasks := 10
	numWorkers := 5
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// Création d'un slice pour stocker les tâches et les résultats attendus
	taskValues := make([]int, numTasks)
	expectedResults := make([]int, numTasks)

	for i := 0; i < numTasks; i++ {
		taskValues[i] = rand.Intn(100) // Génère une tâche aléatoire entre 0 et 99
		expectedResults[i] = taskValues[i] * 2
	}

	// Lancement du pool de workers
	go startWorkerPool(numWorkers, tasks, results)

	// Distribution des tâches
	go func() {
		for _, task := range taskValues {
			tasks <- task
		}
		close(tasks)
	}()

	// Lecture des résultats et vérification
	actualResults := make([]int, 0, numTasks)
	for result := range results {
		actualResults = append(actualResults, result)
	}

	// Vérifie que les résultats sont comme attendus
	for _, expected := range expectedResults {
		found := false
		for _, actual := range actualResults {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Résultat attendu %d non trouvé dans les résultats actuels %v", expected, actualResults)
		}
	}
}
