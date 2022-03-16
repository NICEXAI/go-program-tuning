//go:build !linux
// +build !linux

package go_program_tuning

import (
	"errors"
	"github.com/fatih/color"
)

func init() {
	if err := Tuning(); err != nil {
		color.Red(err.Error())
	}
}

func Tuning() error {
	return errors.New("GoProgramTuning 目前仅支持 Linux 平台，该 SDK 将不会被启用")
}
