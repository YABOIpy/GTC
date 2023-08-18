package checker

import (
	"goftc/internal/utils"
	"goftc/internal/utils/logger"
	"log"
	"time"

	"github.com/bogdanfinn/fhttp"
)

func (in *Instance) Check() (int, time.Time) {
	s := time.Now()
	req, err := http.NewRequest("GET",
		"https://discord.com/api/v9/users/@me/affinities/guilds",
		nil,
	)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("authorization", in.Token)

	resp, err := in.Client.Do(req)
	if err != nil {
		return in.Check()
	}
	defer resp.Body.Close()

	half := utils.HalfToken(in.Token, 0)
	if resp.StatusCode == 200 {
		logger.StrlogV("", half, s)
	} else if resp.StatusCode == 403 {
		logger.StrlogR("", half, s)
	} else {
		logger.StrlogE("", half, s)
	}

	return resp.StatusCode, s
}
