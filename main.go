package go_project

import (
	"fmt"
	"github.com/kivihub/go-fragment/util"
)

func GetGlobalVar() {
	psm := util.PsmFromEnv
	fmt.Println("psm", psm)
}
