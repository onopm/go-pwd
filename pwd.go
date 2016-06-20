package main

import (
	"fmt"
	"os"
	"strconv"
)

const defMaxLen = 30

func main() {

	maxLen := defMaxLen

	if len(os.Args) > 1 {
		l, err := strconv.Atoi(os.Args[1])
		if err != nil {
			maxLen = defMaxLen
		} else {
			maxLen = l
		}
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
