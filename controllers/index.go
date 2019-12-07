package controllers

import (
	"../core/distributor"
	"./Connect"
)

func Init() {
	dis := distributor.Distributor{}

	dis.HandleType("connect", Connect.Init)
	dis.HandleType("disconnect", Connect.Init)
}
