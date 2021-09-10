package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var scanner *bufio.Scanner

func main() {
	var username string
	scanner = bufio.NewScanner(os.Stdin)

	fmt.Print("\033[H\033[2J")

	username = question("Hey there bud, how shall I call you? ")
	fmt.Printf("\nHello, %s! Collecting %s...", dispColor(username, Magenta), dispColor("information", Teal))
	time.Sleep(3 * time.Second)
	fmt.Printf("\n[%s] Information collection for %s went without issue.\n", dispColor("SUCCESS",Green), dispColor(username,Magenta))
}

func dispColor(text string, color string) string {
	return color + text + DefaultColor
}

func question(text string) string {
	fmt.Printf(text)
	scanner.Scan()
	return scanner.Text()
}

const (
	Black        = "\033[1;30m"
	Red          = "\033[1;31m"
	Green        = "\033[1;32m"
	Yellow       = "\033[1;33m"
	Purple       = "\033[1;34m"
	Magenta      = "\033[1;35m"
	Teal         = "\033[1;36m"
	White        = "\033[1;37m"
	DefaultColor = "\033[0m"
)
