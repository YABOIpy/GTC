package source

import (
	"os/exec"
	"bufio"
	"fmt"
	"log"
	"os"
)

func (Xc *Checker) Logo() {
	logo := `
	___`+z+`__  __`+r+`____`+z+`____________`+r+`___`+z+`_____`+r+`
	__  `+z+`/ / /`+r+`__  `+z+`____/__  __/`+r+`_  `+z+`____/
	`+r+`_  `+z+`/ / /`+r+`__  `+z+`/_   `+r+`__  `+z+`/  `+r+`_`+z+`  /     
	/ /_/ / `+r+`_  `+z+`__/   `+r+`_`+z+`  /   / /___   
	\____/  /_/      /_/    \____/   

	[`+r+`Go-Token-Checker`+z+`]		~YABOI#0001
	
	`+r+`Press enter to Continue..`
	fmt.Print(logo)
}

func (Xc *Checker) Cls() {
	cmd := exec.Command("cmd", "/c", "cls") 
	cmd.Stdout = os.Stdout
	cmd.Run()
}


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
