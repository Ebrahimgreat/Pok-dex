package main

import (
	"bufio"
	"crypto/des"
	"fmt"
	"strings"

	"os"
)

type cliCommand struct{
	name string
	description string
	callback func() error

}

func commandExit() error{

}
