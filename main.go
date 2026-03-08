package main

import (
  "fmt"
  "unicode"
)

func checkPasswordStrength (password string) string {
  length := len(password)
  hasLetter := false
  hasNumber := false
  hasSpecial := false

  for _,char := range password {
    switch {
      case unicode.IsLetter(char):
        hasLetter = true
      case unicode.IsDigit(char):
        hasNumber = true
      case unicode.IsPunct(char) || unicode.IsSymbol(char):
        hasSpecial = true
    }
    }

  
  
  
}
