package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	readFile()
	writeToFile()
}

func readFile() {
	var fileName string = "shakespeare-quotes.txt"
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err.Error())
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	fmt.Printf("\n***Some famous Shakespeare quotes...***\n\n")
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("***That's all folks***")

			} else {
				fmt.Println("Error reading file ", err)
			}
			break

		}
		line = strings.TrimSpace(line)
		fmt.Println(line)

	}

}

func writeToFile() {
	var filename string = "shakespeare-quotes.txt"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	//add one more Shakespeare's quote to the file
	n, err := w.Write([]byte("\n\n9. “All that glitters is not gold.”"))
	w.Flush()
	if err != nil {
		fmt.Println("Error writing to file", err)
	} else {
		fmt.Printf("Wrote %d bytes to the file", n)
	}

}
