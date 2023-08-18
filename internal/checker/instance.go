package checker

import (
	"fmt"
	"github.com/bogdanfinn/fhttp"
	"github.com/zenthangplus/goccm"
	"goftc/internal/utils"
	"goftc/internal/utils/logger"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	Logger logger.Logger
	Util   utils.Util
)

type Instance struct {
	Client *http.Client
	Token  string
}

type Stats struct {
	Valid   int
	Locked  int
	Invalid int
	All     int
}

func Configuration() []Instance {
	var (
		mutex     sync.Mutex
		Instances []Instance
		Tokens, _ = utils.ReadFile("tokens.txt")
		wg        = goccm.New(len(Tokens))
	)
	if Tokens == nil {
		log.Print("No Tokens found inside tokens.txt")
		time.Sleep(2 * time.Second)
		return nil
	}
	fmt.Print("\033[s")

	for i := 0; i < len(Tokens); i++ {
		wg.Wait()
		go func(i int) {
			defer wg.Done()

			mutex.Lock()
			fmt.Print("\033[u\033[K")
			fmt.Printf(utils.LoadingChace, strings.Split(Tokens[i], ".")[0])

			Instances = append(Instances, Instance{
				Client: &http.Client{},
				Token:  Tokens[i],
			})
			mutex.Unlock()
		}(i)
	}
	wg.WaitAllDone()

	return Instances
}
