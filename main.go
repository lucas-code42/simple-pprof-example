package main

import (
	"os"
)

func main() {
	if len(os.Args) != 3 {
		panic("argumentos invalidos")
	}

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(os.Args[2], b, 0666)
	if err != nil {
		panic(err)
	}
}
