package task

import (
	"fmt"
	"goftc/internal/checker"
	"goftc/internal/utils"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/zenthangplus/goccm"
)

func CheckerTask(in []checker.Instance) {
	f := []string{
		"data/valid.txt",
		"data/locked.txt",
		"data/invalid.txt",
		"data/failed.txt"
	}
	for i := 0; i < len(f); i++ {
		os.Truncate(f[i], 0)
	}

	var delta time.Duration
	var stats checker.Stats
	wg := goccm.New(len(in))
	token := make([]string, 0)

	s := time.Now()
	for i := 0; i < len(in); i++ {
		c := in[i]
		wg.Wait()
		go func(i int) {
			defer wg.Done()
			resp, t, body := c.Check()
			switch resp {
			case 200:
				utils.WriteFile("data/valid.txt", c.Token)
				token = append(token, c.Token)
				stats.Valid++
			case 403:
				utils.WriteFile("data/locked.txt", c.Token)
				stats.Locked++
			case 429:
				utils.WriteFile("data/failed.txt", c.Token)
				stats.Invalid++
			case 400:
				stats.Invalid++
				if strings.Contains(string(body), "cloudflare") {
					utils.WriteFile("data/failed.txt", c.Token)
				} else {
					utils.WriteFile("data/invalid.txt", c.Token)
				}
			default:
				utils.WriteFile("data/invalid.txt", c.Token)
				stats.Invalid++
			}
			stats.All++
			atomic.AddInt64((*int64)(&delta), int64(time.Since(t)))
		}(i)
	}
	wg.WaitAllDone()

	fmt.Printf(utils.CheckerFormat, time.Since(s).String()[:4], stats.Locked, stats.Invalid, stats.Valid, stats.All)
	fmt.Printf("Delta:\u001B[34;1m %s\n\u001B[0m", delta/time.Duration(len(in)))
	if stats.All != stats.Valid && stats.Valid != 0 {
		if utils.Input(utils.WriteValidMention) == "y" {
			os.Truncate("tokens.txt", 0)
			var d []string
			for i := 0; i < len(token); i++ {
				d = append(d, token[i])
			}
			utils.WriteFileArray("tokens.txt", d)
		}
	}
	stats = checker.Stats{}
	time.Sleep(5 * time.Second)
}
