# Password Strength Checker 

A simple **CLI password strength checker written in Go** that evaluates how strong a password is based on its length and character composition.

This project was built as a **small Go utility** to practice:

* CLI programs in Go
* String and character analysis
* Writing simple security-related tools
* Building consistent GitHub contributions

---

## Features

* Checks password strength directly from the terminal
* Detects:

  * Letters
  * Numbers
  * Special characters
* Categorizes passwords into:

  * Weak 
  * Medium 
  * Strong 

---

## How It Works

The program evaluates passwords using simple rules:

| Condition                                                 | Result |
| --------------------------------------------------------- | ------ |
| Length < 8                                                | Weak   |
| Length ≥ 8 and contains letters and numbers               | Medium |
| Length ≥ 10 with letters, numbers, and special characters | Strong |

---

## Project Structure

```
password-strength-checker
│
├── main.go
└── README.md
```

---

## Installation

Clone the repository:

```bash
git clone https://github.com/Wambita/password-strength-checker.git
cd password-strength-checker
```

---

## Run the Program

```bash
go run main.go
```

---

## Example Usage

```
Enter password: mypassword
Password strength: Medium 
```

```
Enter password: P@ssw0rd123
Password strength: Strong 
```

---

## Example Code

The program analyzes each character to determine if it is:

* a letter
* a number
* a special character

and then determines the strength level accordingly.

---

## Future Improvements

Possible improvements to this tool:

* Hide password input while typing
* Add entropy-based strength calculation
* Generate password suggestions
* Turn it into a reusable Go package
* Add unit tests

---

## Author

**Sheila Fana**

GitHub:
[https://github.com/Wambita](https://github.com/Wambita)
---

## License

MIT License

