package main

import "fmt"

func main() {

   // https://gobyexample.com/variables
	
   	
   const s string = "constant"
   const n = 500000000
	
    const d = 3e20 / n
    fmt.Println(d)
    //A numeric constant has no type until it’s given one, such as by an explicit conversion.

    fmt.Println(int64(d))
   //var declares 1 or more variables  
   var a = "initial"
   fmt.Println(a)
   
   //You can declare multiple variables at once.
    var b, c int = 1, 2
    fmt.Println(b, c)

	//Go will infer the type of initialized variables.
    var d = true
    fmt.Println(d)
	
    // Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
    var e int
    fmt.Println(e)
   
    //The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.

    f := "apple"
    fmt.Println(f)
	
	x := 5
	y := 7
	sum := x + y

	fmt.Println(sum)
}
