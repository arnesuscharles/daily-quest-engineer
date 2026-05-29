# 🟢 Level 006 — Arrays & Slices

## Quest Type
Main Quest

## Difficulty
🟢 Easy

---

## 📖 Story

The adventurer has learned reusable abilities.

But the world contains more than one piece of data.

To handle many values at once,
the adventurer must learn collections.

One value becomes many.
And many values become systems.

---

## 🎯 Objectives
- [x] Learn arrays
- [x] Learn slices
- [x] Learn indexing
- [x] Learn len()
- [x] Learn looping through collections
- [x] Understand dynamic collections
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

#### VS Code — Auto Format Go Code
Shortcut:
```text
Shift + Alt + F
```
Used for auto format Go Code

---

## Step-by-Step Walkthrough

### 📂 STEP 1 — Create Level Folder
move to repo:
```bash
cd ~/workspace/daily-quest-engineer
```
create new folder:
```bash
mkdir level-006-arrays-and-slices
cd level-006-arrays-and-slices
```
### 📂 STEP 2 — Initialize Go Module
```bash
go mod init level-006-arrays-and-slices
```

### 📂 STEP 3 — Create main.go
```bash
touch main.go
```
open VS Code:
```bash
code .
```

#### 🧠 What is Array?
Array is: a group of data with fixed size.

Example:
- list of levels,
- list of items,
- list of players,
- list of numbers.


fill main.go
```Go
package main

import "fmt"

func main() {

    levels := [3]int{1, 2, 3}

    fmt.Println(levels)
}
```

#### 🧠 Explanation
```text
[3]int
```
means:
- array consist of 3 data,
- data type is integer.

#### 🧠 Inside array
```text
1
2
3
```
kept inside one variable:
```text
levels
```


### 📂 STEP 4 — Run Program
```bash
go run main.go
```
#### 🎯 Expected Output
```text
[1 2 3]
```

#### 🧠 What is Index?
Every data has position.
Starts from:
```text
0
```
Not 1.

#### ✨ Index Example
| Value | Index |
| ----- | ----- |
| 1     | 0     |
| 2     | 1     |
| 3     | 2     |


### 📂 STEP 5 — Access Specific Data
update main.go to:
```Go
package main

import "fmt"

func main() {

    levels := [3]int{1, 2, 3}

    fmt.Println(levels[0])
    fmt.Println(levels[1])
    fmt.Println(levels[2])
}
```

#### 🧠 Explanation
```text
levels[0]
```
Get data at index 0.


### 📂 STEP 6 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
1
2
3
```

#### 🧠 What is Slice?

Slice is: flexible version of array.

Slice size can be changed.

Slice is often used in Go compared to array.

### 📂 STEP 7 — Slice Program
update main.go to:
```Go
package main

import "fmt"

func main() {

    titles := []string{
        "Code Initiate",
        "Logic Apprentice",
        "Function Smith",
    }

    fmt.Println(titles)
}
```

#### 🧠 Explanation
```text
[]string
```

means slice consist of string.

No fixed size like [3].

### 📂 STEP 8 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
[Code Initiate Logic Apprentice Function Smith]
```

### 📂 STEP 9 — Learn append()
update main.go to:
```Go
package main

import "fmt"

func main() {

    titles := []string{
        "Code Initiate",
        "Logic Apprentice",
    }

    titles = append(titles, "Collection Handler")

    fmt.Println(titles)
}
```

### 📂 STEP 10 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
[Code Initiate Logic Apprentice Collection Handler]
```

#### 🧠 What is append() ?
used for: add data to slice.

### 📂 STEP 11 — Loop Through Slice
update main.go to:
```Go
package main

import "fmt"

func main() {

    titles := []string{
        "Code Initiate",
        "Logic Apprentice",
        "Function Smith",
    }

    for i, title := range titles {
        fmt.Println(i, "-", title)
    }
}
```

#### 🧠 New Concept
```text
range
```
used for: looping collection.

#### 🧠 i, title
- i -> index
- title -> value

### 📂 STEP 12 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
0 - Code Initiate
1 - Logic Apprentice
2 - Function Smith
```

### 📂 STEP 13 — Learn len()
update main.go to:
```Go
package main

import "fmt"

func main() {

    titles := []string{
        "Code Initiate",
        "Logic Apprentice",
        "Function Smith",
    }

    fmt.Println("Total Titles:", len(titles))
}
```

### 📂 STEP 14 — Run Again
```bash
go run main.go
```

#### 🎯 Expected Output
```text
Total Titles: 3
```

#### 🧠 What is len() ?
used for: count total data.

#### 🧠 Important Concept
Arrays & Slices are foundation for:
- API responses,
- database rows,
- JSON arrays,
- Spark datasets,
- Kubernetes resource lists,
- ML datasets.

Almost every data engineering & backend systems works with: collections of data.


## What I Learned
- Learned arrays
- Learned slices
- Learned indexing
- Learned len()
- Learned looping through collections
- Understand dynamic collections

## Rewards
| Attribute           | EXP |
| ------------------- | --- |
| 🖥️ Programming EXP | +25 |
| 🗃️ Data EXP | +15  |
| 📚 Wisdom EXP       | +5  |


## Title Unlocked

🟢 Collection Handler