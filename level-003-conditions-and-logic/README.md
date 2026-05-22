# 🟢 Level 003 — Conditions & Logic

## Quest Type
Main Quest

## Difficulty
🟢 Easy

---

## 📖 Story

The adventurer has learned how to store knowledge.

Now the adventurer must learn: judgment.

Machines do not think like humans.
They follow logic.

True.
False.
Yes.
No.

From these simple rules,
entire software systems are born.

---

## 🎯 Objectives
- [x] Learn if
- [x] Learn else
- [x] Learn comparison operators
- [x] Learn logical flow
- [x] Create simple decision-making program
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
mkdir level-003-conditions-and-logic
cd level-003-conditions-and-logic
```
### 📂 STEP 2 — Initialize Go Module
```bash
go mod init level-003-conditions-and-logic
```

#### 🧠 What is Condition?
Condition is a statement that only resulting in Yes or No.

example:
```text
is Level > 10?
```

The answer is only Yes or No
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

    level := 3

    if level >= 3 {
        fmt.Println("New title unlocked!")
    }
}
```

#### 🧠 Explanation
```
if
```
is used to run code if only the condition return true

#### 🧠 Comparison Operator
In this quest:
```text
>=
```
means more than or equal

So if:
```
level >= 3
```
is true, then
```text
fmt.Println(...)
```
will be executed


### 📂 STEP 4 — Run Program
```bash
go run main.go
```
#### 🎯 Expected Output
```text
New title unlocked!
```


### 📂 STEP 5 — Learn Else
update main.go to:
```Go
package main

import "fmt"

func main() {

    level := 1

    if level >= 3 {
        fmt.Println("New title unlocked!")
    } else {
        fmt.Println("Level too low.")
    }
}
```

#### 🧠 What is else?
else executed if condition return false

#### 🧠 Logic Program
if
```text
level >=3
```
FALSE, then
```Go
else
```
will be executed

### 📂 STEP 6 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
Level too low.
```

### 📂 STEP 7 — Learn Multiple Conditions
update main.go to:
```Go
package main

import "fmt"

func main() {

    level := 7

    if level >= 10 {
        fmt.Println("Elite Adventurer")
    } else if level >= 5 {
        fmt.Println("Intermediate Adventurer")
    } else {
        fmt.Println("Beginner Adventurer")
    }
}
```

#### 🧠 What is else if?
is used to check additional condition

#### 🧠 Flow Logic
Program reads from top to bottom. If the first condition return false, the program continue to the next condition.

### 📂 STEP 8 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
Intermediate Adventurer
```

#### 🧠 Important Concept
Programming logic basically is: decision trees.

example:
- login success/fail,
- payment approved/rejected,
- AI confidence checks,
- Kubernetes health checks,
- Cloud auto scaling,
- fraud detection.

all of those come from:
- conditions,
- comparisons,
- logical flow


## What I Learned
- Learned if
- Learned else
- Learned else if
- Learned comparison operators
- Learned logical flow
- Created simple decision-making program 

## Rewards
| Attribute           | EXP |
| ------------------- | --- |
| 🖥️ Programming EXP | +15 |
| 🧩 Debugging EXP | +10  |
| 📚 Wisdom EXP       | +5  |


## Title Unlocked

🟢 Logic Apprentice
