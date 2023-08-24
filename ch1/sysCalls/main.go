package main

import (
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	c := make([]byte, 512)

	log.Println("Getpid: ", unix.Getpid())
	log.Println("Getpgrp: ", unix.Getpgrp())
	log.Println("Getppid: ", unix.Getppid())
	log.Println("Gettid: ", unix.Gettid())

	_, err := unix.Getcwd(c)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(c))
}
