package dicontainer

import (
	"fmt"
	. "github.com/instance-id/GoUI/components"
)

type DIContainer struct {
	Cnt *Containers
}

var DiCon *DIContainer

func Init() {
	InitDi()
}

func InitDi() {
	DiCon.Cnt = Cntnrs

	fmt.Printf("Testing Print from one of the containers: %s", DiCon.Cnt.Dbd.Data.Address)

}
