package main


import (
	"sync"
	"fmt"
	"time"
	"gochecker/src"
)


func main() {
	var c = source.X()
	c.Cls()
	c.Logo()
	fmt.Scanln()
	token, err := c.Read_tokens("tokens.txt")
	c.Errs(err)
	var wg sync.WaitGroup
	wg.Add(len(token))
	start := time.Now()
	for i := 0; i < len(token); i++ {
		go func(i int) {
			defer wg.Done()
			c.Check(token[i])
			if c.Resp == true {
				c.OpnFile("valid.txt",token[i])
			} else if c.Resp == false {
				c.OpnFile("locked.txt",token[i])
			} else {
				c.OpnFile("invalid.txt",token[i])
			}
			
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("[\033[32m✓\033[39m] (TIME\033[39m):", elapsed.String()[:4]+"Ms", "\033[39m(\033[33mLOCKED\033[39m):", c.Locked, "(\033[31mINVALID\033[39m):", c.Invalid, "(\033[32mVALID\033[39m):", c.Valid , "(\u001b[34;1mTOTAL\033[39m):", c.All)

}
