package main

import (
	"cert-exp-check/check"
	"cert-exp-check/config"
)

func init() {
	config.Setup()
}

func main() {
	check.Check()

}
