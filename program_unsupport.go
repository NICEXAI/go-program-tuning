//go:build !linux
// +build !linux

package go_program_tuning

import (
	"fmt"
	"github.com/fatih/color"
)

func init() {
	if err := Tuning(); err != nil {
		color.Red("GoProgramTuning 目前仅支持 Linux 平台，该 SDK 将不会被启用")
	}
}

func Tuning() error {
	return fmt.Errorf("该 SDK 目前只支持 Linux 平台运行，请在 Linux 环境下安装并执行")
}
