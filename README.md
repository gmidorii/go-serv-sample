# Go-Serv-Sample

### Memo
hook_unix.go
```go
package net

import "syscall"

var (
	testHookDialChannel  = func() {} // for golang.org/issue/5349
	testHookCanceledDial = func() {} // for golang.org/issue/16523

	// Placeholders for socket system calls.
	socketFunc        func(int, int, int) (int, error)         = syscall.Socket
	closeFunc         func(int) error                          = syscall.Close
	connectFunc       func(int, syscall.Sockaddr) error        = syscall.Connect
	listenFunc        func(int, int) error                     = syscall.Listen
	acceptFunc        func(int) (int, syscall.Sockaddr, error) = syscall.Accept
	getsockoptIntFunc func(int, int, int) (int, error)         = syscall.GetsockoptInt
)
```

#### syscall
- socket
- accept
	- get connection from socket
	- create new socket(connected)
	- return new socket
