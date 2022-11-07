package main

import (
	"fmt"
	"regexp"
	
)

func main(){
	txt := "This is how I tried to split a paragraph into a sentence. But, there is a problem. My paragraph includes dates like Jan.13, 2014 , words like U.S and numbers like 2.2. They all got split by the above code."

	re := regexp.MustCompile("[\\!\\?\\.] +")

    sentences := re.Split(txt, -1)
    all_together := []string{}

    for i := range sentences {
        all_together = append(all_together, sentences[i])
		fmt.Println(sentences[i])
    }
}