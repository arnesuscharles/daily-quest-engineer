# 🟢 Level 002 — Variables & Data Types

## Quest Type
Main Quest

## Difficulty
🟢 Easy

---

## 📖 Story

After unlocking the power of Golang,
the adventurer now seeks to understand
how machines store knowledge.

Numbers.
Words.
Truth.
Memory.

These are the foundations of all software systems.

---

## 🎯 Objectives
- [x] Create variables
- [x] Learn basic data types
- [x] Print variables
- [x] Modify variable values
- [x] Understand strings, integers, and booleans
- [x] Document the journey
- [x] Push progress to GitHub

---

## Step-by-Step Walkthrough

### 📂 STEP 1 — Create Level Folder
move to repo:
```bash
cd ~/workspace/daily-quest-engineer
```
create new folder:
```bash
mkdir level-002-variables-and-data-types
cd level-002-variables-and-data-types
```
### 📂 STEP 2 — Initialize Go Module
```bash
go mod init level-002-variables-and-data-types
```

#### 🧠 What is Variable?
Variable is a place to keep data inside compute memory. Just like, a box, it has name, and has a value.

example:
```text
name = "Charles"
```

- Computer saves "Charles" text, inside a box called name
### 📂 STEP 3 — Create main.go
```bash
touch main.go
```
open VS Code:
```bash
code .
```
fill main.go
```Go
package main

import "fmt"

func main() {

    // String
    var name string = "Charles"

    // Integer
    var level int = 2

    // Boolean
    var isLearning bool = true

    fmt.Println("Name:", name)
    fmt.Println("Level:", level)
    fmt.Println("Learning:", isLearning)
}
```

#### 🧠 Explanation
```
string
```
used for text.
example:
```text
"Hello"
"Charles"
"Golang"
```

```
int
```
used for integers
example:
```text
1
10
999
```

```
bool
```
Boolean only have 2 values:
```text
true
false
```
used for: condition, status, logic

### 📂 STEP 4 — Run Program
```bash
go run main.go
```
#### 🎯 Expected Output
```text
Name: Charles
Level: 2
Learning: true
```
#### 🧠 What is happening?
Program:
- Creates variable,
- save data to memory,
- get data,
- show data to terminal.


### 📂 STEP 5 — Modify Variables
update main.go to:
```Go
package main

import "fmt"

func main() {

    var name string = "Charles"

    fmt.Println("Before:", name)

    name = "Adventurer"

    fmt.Println("After:", name)
}
```

#### 🧠 What is happening?
Variable can:
- change
- updated
- save new value

This is called mutable state. 

### 📂 STEP 6 — Run Again
```bash
go run main.go
```

#### 🧠 Important Concept
Programming basically is: change data state to another state

example:
- login status,
- bank balance,
- shopping cart,
- game score,
- AI memory.

all of those come from:
- variables,
- state,
- data transformation

#### 🎯 Expected Output
```text
Before: Charles
After: Adventurer
```
## What I Learned
- How to create variables
- Learn basic data types
- How to print variables
- How to modify variable values
- Understand strings, integers, booleans

## Rewards
| Attribute           | EXP |
| ------------------- | --- |
| 🖥️ Programming EXP | +15 |
| ⚙️ Backend EXP      | +5  |
| 📚 Wisdom EXP       | +5  |


## Title Unlocked

🟢 Data Keeper