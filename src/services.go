package main

import (	
	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/service"
	"github.com/jarrancarr/website/ecommerse"
)

func addServices() {
	logger.Trace.Println("addServices()")
	acs = website.CreateAccountService()
	weePlayDate.AddService("account", acs)
	weePlayDate.AddSiteProcessor("secure", acs.CheckSecure)
	ecs = ecommerse.CreateService(acs)
	weePlayDate.AddService("ecommerse", ecs)
	mss = service.CreateService(acs)
	weePlayDate.AddService("message", mss)
}