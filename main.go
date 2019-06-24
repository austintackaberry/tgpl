package main

import (
	"os"
	ch1 "tgpl/ch1"
)

func main() {
	exName := os.Args[1]
	if exName == "1_1" {
		ch1.Ex1()
	}
	if exName == "1_2" {
		ch1.Ex2()
	}
	if exName == "1_4" {
		ch1.Ex4()
	}
	if exName == "1_5" {
		ch1.Ex5()
	}
	if exName == "1_6" {
		ch1.Ex6()
	}
	if exName == "1_7" {
		ch1.Ex7()
	}
	if exName == "1_8" {
		ch1.Ex8()
	}
	if exName == "1_9" {
		ch1.Ex9()
	}
	if exName == "1_10" {
		ch1.Ex10()
	}
	if exName == "1_11" {
		ch1.Ex11()
	}
	if exName == "1_12" {
		ch1.Ex12()
	}
}
