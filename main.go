package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type JsonObject struct {
	Categories, Samples Source string
}

func main() {
	file, err := os.Open("tld.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		//if text doesnt start with // or *
			//obj := JsonObj{
				//Source : regexexamples.com,
				//Samples: [],
				//Categories: [Website, Email]
				//Regexp: asdasdasd
			//}
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
