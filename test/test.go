package main

import (
	"fmt"

	"github.com/adlindo/gocom"
)

func main() {

	fmt.Println("====>> ADL Common Lib Test <<====")

	gocom.AddCtrl(GetTestCtrl())
	gocom.AddCtrl(GetKeyValCtrl())
	gocom.AddCtrl(GetSecretCtrl())
	gocom.AddCtrl(GetJWTCtrl())
	gocom.AddCtrl(GetPubSubCtrl())
	gocom.AddCtrl(GetDistObjCtrl())
	gocom.AddCtrl(GetDistLockCtrl())

	gocom.Start()
}
