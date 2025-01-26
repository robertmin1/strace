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

	// Setup seccomp filtering.
	if err := setupSeccomp(); err != nil {
		panic(err)
	}

	if err := strace.New(c, false, func(task strace.Task, record *strace.TraceRecord) error {
		if record.Event == strace.SyscallExit || record.Event == strace.SyscallEnter {
			log.Printf("\033[1;34mpid %d: \033[1;33mSyscall Number %d\033[0m", record.PID, record.Syscall.Sysno)
		}
		return nil
	}); err != nil {
		panic(err)
	}
}

func setupSeccomp() error {
	// Create a new filter with a default action to allow all syscalls.
	filter, err := seccomp.NewFilter(seccomp.ActAllow)
	if err != nil {
		return fmt.Errorf("failed to create seccomp filter: %w", err)
	}

	if err := filter.AddRule(unix.SYS_CONNECT, seccomp.ActTrace); err != nil {
		return fmt.Errorf("failed to add connect syscall to seccomp filter: %w", err)
	}

	// Load the filter into the kernel.
	if err := filter.Load(); err != nil {
		return fmt.Errorf("failed to load seccomp filter: %w", err)
	}

	return nil
}
