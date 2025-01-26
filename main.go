package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/u-root/u-root/pkg/strace"
)

var errUsage = errors.New("usage: strace <command> [args...]")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println(errUsage)
		os.Exit(1)
	}

	c := exec.Command(args[0], args[1:]...)
	c.Stdin, c.Stdout, c.Stderr = os.Stdin, os.Stdout, os.Stderr

	if err := strace.New(c, false, func(task strace.Task, record *strace.TraceRecord) error {
		if record.Syscall.Sysno > 0 {
			log.Printf("pid %d: %d", record.PID, record.Syscall.Sysno)
		}
		return nil
	}); err != nil {
		panic(err)
	}
}
