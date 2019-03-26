# Golang-Entry-Level
Golang Entry Tutorial

Installation
---
https://golang.org/

hello.go
---
1. Every project need a package called **main**
2. import a Package called  **fmt**
3. **fmt**: implements formatted I/O with functions analogous to C's **printf and scanf.**

```go
package main

import (
    "fmt"
)

func main() {
  fmt.Println("Hello World!")
}
```
```shell=
go run hello.go
```

sum.go
---
1. variable statement : var ? type = ?
```go
package main

import (
    "fmt"
)

func main() {
  var x int = 5
  var y int = 7
  var sum int = x + y
  fmt.Println(sum)
}
```
```go
package main

import (
    "fmt"
)

func main() {
  var x int
  x = 5
  var y int
  y = 7
  var sum int
  sum = x + y
  fmt.Println(sum)
}
```

sum-assign.go
---
1. change to := assign
```go
package main

import (
    "fmt"
)

func main() {
  x := 5
  y := 7
  sum := x + y
  fmt.Println(sum)
}
```
