package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	command := strings.Join(args, "")

	reg := regexp.MustCompile(`(\d+)\s*d\s*(\d+)(\s*\+\s*(\d+)\s*)?(x\s*(\d+))?`) // 2d6+3 x10
	parts := reg.FindStringSubmatch(command)

	if len(parts) == 0 {
		printHelpAndQuit()
	}

	sNum := parts[1]
	sDie := parts[2]
	sBon := parts[4]
	sRep := parts[6]

	iNum, err := strconv.Atoi(sNum)
	if err != nil {
		printHelpAndQuit()
	}

	iDie, err := strconv.Atoi(sDie)
	if err != nil {
		printHelpAndQuit()
	}

	iBon, err := strconv.Atoi(sBon)
	if err != nil {
		iBon = 0
	}

	iRep, err := strconv.Atoi(sRep)
	if err != nil {
		iRep = 1
	}

	rand.Seed(time.Now().UnixNano())
	calc(iNum, iDie, iBon, iRep)
}

func printHelpAndQuit() {
	fmt.Println(`Example usage: d 2d8+3 x4"`)
	fmt.Println(`Minimal usage: d 1d6"`)
	os.Exit(1)
}

func calc(num, die, bon, rep int) {
	for r := 0; r < rep; r++ {
		result := 0
		for n := 0; n < num; n++ {
			result = result + rand.Intn(die) + 1
		}

		// if we're rolling a 1d20, add a special reaction for nat-1 and nat-20
		if num == 1 && result == 20 {
			fmt.Printf("%3d  :D\n", result+bon)
		} else if num == 1 && result == 1 {
			fmt.Printf("%3d  :(\n", result+bon)
		} else {
			fmt.Printf("%3d\n", result+bon)
		}
	}
}
