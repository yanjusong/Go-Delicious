package main

import (
	"fmt"
)

func getScoreDesc(score int) string {
	var desc string

	if score > 100 || score < 0 {
		desc = "invlaid score."
	} else if score < 60 {
		desc = "不及格"
	} else if score < 90 {
		desc = "良好"
	} else {
		desc = "优秀"
	}

	return desc
}

func printScoreDesc(score int) {
	desc := getScoreDesc(score)
	fmt.Printf("score:%d, desc:%s\n", score, desc)
}

func main() {
	printScoreDesc(-1)
	printScoreDesc(0)
	printScoreDesc(1)
	printScoreDesc(59)
	printScoreDesc(60)
	printScoreDesc(61)
	printScoreDesc(89)
	printScoreDesc(90)
	printScoreDesc(91)
	printScoreDesc(99)
	printScoreDesc(100)
	printScoreDesc(101)
}
