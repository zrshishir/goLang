package main

import (
	"fmt"
	"regexp"
	
)

func main(){
	// txt := "This is how I tried to split a paragraph into a sentence. But, there is a problem. My paragraph includes dates like Jan.13, 2014 , words like U.S and numbers like 2.2. They all got split by the above code. Dr. Doctor Name is a good Mr.. "

	// // re := regexp.MustCompile("[\\!\\?\\.] +")
	// re := regexp.MustCompile(`(?m)(?<=[.!?])\s+(?=[A-Z])`)

    // sentences := re.Split(txt, -1)
    // all_together := []string{}

    // for i := range sentences {
    //     all_together = append(all_together, sentences[i])
	// 	fmt.Println(sentences[i])
    // }

	var re = regexp.MustCompile(`(\S.+?[.!?])(?=\s+|$)`)
    var str = `Hello world! How are you? I am fine. This is a difficult sentence because I use I.D.
Newlines should also be accepted. Numbers should not cause sentence breaks, like 1.23. Dr. Doctor Name ? `
    
    for i, match := range re.FindAllString(str, -1) {
        fmt.Println(match, "found at index", i)
    }
}