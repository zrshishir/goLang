# List

### Definition
A list is a sequence of elements. Each element can be connected to another with a link in a forward or backward direction. The element can have other payload properties. This data structure is a basic type of container. Lists have a variable length and developer can remove or add elements more easily than an array. Data items within a list need not be contiguous in memory or on disk. Linked lists were proposed by Allen Newell, Cliff Shaw, and Herbert A. Simon at RAND Corporation.

To get started, a list can be used in Go, as shown in the following example; elements are added through the PushBack method on the list, which is in the container/list

### Exammple: 
    package main

    import (
    "bufio"
    "container/list"
    "fmt"
    "log"
    "os"
    "strconv"
    )

    func main() {
        var inputRange int
        var listStack list.List
        scanner := bufio.NewScanner(os.Stdin)
    
        fmt.Println("Please input how many value do you want to input: ")
        if _, err := fmt.Scan(&inputRange); err != nil {
            log.Print("get input failed due to ", err)
            return
        }

        for i := 1; i <= inputRange; i++ {
            fmt.Println("Input value %d", i)
            scanner.Scan()
            input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
            listStack.PushBack(input)
        }
    
        fmt.Println("Output are: ")
        for element := listStack.Front(); element != nil; element = element.Next() {
            fmt.Println(element.Value.(int64))
        }
    }

