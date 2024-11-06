package main

import (
	"fmt"
)

// Créez une map globale pour stocker les étudiants et leurs notes
var students = make(map[string]float64)

// 1. Fonction pour ajouter un étudiant à la map
func AddStudent(name string, grade float64) {
	students[name] = grade
}

// 2. Fonction pour obtenir la note d'un étudiant
func GetGrade(name string) (float64, bool) {
	grade, found := students[name]
	if !found {
		fmt.Printf("L'étudiant %s n'a pas été trouvé.\n", name)
		return 0, false
	}
	return grade, true
}

// 3. Fonction pour supprimer un étudiant de la map
func RemoveStudent(name string) {
	if _, found := students[name]; found {
		delete(students, name)
		fmt.Printf("L'étudiant %s a été supprimé.\n", name)
	} else {
		fmt.Printf("L'étudiant %s n'a pas été trouvé.\n", name)
	}
}

// 4. Fonction pour calculer la moyenne des notes
func AverageGrade() float64 {
	if len(students) == 0 {
		return 0
	}
	total := 0.0
	for _, grade := range students {
		total += grade
	}
	return total / float64(len(students))
}

func main() {
	// Ajout d'étudiants
	AddStudent("Alice", 85.5)
	AddStudent("Bob", 92.0)
	AddStudent("Charlie", 78.3)

	// Afficher la note d'un étudiant
	grade, found := GetGrade("Alice")
	if found {
		fmt.Printf("Note de Alice: %.2f\n", grade)
	}

	// Supprimer un étudiant
	RemoveStudent("Bob")

	// Afficher la note moyenne de tous les étudiants
	avg := AverageGrade()
	fmt.Printf("Moyenne des notes: %.2f\n", avg)
}
