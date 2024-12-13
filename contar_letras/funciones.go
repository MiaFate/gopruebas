package contar_letras

import "unicode"

// Count letter for the parameter's phrase
func ContarLetras(frase string) int {
  var counter int
  for _, r := range frase{
    if unicode.IsLetter(r) {
      counter++
    }
  }
 return counter 
}
