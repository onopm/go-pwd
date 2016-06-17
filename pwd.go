package main

import (
	"fmt"
	"os"
	"strconv"
)

const defMaxLen = 30

func main() {

	arg := os.Args[1]
	maxLen, err := strconv.Atoi(arg)
	if err != nil {
		maxLen = defMaxLen
	}

	home := os.Getenv("HOME")
	pwd := os.Getenv("PWD")

	fmt.Printf("%s", Pwd(pwd, home, maxLen))
}

func Pwd(pwd string, home string, maxLen int) string {

	prefixHome := "~..."
	prefixRoot := "/..."

	if len(pwd) >= len(home) {

		if pwd[:len(home)] == home {
			path := "~" + pwd[len(home):]

			if len(path) > maxLen {
				l := maxLen - len(prefixHome)
				path = prefixHome + path[len(path)-l:]
				return path
			}
			return path
		}
	}

	if len(pwd) > maxLen {
		l := maxLen - len(prefixRoot)
		path := prefixRoot + pwd[len(pwd)-l:]
		return path
	}

	return pwd
}
