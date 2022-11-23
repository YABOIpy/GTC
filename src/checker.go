package source

import (
	"net/http"
	"fmt"
)

func (Xc *Checker) Check(token string) string {

	req, err := http.NewRequest("GET", url, nil)
	Xc.Errs(err)
	req.Header.Set("authorization", token)
	resp, err := Xc.Client.Do(req)
	Xc.Errs(err)
	var typ = Xc.Resp

	if resp.StatusCode == 200 {
		typ = true
		fmt.Println("["+g+"✓"+r+"] ("+u+""+g+"VALID"+u+""+rb+"):", token)
		Xc.Valid++
		
	} else if resp.StatusCode == 403 {		
		typ = false
		fmt.Println("[\033[33m/"+r+"] ("+u+""+c+"LOCKED"+u+""+rb+"):", token)
		Xc.Locked++
	} else {
		fmt.Println("\033[31mX"+r+"] ("+u+""+c+"INVALID"+u+""+rb+"):", token)
		Xc.Invalid++
	}
	
	Xc.All++
	Xc.Resp = typ
	Xc.Token = token
	return Xc.Token
}

func X() Checker {
	x := Checker{}
	return x
}
