package contar_letras

import "unicode"

func ContarLetras(frase string) int {
  var counter int
  for _, r := range frase{
    if unicode.IsLetter(r) {
      counter++
    }
  }
 return counter 
}
