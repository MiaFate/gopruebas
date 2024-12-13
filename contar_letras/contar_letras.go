// This package contains exercices
package contar_letras

import "unicode"

// Count letter for the parameter's phrase and returns an integer
func ContarLetras(frase string) int {
  var counter int
  for _, r := range frase{
    if unicode.IsLetter(r) {
      counter++
    }
  }
 return counter 
}
