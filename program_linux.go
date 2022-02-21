//go:build linux
// +build linux

package go_program_tuning

import (
	"syscall"
)

func init() {
	if err := Tuning(); err != nil {
		panic(err)
	}
}

const (
	defaultFileMax = 1048576
)

func Tuning() error {
	var (
		rLimit syscall.Rlimit
		err    error
	)

	if err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		return err
	}

	if rLimit.Max < defaultFileMax {
		rLimit.Max = defaultFileMax
		rLimit.Cur = defaultFileMax
	}

	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}
