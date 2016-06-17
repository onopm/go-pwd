package main

import (
	"testing"
)

func TestRoot(t *testing.T) {
	home := "/home/jack"

	if pwd := Pwd("/", home, 20); pwd != "/" {
		t.Error("/, ", pwd)
	}

	if pwd := Pwd("/usr/local/bin/go", home, 18); pwd != "/usr/local/bin/go" {
		t.Error("over, ", pwd)
	}
	if pwd := Pwd("/usr/local/bin/go", home, 17); pwd != "/usr/local/bin/go" {
		t.Error("equal, ", pwd)
	}
	if pwd := Pwd("/usr/local/bin/go", home, 16); pwd != "/...local/bin/go" {
		t.Error("under, ", pwd)
	}

}
func TestHome(t *testing.T) {

	home := "/home/jack"

	if pwd := Pwd("/home/jack", home, 20); pwd != "~" {
		t.Error("~, ", pwd)
	}

	if pwd := Pwd("/home/jack/local/bin/go", home, 30); pwd != "~/local/bin/go" {
		t.Error("over, ", pwd)
	}
	if pwd := Pwd("/home/jack/local/bin/go", home, 13); pwd != "~...al/bin/go" {
		t.Error("under, ", pwd)
	}
	if pwd := Pwd("/home/jack/local/bin/go", home, 14); pwd != "~/local/bin/go" {
		t.Error("equal, ", pwd)
	}
}
