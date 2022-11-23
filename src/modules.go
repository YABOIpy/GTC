package source

import (
	"bufio"
	"log"
	"os"
)


func (Xc *Checker) Read_tokens(files string) ([]string, error) {
	file, err := os.Open(files)
	Xc.Errs(err)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}


func (Xc *Checker) OpnFile(file string, token string) {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0660)
	Xc.Errs(err)
	defer f.Close()
	_, ers := f.WriteString(token + "\n")
	Xc.Errs(ers)
}


func (Xc *Checker) Errs(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
