package checker

import (
	"goftc/internal/utils"
	"goftc/internal/utils/logger"
	"io"
	"log"
	"time"

	"github.com/bogdanfinn/fhttp"
)

func (in *Instance) Check() (int, time.Time, []byte) {
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
	body, _ := io.ReadAll(resp.Body)

	half := utils.HalfToken(in.Token, 0)
	switch resp.StatusCode {
	case 200:
		logger.StrlogV("", half, s)
	case 403:
		logger.StrlogR("", half, s)
	default:
		logger.StrlogE("", half, s)
	}

	return resp.StatusCode, s, body
}
