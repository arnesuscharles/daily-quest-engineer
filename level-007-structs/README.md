# 🟢 Level 007 — Structs

## Quest Type
Main Quest

## Difficulty
🟢 Easy

---

## 📖 Story

The adventurer has learned to manage collections.

But a hero is more than a list of values.

A hero has:

a name,
a level,
titles,
experience.

To represent real-world entities,
the adventurer must learn Structs.

---

## 🎯 Objectives
- [x] Learn structs
- [x] Create custom types
- [x] Access struct fields
- [x] Update struct values
- [x] Use structs with functions
- [x] Understand data modeling
- [x] Document the journey
- [x] Push progress to Github

---

## 🧰 Developer Tips

#### VS Code — Quick Command Palette
Shortcut:
```text
Ctrl + Shift + P
```
Used for:
- find command,
- rename symbol,
- format document,
- restart Go language server,
- etc.

#### VS Code — Rename Symbol
Shortcut:
```text
F2
```
Rename struct, variable, or function in all project

#### VS Code — Go To Definition
Shortcut:
```text
F12
```
Jump to code definition

---

## Step-by-Step Walkthrough

### 📂 STEP 1 — Create Level Folder
move to repo:
```bash
cd ~/workspace/daily-quest-engineer
```
create new folder:
```bash
mkdir level-007-structs
cd level-007-structs
```
### 📂 STEP 2 — Initialize Go Module
```bash
go mod init level-007-structs
```

### 📂 STEP 3 — Create Project Structure
```bash
mkdir 01-basic-struct
mkdir 02-access-fields
mkdir 03-update-fields
mkdir 04-struct-with-function

touch README.md
```
open VS Code:
```bash
code .
```

### 📂 STEP 4 — Example 1: Basic Struct

Create file:
```bash
touch 01-basic-struct/main.go
```

fill main.go
```Go
package main

import "fmt"

type Character struct {
	Name  string
	Level int
	Title string
}

func main() {
	character := Character{
		Name:  "Charles",
		Level: 7,
		Title: "Collection Handler",
	}

	fmt.Println(character)
}
```

#### ▶️ Run Example 1
```bash
cd 01-basic-struct

go run .
```

Expected:
```
{Charles 7 Collection Handler}
```

Back to root level:
```bash
cd ..
```

#### 🧠 Concept Learned
Struct = custom data type.

We created new type:

```Go
type Character struct
```

that represents an entity.

### 📂 STEP 5 — Example 2: Access Fields

Create file:
```bash
touch 02-access-feilds/main.go
```

fill main.go
```Go
package main

import "fmt"

type Character struct {
	Name  string
	Level int
	Title string
}

func main() {

	character := Character{
		Name:  "Charles",
		Level: 7,
		Title: "Collection Handler",
	}

	fmt.Println("Name:", character.Name)
	fmt.Println("Level:", character.Level)
	fmt.Println("Title:", character.Title)
}
```

#### ▶️ Run Example 2
```bash
cd 02-access-fields

go run .
```

Expected:
```
Name: Charles
Level: 7
Title: Collection Handler
```

Back to root level:
```bash
cd ..
```

#### 🧠 Concept Learned
Access field using:

```Go
character.Name
```

is called: Dot Notation

### 📂 STEP 6 — Example 3: Update Fields

Create file:
```bash
touch 03-update-fields/main.go
```

fill main.go
```Go
package main

import "fmt"

type Character struct {
	Name  string
	Level int
	Title string
}

func main() {

	character := Character{
		Name:  "Charles",
		Level: 7,
		Title: "Collection Handler",
	}

	character.Level = 8
	character.Title = "Model Architect"

	fmt.Println(character)
}
```

#### ▶️ Run Example 3
```bash
cd 03-update-fields

go run .
```

Expected:
```
{Charles 8 Model Architect}
```

Back to root level:
```bash
cd ..
```

#### 🧠 Concept Learned
Field struct can be updated:

```Go
character.Level = 8
```

just like a normal variable.

### 📂 STEP 7 — Example 4: Struct With Function

Create file:
```bash
touch 04-struct-with-function/main.go
```

fill main.go
```Go
package main

import "fmt"

type Character struct {
	Name  string
	Level int
	Title string
}

func displayCharacter(character Character) {
	fmt.Println("Name:", character.Name)
	fmt.Println("Level:", character.Level)
	fmt.Println("Title:", character.Title)
}

func main() {

	character := Character{
		Name:  "Charles",
		Level: 7,
		Title: "Collection Handler",
	}

	displayCharacter(character)
}
```

#### ▶️ Run Example 4
```bash
cd 04-struct-with-function

go run .
```

Expected:
```
Name: Charles
Level: 7
Title: Collection Handler
```

Back to root level:
```bash
cd ..
```

#### 🧠 Concept Learned
Struct can be a function parameter.

This is foundation for:
- API Requests
- API Response
- Service Layer
- Database Models


#### 🧠 Important Concept
We have learned:

- variables
- conditions
- loops
- functions
- slices
- structs

means that we can start model data in a real world, and this is the foundation for almost all backend applications.

#### 🔥 Real Engineering Connection

Struct can be expanded into:

Backend
```Go
type User struct {}
type Product struct {}
type Order struct {}
```

Kubernetes
```Go
type Deployment struct {}
type Pod struct {}
```

Cloud
```Go
type EC2Instance struct {}
```

AI
```Go
type PromptRequest struct {}
```

## What I Learned
- Learned structs
- Created custom types
- Access struct fields
- Access struct values
- Use struct with functions
- Understand data modeling

## Rewards
| Attribute           | EXP |
| ------------------- | --- |
| 🖥️ Programming EXP | +30 |
| 📐 Architecture EXP | +15 |
| 📚 Wisdom EXP       | +10 |


## Title Unlocked

🟢 Model Architect