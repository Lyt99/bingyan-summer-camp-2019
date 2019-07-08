# Go Notes

### Imports

Basic syntax:  no big deal.

```go
import (
	"fmt"
    "math"
)
```

### Exported names

Begins with a **Capital** letter. Only exported names from a package are accessible from the outside.



### Functions

```go
func add(x int, y int) int {
    return x + y
}
```

Omit the type:

```go
func add(x, y int) int{
	return x + y
}
```

#### Multiple results

```go
package main

import "fmt"

func swap(x, y string) (string, string){
    return y, x
}

func main() {
    a, b := swap("world", "hello")
    fmt.Println(a, b)
}
```

#### Named return values

*Naked return: used in short func*

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```



### Variables

```go
var c, python, java bool //package level

func main() {
    var i int //func level
    fmt.Println(i, c, python, java)
}
```

#### Variables with initializers

```go
var i, j int = 1, 2

func main() {
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}
```

#### Short variable declarations

Only available **in** a func:

```go
func main() {
    var i, j int = 1, 2
    k := 3 // similar to auto in C++?
    c, python, java := true, false, "no!"
    fmt.Println(i, j, k, c, python, java)
}
```

#### Basic types

```go
package main

import (
	"fmt"
    "math/cmplx"
)

var (
	ToBe bool = false
    MaxInt uint64 = a<<64 - 1
    z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Println("Type: %T Value: %v\n", Tobe, Tobe)
    fmt.Println("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Println("Type: %T Value: %v\n", z, z)
}
```

#### Type conversions

```go
package main

import (
	"fmt"
    "math"
)

func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f)
    fmt.Println(x, y, z)
}
```

#### Type inference

`:=` syntax or `var = ` syntax type is inferred from the value on the right hand side



### Constants

`const` keyword. No `:=` syntax

```go
const Truth = true
```

#### Numeric Constants

High precision *values*

```go 
const (
	Big = 1 << 100
    Small = Big >> 99
)
```



## Loop

### For loop

statements, separated by simicolons:

- init
- condition
- post

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.

```go
for i := 0; i < 10; i++{ //no semicolon after post statement
    sum += i
}
```

The init and post statements are optional.

```go
func main() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}
```

#### For version of While

```go
for sum < 1000{
	sum += sum
}
```



## Conditions

### If

```go
func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}
```

#### With a short statement

Variables declared by the statement are only in scope until the end of if.

```go
func pow(x, n, lim float64) float64{
    if v:= math.Pow(x, y); v < lim{
        return v
    }
    return lim
}
```

#### If and else

```go
package main

import(
	"fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v:= math.Pow(x, n); v < lim{
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
        return lim
    }
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}
```

### Exercise

Sqrt func: z -= (z\*z - x) / (2\*z)

Newton's method!

```go
package main

import "fmt"

func sqrt(x float64) int {
    z := float64(1) // type conversion!
    i := 1
    for ; (z*z - x)*(z*z - x) > 0.0001; i++{
        z -= (z*z - x) / (2 * z)
    }
    return i
}

func main() {
    fmt.Println(sqrt(2))
}
```

### Switch

~~Constants~~, ~~integers~~, ~~break~~

top -> bottom

```go
package main

import (
	"fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on ")
    switch os:= runtime.GOOS; os {
        case "darwin":
        	fmt.Println("OS X.")
        case "linux":
        	fmt.Println("Linux.")
        default:
        fmt.Printf("%s.\n", os)
    }
}
```

#### Switch with no conditions

Also written as `switch true`

```go
t := time.Now()
switch {
    case t.Hour() < 12:
    	fmt.Println("Good morning")
    case t.Hour() < 17
    	fmt.Println("Good afternoon")
    default:
    	fmt.Println("Good evening")
}
```



## Defer

A defer statement defers the execution of a function until the surrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
func main() {
    defer fmt.Println("world")
    fmt.Printlb("hello")
}
```

#### Stacking defers

Deferred function calls are pushed onto a **stack**.

```go
func main() {
    fmt.Println("counting") //first
    for i := 0; i < 10; i++ {
        defer fmt.Println(i) //push!
    }
    fmt.Println("done") //second
}
```

