package main

import (
	"math/rand"
	"testing"
	"time"
)

// Génère un nom d'étudiant aléatoire
func generateRandomName() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	nameLength := rand.Intn(5) + 5 // Nom entre 5 et 10 caractères
	name := make([]rune, nameLength)
	for i := range name {
		name[i] = letters[rand.Intn(len(letters))]
	}
	return string(name)
}

// Génère une note aléatoire entre min et max
func generateRandomGrade(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func TestAddAndGetStudentWithRandomData(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Ajout d'étudiants aléatoires
	for i := 0; i < 10; i++ {
		name := generateRandomName()
		grade := generateRandomGrade(0, 100)
		AddStudent(name, grade)

		// Récupération de la note pour vérifier qu'elle a bien été ajoutée
		retrievedGrade, found := GetGrade(name)
		if !found {
			t.Errorf("GetGrade() did not find student %s, expected grade %.2f", name, grade)
		}
		if retrievedGrade != grade {
			t.Errorf("GetGrade() for student %s = %.2f; want %.2f", name, retrievedGrade, grade)
		}
	}
}

func TestRemoveStudentWithRandomData(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Ajout d'un étudiant pour tester la suppression
	name := generateRandomName()
	grade := generateRandomGrade(0, 100)
	AddStudent(name, grade)

	// Suppression de l'étudiant
	RemoveStudent(name)

	// Vérifie que l'étudiant n'est plus dans la map
	_, found := GetGrade(name)
	if found {
		t.Errorf("RemoveStudent() did not remove student %s", name)
	}
}

func TestAverageGradeWithRandomData(t *testing.T) {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Efface tous les étudiants précédents
	students = make(map[string]float64)

	// Ajoute des étudiants avec des notes aléatoires et calcule la moyenne attendue
	numStudents := 10
	total := 0.0
	for i := 0; i < numStudents; i++ {
		name := generateRandomName()
		grade := generateRandomGrade(0, 100)
		AddStudent(name, grade)
		total += grade
	}

	expectedAverage := total / float64(numStudents)
	calculatedAverage := AverageGrade()

	if calculatedAverage != expectedAverage {
		t.Errorf("AverageGrade() = %.2f; want %.2f", calculatedAverage, expectedAverage)
	}
}
