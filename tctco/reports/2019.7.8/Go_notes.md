# Go notes

## Interfaces

```go
type Abser interface{
    Abs() float64
}
```

In this case:

```go
type Rect struct {
    Width float32
    Height float32
}

type Circle struct {
    Radii float32
}

func (r Rect) Area() float32 {
    return r.Width * r.Height
}

func (c Circle) Area() float32 {
    return math.Pi * c.Radii * c.Radii
}

func (r Rect) Circumference() float32 {
    return 2 * (r.Width + r.Height)
}

func (c Circle) Circumference() float32 {
    return 2 * math.Pi * c.Radii
}
```

If we want to declare a function `showInfo`, we will have to make 2, one for Rect and the other for Circle! But we can use **interface**.

```go
type Shaper interface {
    Area() float32
    Circumference() float32
}

func showInfo(s Shaper) {
    fmt.Println(s.Area(), s.Circumference())
}
```

An *interface type* is defined as a set of method signatures. A value of interface type can hold any value that implements those methods

### Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

```go
type I interface {
    M()
}

type T struct {
    S string
}

//This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M () {
    fmt.Printlb(t.S)
}
```

### Interface values

```go
package main

import (
	"fmt"
    "math"
)

type T struct {
    S string
}

type F float64

type I interface {
    M()
}

func (t *T) M() {
    fmt.Println(t.S)
}

func (f F) M() {
    fmt.Println(f)
}

func describe(i interface) {
    fmt.Printf("(%v %T)\n", i, i)// tuple(value, type)
}

func main() {
    var i I
    i = &T{"Hello"}
    describe(i)
    i.M()
    
    i = F(math.Pi)
    descirbe(i)
    i.M()
}
```

### Interface values with nil underlying values

If the concrete value inside the interface itself is nil, the **method** will be called with a nil receiver.

```go
func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}
```

### Nil interface values

Calling a method on a nil interface is a sun-time error because there is no type inside the interface tuple to indicate which **concrete method** to call.

### The empty interface

`interface{}` is known as the *empty interface*. It is used to store any type of value since every implements at least zero methods)

Empty interfaces are used by code that handles values of unknown type. *Eg: `fmt.Print` takes any number of arguments of type `interface`*

```go
func main() {
    var i interface{} // this is like a variable in python!
    describe(i)
    
    i = 42 // any type!
    describe(i)
    
    i = "hello"
    describe(i)
}

func describe(i interface{}) { // don't know which type is in
    fmt.Printf("(%v %T)\n", i, i)
}
```

### Type assertions

`t := i.(T)` asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.

To *test* whether an interface value holds a specific type, a type assertion can return 2 values.

```go
f, okay := i.(float64) // value of i and bool
```

*Review: in map, we have the similar expression*

```go
m := make(map[string]int)
m["Answer"] = 99
delete(m, "Answer")
v, okay = m["Answer"] // 0, false
```

### Type switches

```go
func do(i interface{}) {
    switch v := i.(type) {
        case int:
        	fmt.Printf("Twice %v is %v\n", v, v*2)
        case string:
        	fmt.Printf("%q is %v bytes long\n", v, len(v))
        default:
        fmt.Println("don't know about type %T!\n", v)
    }
}

func main() {
    do(21)
    do("Hello")
    do(true)
}
```

### Stringers

```go
type Stringer interface {
    String() string
}
```

One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package. The `fmt` package (and many others) look for this interface to print values.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
    fmt.Println(a, z)
}
```

#### Exercises

Make the `IPAddr` type implement `fmt.Stringer` to print the address as a dotted quad.

For instance, `IPAddr{1, 2, 3, 4}` should print as `"1.2.3.4"`.

```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("\"%d.%d.%d.%d\"", ip[0], ip[1], ip[2], ip[3])
}


func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

### Errors

The `error` type is a built-in interface similar to `fmt.Stringer`:

```go
type error interface {
    Error() string
}
```

The `fmt` package looks for the `error` interface when printing.

```go
package main

import (
	"fmt"
    "time"
)

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "it didn't work"
    }
}

func main() {
    if err := run(); err != nil{
        fmt.Println(err)
    }
}
```

#### Exercises Errors

`Sqrt` should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type

```
type ErrNegativeSqrt float64
```

and make it an `error` by giving it a

```
func (e ErrNegativeSqrt) Error() string
```

method such that `ErrNegativeSqrt(-2).Error()` returns `"cannot Sqrt negative number: -2"`.

**Note:** A call to `fmt.Sprint(e)` inside the `Error` method will send the program into an infinite loop. You can avoid this by converting `e` first: `fmt.Sprint(float64(e))`. Why?

Change your `Sqrt` function to return an `ErrNegativeSqrt` value when given a negative number.

```go
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number : %.2f", e)
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		v := float64(1)
		for (v*v-x)*(v*v-x) > 0.00001 {
			v -= (v*v - x) / (2 * v)
		}
		return v, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

### Reader

The `io` package specifies the `io.Reader` interface, which represents the read end of a stream of data.

The `io.Reader` interface has a `Read` method, which populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.

```go
package main

import (
	"fmt"
    "io"
    "string"
)

func main() {
    r := strings.NewReader("Hello, Reader")
    
    b := make([]byte 8) 
    for{
        n, err := r.Read(b) //reader and read!
        fmt.Println(n)
        if err == io.EOF {
            break
        }
    }
}
```

