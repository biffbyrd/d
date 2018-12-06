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

	reg := regexp.MustCompile(`((\d+)\s*d\s*(\d+))?(\s*\+\s*(\d+)\s*)?(x\s*(\d+))?`) // 2d6+3 x10
	parts := reg.FindStringSubmatch(command)

	if len(parts) == 0 {
		printHelpAndQuit()
	}

	sNum := parts[2]
	sDie := parts[3]
	sBon := parts[5]
	sRep := parts[7]

	iNum, err := strconv.Atoi(sNum)
	if err != nil {
		iNum = 1
	}

	iDie, err := strconv.Atoi(sDie)
	if err != nil {
		iDie = 20
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
	fmt.Println(`Example usage: d 2d8+3 x4`)
	fmt.Println(`               d 1d6`)
	fmt.Println(`               d +4`)
	fmt.Println(`               d`)
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
