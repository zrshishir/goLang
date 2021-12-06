# Tuples

A tuple is a finite sorted list of elements. It is a data structure that groups data. Tuples are typically immutable sequential collections. The element has related fields of different datatypes. The only way to modify a tuple is to change the fields. Operators such as + and * can be applied to tuples

### Example: 

    package main

    import (
    "fmt"
    "log"
    )
    
    func main() {
        var inputNumber int
        var squire int
        var cube int
    
        fmt.Println("Please input for squire and cube")
    
        if _, err := fmt.Scan(&inputNumber); err != nil {
            log.Print("Get input if failed due to ", err)
            return
        }
    
        squire, cube = powerSeries(inputNumber)

	    fmt.Println("Squire is: ", squire, " And cube is : ", cube)
    }

    func powerSeries(input int) (int, int) {
        return input * input, input * input * input
    }
