package util

import (
	"fmt"
	"github.com/kivihub/go-fragment/util"
)

func SayHi() {
	fmt.Println("Hello Go!")
}

func Echo(msg string) {
	fmt.Print(msg)
}

func GetGlobalVar() {
	psm := util.PsmFromEnv
	fmt.Println("psm", psm)
}

func GetVersion() {
	fmt.Print("Version", util.Version)
}
