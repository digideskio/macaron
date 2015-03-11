package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Unknwon/com"
)

func PathInsideGoPath(location string) error {
	var insideGoPath bool
	gopaths := com.GetGOPATHs()

	if len(gopaths) < 0 {
		// add error about setting gopath
		return errors.New("blah")
	}

	for _, wg := range gopaths {
		fmt.Println(wg)

		if strings.HasPrefix(strings.ToLower(location), strings.ToLower(wg)) {
			insideGoPath = true
			break
		}
	}

	if !insideGoPath {
		return errors.New("blah")
	}

	return nil
}
