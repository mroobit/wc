package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// pull in specified file to read
	// pull in flags, adjust behaviors based on flags
	// ccwc -c should output number of bytes in a file
	// so too should --bytes

	filename := os.Args[len(os.Args)-1] // file to read

	bytesFlagA := flag.Bool("c", false, "print bytes count\t(boolean)")
	//bytesFlagB := flag.Bool("bytes", false, "print bytes count\t(boolean)")
	linesFlag := flag.Bool("l", false, "print line count\t(boolean)")
	wordsFlag := flag.Bool("w", false, "print word count\t(boolean)")
	charsFlag := flag.Bool("m", false, "print character count\t(boolean)")

	flag.Parse()

	words, lines, chars := 0, 0, 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	defer file.Close()

	read := bufio.NewScanner(file)
	for read.Scan() {
		line := read.Text()
		w := strings.Fields(line)
		words += len(w)
		chars += len(string(line))
		lines++
	}

	wholeFile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading in whole file: ", err)
		return
	}
	stringified := fmt.Sprintf("%s", wholeFile)
	chars = len(stringified)

	bytes := len(wholeFile)

	if *linesFlag {
		fmt.Printf(" %d", lines)
	}
	if *wordsFlag {
		fmt.Printf(" %d", words)
	}
	if *bytesFlagA { //|| *bytesFlagB {
		fmt.Printf(" %d", bytes)
	}
	if *charsFlag {
		fmt.Printf(" %d", chars)
	}
	fmt.Printf(" %s\n", filename)
}
