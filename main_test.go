package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// injetando os argumentos para o programa funcionar
	os.Args = []string{"gp", "foo.txt", "bar.txt"}
	main()
}
