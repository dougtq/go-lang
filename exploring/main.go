package main

import (
  "fmt"
  "errors"
  "math"
)

type person struct {
  name string
  age int
}

func main () {
    x := 5  // shorthand of var x = 5
    y := 13

    //arr := [5]int { 1, 2, 3, 4, 5} // shorthand to create a array
    slc := []int { 1, 2, 3, 4, 5} // shorthand to create a slice
    slc = append(slc, 3)

    vert := make(map [string]int)

    vert["triangle"] = 3
    vert["square"] = 4
    vert["dodecagon"] = 12

    delete(vert, "dodecagon")

    for i := 0; i <= 10; i += 2 {
      fmt.Println(i)
    }

    for index, value := range(slc) {
      fmt.Print("Index ", index, " ")
      fmt.Print("Value ", value, "\n")
    }

    for key, value := range(vert) {
      fmt.Print("Key ", key, " ")
      fmt.Print("Value ", value, "\n")
    }

    sum := sum(x, y)

    if sum > 10 {
      fmt.Println("Sum result is", sum)
    } else if x < 5 {
      fmt.Println("Value below 5")
    } else {
      fmt.Println("Value between 10 and 5")
    }

    result, err := sqrt(-3)

    if err != nil {
      fmt.Println(err)
    } else {
      fmt.Println(result)
    }

    p := person { name: "Douglas", age: 25 }

    fmt.Println(p)
    fmt.Println(&p.name)

    val := 10
    inc(&val)
    fmt.Println(val)
    inc(&val)
    fmt.Println(val)
}

func sum (a int , b int ) int {
  return a + b
}

func sqrt (x float64) (float64, error) {
  if x < 0 {
    return 0, errors.New("Undefined for negative number")
  }

  return math.Sqrt(x), nil
}

func inc (val *int)  {
  *val++
}
