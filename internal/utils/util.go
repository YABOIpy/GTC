package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile(path string) ([]string, error) {
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

	if strings.Contains(path, "token") {
		tokens := make([]string, 0, len(lines))
		for _, line := range lines {
			if strings.Contains(line, ":") {
				tokens = append(tokens, strings.Split(line, ":")[2])
			}
		}
		if len(tokens) > 0 {
			lines = tokens
		}
	}

	return lines, scanner.Err()
}

func WriteFileArray(files string, item []string) {
	f, err := os.OpenFile(files, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	for i := 0; i < len(item); i++ {
		f.WriteString(item[i] + "\n")
	}
}

func WriteFile(files string, item string) {
	f, err := os.OpenFile(files, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	_, ers := f.WriteString(item + "\n")
	if ers != nil {
		log.Println(ers)
		return
	}
}

func Input(text string) string {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	reader.Scan()
	return reader.Text()
}

func HalfToken(T string, v int) string {
	return strings.Split(T, ".")[v]
}
