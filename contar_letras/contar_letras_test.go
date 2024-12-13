package contar_letras_test

import (
	"fmt"

	"github.com/miafate/gopruebas/contar_letras"
)

func ExampleContarLetras() {
  letters := contar_letras.ContarLetras("frase de prueba")
  fmt.Println(letters)
  // Output:
  // 13
}
