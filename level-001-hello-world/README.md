# 🟢 Level 001 — Hello Golang

## Quest Type
Main Quest

## Difficulty
🟢 Easy

---

# 📖 Story

The adventurer has learned how to communicate with machines using code.

A new weapon has been unlocked:
Golang.

---

# 🎯 Objectives

- [x] Install Golang
- [x] Verify Go installation
- [x] Create first Go module
- [x] Create `main.go`
- [x] Run first Go program

---

# 🗺️ Step-by-Step Walkthrough

## STEP 1 — Install Golang

```bash
sudo apt update
sudo apt install golang-go -y
```
### Explanation
```
apt update
```
refreshes Linux package information.
```
apt install
```
installs software from Ubuntu repositories.
```
golang-go
```
is the Golang compiler package.
```
-y
```
automatically answers "yes" during installation.

## STEP 2 — Verify Installation 
```bash
go version
```
### Expected Output
```text
go version go1.22.12 linux/amd64
```
### Explanation

This command checks:

- whether Go is installed,
- which version is installed,
- and system architecture.

## STEP 3 — Create Project Folder
```bash
mkdir level-001-hello-world
cd level-001-hello-world
```

### Explanation
```
mkdir
```
creates a new directory.
```
cd
```
moves into the directory.

## STEP 4 — Initialize Go Module
```bash
go mod init level-001-hello-world
```
### Explanation
This creates:
```
go.mod
```

The go.mod file is:

- the identity of the Go project,
- dependency manager,
- project configuration file.

## STEP 5 — Create main.go
```bash
touch main.go
```

### Explanation
Creates an empty Go source file.

## STEP 6 — Write First Program
```Go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Adventurer!")
}
```

### Explanation
```
package main
```
Defines this as an executable application.
```
import "fmt"
```
Imports Go's formatting package.
```
func main()
```
Program entry point. Execution always starts here.
```
fmt.Println
```
Prints text to terminal output.

## STEP 7 — Run Program
```bash
go run main.go
```

### Output
```
Hello, Adventurer!
```

### Explanation
```
go run
```
compiles and runs the program instantly. 
The computer is now executing instructions written by the adventurer.
---


# 🧠 What I Learned
- How to install Golang
- What Go modules are
- How Go programs execute
- What main() means
- Basic terminal workflow
- Basic compiler workflow

# 🏆 Rewards
Attribute	EXP
🖥️ Programming EXP	+10
⚙️ Backend EXP	+5
📚 Wisdom EXP	+5

# 🏅 Title Unlocked
🟢 Code Initiate

# Unlocks
- Go Compiler
- Go Modules
- Program Entry Point
- CLI Program Execution