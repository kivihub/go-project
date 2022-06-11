package go_project

import (
	"fmt"
	"github.com/kivihub/go-fragment/util"
)

func GetGlobalVar() {
	psm := util.PsmFromEnv
	fmt.Println("psm", psm)
}

func GetVersion()  {
	fmt.Print("Version", util.Version)
}