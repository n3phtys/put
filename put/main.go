package main

import "fmt"
import (
	"flag"
	"os"
	"bufio"
	"log"
	"strings"
)

func main() {

	//Command line arguments
	filePtr := flag.String("file", "./file.txt", "the input text file that is to be edited")
	insertPtr := flag.String("insert", "abc", "the input line that is to be inserted if missing")
	linePtr := flag.Int("n", -1, "if inserting, this is the line index where the line will be inserted")
	/*strictPtr := flag.Bool("s", false, "strict mode using exact whitespace comparison")
	yesPtr := flag.Bool("y", false, "automatically overwrite, do not ask for confirmation")
	noPtr := flag.Bool("c", false, "only check, never overwite data")*/

	var ask bool = false;
	var defaultAnswer bool = false;
	flag.Parse()

	input, err := parseAndSplitFile(filePtr)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if *linePtr < 0 {
		*linePtr = len(input)
	}

	var contains bool = false
	//check if contains
	contains = findInLines(input, *insertPtr)

	if contains {
		fmt.Println("File contains search term, returning")
		return
	}

	//ask for overwrite and potentially overwrite
	if confirmChanges(filePtr, *insertPtr, ask, defaultAnswer) {
		replaceFilecontent(filePtr, insertInLines(input, *insertPtr, *linePtr))
	}

	fmt.Printf("Put if Absent ran through.\n")
}

func getLineEndingFromFile(filepath *string) string {
	return "\n"
}

func parseAndSplitFile(filepath *string) ([]string, error) {
	return readLines(*filepath)
}

func findInLines(lines []string, term string) bool {
	for i := 0; i < len(lines); i++ {
		if (strings.TrimSpace(lines[i]) == strings.TrimSpace(term)) {
			return true
		}
	}
	return false
}

func insertInLines(lines []string, line string, index int) []string {
	var arr []string
	var k int = 0
	l := len(lines)
	m := l
	if index <= m {
		m += 1
	}
	for i := 0; i < m; i++ {
		if i == index {
			arr = append(arr, line)
			k -= 1
		} else {
			arr = append(arr, lines[k])
		}
		k += 1
	}
	return arr
}

func replaceFilecontent(filepath *string, lines []string) error {
	return writeLines(lines, *filepath)
}

func confirmChanges(filepath *string, line string, ask bool, def bool) bool {
	fmt.Println("The file %s does not contain the line %s", *filepath, line)
	if ask {
		return def
	} else {
		fmt.Println("Do you want to overwite this file with the new insert?")
		return askForConfirmation()
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// askForConfirmation uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user. Typically, you should use fmt to print out a question
// before calling askForConfirmation. E.g. fmt.Println("WARNING: Are you sure? (yes/no)")
func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation()
	}
}

// You might want to put the following two functions in a separate utility package.

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
