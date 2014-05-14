// +build darwin freebsd netbsd openbsd

package terminalgo

import "syscall"

const ioctlReadTermios = syscall.TIOCGETA
