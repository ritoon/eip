package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Crée un pool qui stocke des strings
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Création d'un nouvel objet dans le pool")
			return "nouvel objet"
		},
	}

	// Place un objet dans le pool
	pool.Put("objet1")
	pool.Put("objet2")
	fmt.Println("Objet ajouté dans le pool")

	// Appel du GC pour simuler une collecte
	runtime.GC()

	// Tente de récupérer un objet dans le pool
	obj := pool.Get()
	fmt.Printf("Objet récupéré : %s\n", obj)

	obj = pool.Get()
	fmt.Printf("Objet récupéré : %s\n", obj)

	// Si le pool a été vidé par le GC, un nouvel objet est créé
	fmt.Printf("Objet récupéré après GC : %s\n", obj)
}
