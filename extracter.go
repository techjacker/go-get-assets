package main

import (
	"regexp"
)

type Extracter struct {
	reg string
}

func (e *Extracter) Gdrive(url string) string {
	idReg := regexp.MustCompile(e.reg)
	id := idReg.FindStringSubmatch(url)
	// didn't find a match
	if len(id) < 2 {
		return ""
	}
	return id[1]
}
