# strace
Simple demonstration of tracing processes in Go using `ptrace` and `seccomp`

# Overview
This program allows you to trace syscalls for a specified command, and it integrates seccomp filtering to intercept and trace the connect syscall specifically.

# Installation
```
git clone https://github.com/robertmin1/strace
```

```
go mod tidy
```
# Usage
To trace a command and its syscalls, run the following command:
```
go run main.go <command> [args...]
```
Where <command> is the command you want to trace, and [args...] are its arguments. For example:
```
go run main.go wget google.com
```

# Seccomp Filtering
The program sets up a seccomp filter to trace only connect syscalls while allowing all other syscalls to pass through. 

`PS:` The upstream implementation for tracking the entry and exit of syscalls while seccomp is active is not fully implemented.

Using `seccomp` approach speeds up tracing by avoiding unnecessary syscalls, but is still slow for larger applications. 
For better performance, I recommend using seccomp's user-space notification mechanism, which is up to three times faster based on personal testing. 

I have a PoC in Golang that I’ll write about soon.
