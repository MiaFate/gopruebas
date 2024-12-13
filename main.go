package main

import (
	"fmt"

	"github.com/miafate/gopruebas/contar_letras"
)

func main()  {
  letras := contar_letras.ContarLetras("hola amikos") 
  fmt.Println(letras)
}
