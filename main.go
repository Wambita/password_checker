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

    if length >= 10 && hasLetter && hasNumber && hasSpecial {
    return "Strong"
  } else if length >= 8 && hasLetter && hasNumber {
    return "Medium"
  } 
  return "Weak"
  }

func main (){
var password string
fmt.Print("Enter password: ")
fmt.Scanln(&password)
strength := checkPasswordStrength(password)
fmt.Println("Password strenth is:", strength)  
  
}
