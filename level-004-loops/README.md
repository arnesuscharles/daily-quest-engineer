# 🟢 Level 004 — Conditions & Logic

## Quest Type
Main Quest

## Difficulty
🟢 Easy

---

## 📖 Story

The adventurer has learned judgment and logic.

Now the adventurer seeks efficiency.

Why repeat work manually,
when machines can repeat it endlessly?

The power of iteration is about to awaken.

---

## 🎯 Objectives
- [x] Learn for loop
- [x] Repeat actions automatically
- [x] Understand iteration
- [x] Learn loop conditions
- [x] Learn loop counters
- [x] Build simple repetitive programs
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
mkdir level-004-loops
cd level-004-loops
```
### 📂 STEP 2 — Initialize Go Module
```bash
go mod init level-004-loops
```

#### 🧠 What is loop?
Loop is: automatically repeat executing code.

without loop:
```text
print
print
print
print
print
```

with loop:
```text
repeat 5 times
```

Loop makes program:
- scalable,
- efficient,
- automated.

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

    for i := 1; i <= 5; i++ {
        fmt.Println("Loop iteration:", i)
    }
}
```

#### 🧠 Explanation
```
for
```
used for looping in Go.

Go only has:
```
for
```

No:
- while
- do while

#### 🧠 Structure Loop
```Go
for initialization; condition; increment
```

#### 🧠 Breakdown
i := 1

Start counter from 1.

i <= 5

Loop runs as long as condition is TRUE.

i++

Add 1 every iteration to i. 


### 📂 STEP 4 — Run Program
```bash
go run main.go
```
#### 🎯 Expected Output
```text
Loop iteration: 1
Loop iteration: 2
Loop iteration: 3
Loop iteration: 4
Loop iteration: 5
```

#### 🧠 What is Happening
Program:
1. creates counter,
2. check condition,
3. runs code,
4. increase counter,
5. repeat

This is called: iteration cycle.


### 📂 STEP 5 — Countdown Loop
update main.go to:
```Go
package main

import "fmt"

func main() {

    for i := 5; i >= 1; i-- {
        fmt.Println("Countdown:", i)
    }

    fmt.Println("Quest Started!")
}
```

#### 🧠 New concept
```
i--
```
Decrease value by 1.

#### 🧠 Loop can:
- increase
- decrease
- infinite
- conditional
- nested

### 📂 STEP 6 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
Countdown: 5
Countdown: 4
Countdown: 3
Countdown: 2
Countdown: 1
Quest Started!
```

### 📂 STEP 7 — Loop + Condition
update main.go to:
```Go
package main

import "fmt"

func main() {

    for i := 1; i <= 10; i++ {

        if i%2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
    }
}
```

#### 🧠 New Concept
% called modulo operator. used to find remainder of a division

example:

```text
4 % 2 = 0
```
means 4 is divisible by 2: even.

### 📂 STEP 8 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
1 is odd
2 is even
3 is odd
...
```

#### 🧠 Important Concept
Loop is on biggest foundation in software engineering.

Because almost every system do:
- repetition,
- iteration,
- scanning,
- processing.

example:
- reads thousand database rows,
- processing logs,
- train ML model,
- Kubernetes reconciliation loop,
- Spark distributed processing

all comes from: loops.


## What I Learned
- Learned for loop
- Repeat actions automatically
- Understand iteration
- Learned loop conditions
- Learned loop counters
- Built simple repetitive programs

## Rewards
| Attribute           | EXP |
| ------------------- | --- |
| 🖥️ Programming EXP | +20 |
| ⚙️ Backend EXP | +5  |
| 📚 Wisdom EXP       | +5  |


## Title Unlocked

🟢 Iteration Apprentice