// +build linux

package terminalgo

import "syscall"

const ioctlReadTermios = syscall.TCGETS
