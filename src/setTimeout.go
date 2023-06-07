package main

import "time"

func setTimeout(someFunc func(), milliseconds int) {
	timeout := time.Duration(milliseconds) * time.Millisecond
	time.AfterFunc(timeout, someFunc)
}
