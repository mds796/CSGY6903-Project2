package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func stopServer(pidFile string) {
	pid := readPid(pidFile)
	if pid == -1 {
		return
	}

	process, err := os.FindProcess(pid)
	panicOnError(err)

	err = process.Kill()
	if err != nil {
		log.Println(err)
	}

	err = process.Release()
	panicOnError(err)

	err = os.Remove(pidFile)
	if err != nil {
		log.Println(err)
	}
}

func readPid(pidFile string) int {
	bytes, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return -1
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(bytes)))
	panicOnError(err)

	return pid
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func writePidFile(pidFile string) {
	bytes := []byte(strconv.Itoa(os.Getpid()) + "\n")
	err := ioutil.WriteFile(pidFile, bytes, 400)

	if err != nil {
		panic(err)
	}
}
