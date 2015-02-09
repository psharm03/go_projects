package main

import (
  "fmt"
  "time"
)

type Vertex struct {
  X, Y int
}

var (
  v1 = Vertex{1, 2}  // has type Vertex
  v2 = Vertex{X: 1}  // Y:0 is implicit
  v3 = Vertex{}      // X:0 and Y:0
  p  = &Vertex{1, 2} // has type *Vertex
)

func say(s string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func main() {
  var a [2]string
  a[0] = "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  p := []int{2, 3, 5, 7, 11, 13}
  fmt.Println("p ==", p)

  for i := 0; i < len(p); i++ {
    fmt.Printf("p[%d] == %d\n", i, p[i])
  }

  var abc []int
  printSlice("a", abc)
  
  abc = append(abc, 0)
  printSlice("a", abc)

  a := make([]int, 5)
  printSlice("a", a)
  b := make([]int, 0, 5)
  printSlice("b", b)
  c := b[:2]
  printSlice("c", c)
  d := c[2:5]
  printSlice("d", d)

  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)

  fmt.Println(v1, p, v2, v3)

  go say("world")
  say("hello")

}
