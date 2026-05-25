# 🟢 Level 005 — Functions

## Quest Type
Main Quest

## Difficulty
🟢 Easy

---

## 📖 Story

The adventurer has learned repetition.

But repeating code is inefficient.

Wise engineers forge reusable abilities:
small functions that can be called anytime.

Today,
the adventurer learns modular power.

---

## 🎯 Objectives
- [x] Learn functions
- [x] Learn parameters
- [x] Learn return values
- [x] Create reusable logic
- [x] Understand modular programming
- [x] Document the journey
- [x] Push progress to Github

---

## Step-by-Step Walkthrough

### 📂 STEP 1 — Create Level Folder
move to repo:
```bash
cd ~/workspace/daily-quest-engineer
```
create new folder:
```bash
mkdir level-005-functions
cd level-005-functions
```
### 📂 STEP 2 — Initialize Go Module
```bash
go mod init level-005-functions
```

#### 🧠 What is Function?
Function is: a group of code that has a specific task and can be recalled.

Imagine: skill, spell, ability.

Rather than rewrite codes everytime, we can just: call function

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

func greet() {
    fmt.Println("Welcome, Adventurer!")
}

func main() {
    greet()
}
```

#### 🧠 Explanation
```text
func greet()
```
create function called 
```text
greet
```


### 📂 STEP 4 — Run Program
```bash
go run main.go
```
#### 🎯 Expected Output
```text
Welcome, Adventurer!
```

### 📂 STEP 5 — Function with Parameters
update main.go to:
```Go
package main

import "fmt"

func greet(name string) {
    fmt.Println("Welcome,", name)
}

func main() {
    greet("Charles")
    greet("Adventurer")
}
```

#### 🧠 What is Parameter?
Parameter is: data that is sent to a function.

#### 🧠 So a Function can be flexible
Instead of hardcoded:
```text
Welcome, Charles
```

we can: send different data

### 📂 STEP 6 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
Welcome, Charles
Welcome, Adventurer
```

### 📂 STEP 7 — Function with Return Value
update main.go to:
```Go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {

    result := add(5, 3)

    fmt.Println("Result:", result)
}
```

#### 🧠 New Concept
```text
return
```

used for: return the result of a function.

#### 🧠 This Function Now
- receive input,
- process data,
- produce output.

and that are foundation of:
- APIs,
- backend logic,
- services,
- data processing.

### 📂 STEP 8 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
Result: 8
```

#### 🧠 Important Concept
Function is the foundation of modular programming.

Without functions:
- code will be longer,
- hard to read,
- hard to maintain.

With functions:
- reusable,
- cleaner,
- scalable,
- easier to debug.

#### 🧠 Real Engineering Connection
Function can be expanded to:
- REST handlers,
- services,
- repositories,
- middleware,
- ETL transformations,
- ML pipelines,
- cloud functions,
- microservices.

All come from: reusable logic blocks.


## What I Learned
- Learned functions
- Learned parameters
- Learned return values
- Create reusable logic
- Understand modular programming

## Rewards
| Attribute           | EXP |
| ------------------- | --- |
| 🖥️ Programming EXP | +25 |
| 📐 Architecture EXP | +10  |
| 📚 Wisdom EXP       | +5  |


## Title Unlocked

🟢 Function Smith