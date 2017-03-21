package main

import (
	"cert-check/check"
	"cert-check/config"
)

func init() {
	config.Setup()
}

func main() {
	check.Check()

}
