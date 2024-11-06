package main

type MyInt int

// 1. Fonction Add qui additionne deux entiers
func (i *MyInt) Add(a int) int {
	res := *i
	res = MyInt(a) + res
	return int(res)
}

func (i *MyInt) Sub(a int) int {
	return int(*i) - a
}

// 2. Fonction Multiply qui multiplie deux entiers
func (i *MyInt) Multiply(a int) int {
	return int(*i) * a
}

func (i *MyInt) Divide(b int) int {
	return int(*i) / b
}

// 3. Fonction Factorial qui calcule le factoriel d'un nombre
func (i *MyInt) Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * i.Factorial(n-1)
}
