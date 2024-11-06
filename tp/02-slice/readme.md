# Exercice : Manipulation de Slices

## Objectif

L'objectif de cet exercice est de comprendre comment manipuler et travailler avec des slices en Go.

## Énoncé

Créez une fonction appelée FilterEvenNumbers qui prend un slice d'entiers en entrée et retourne un nouveau slice contenant uniquement les nombres pairs.

## Prototype :

``go
func FilterEvenNumbers(numbers []int) []int
``
Créez une fonction appelée SumSlice qui prend un slice d'entiers et retourne la somme de tous les éléments.

``go
func SumSlice(numbers []int) int
``
Créez une fonction RemoveDuplicates qui prend un slice d'entiers et retourne un nouveau slice sans les doublons.

``go
func RemoveDuplicates(numbers []int) []int
``

## Consignes

Testez chaque fonction avec plusieurs slices pour vérifier qu'elles fonctionnent correctement.

Dans la fonction principale (main), créez un slice d'entiers avec des nombres de votre choix, incluant des doublons et des nombres impairs, et utilisez vos fonctions pour :

Filtrer les nombres pairs
Calculer la somme du slice initial
Créer un slice sans doublons
Affichez les résultats des différentes opérations dans la console.

Exemple de Code
Voici un exemple de code pour vous donner un point de départ :

Résultats attendus
Par exemple, avec le slice numbers := []int{1, 2, 3, 4, 4, 5, 6, 6, 7, 8, 9, 10}, on pourrait obtenir :

Nombres pairs : [2, 4, 4, 6, 6, 8, 10]
Somme des nombres : 65
Nombres uniques : [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
Objectif pédagogique

Cet exercice permet de pratiquer :

- La création et manipulation de slices en Go
- Les boucles et conditions pour filtrer et transformer des slices
- Les fonctions en tant qu’outils de réutilisation du code