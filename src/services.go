package main

import (	
	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/ecommerse"
)

func addServices() {
	logger.Trace.Println("addServices()")
	acs = website.CreateAccountService()
	wePlayDate.AddService("account", acs)
	wePlayDate.AddSiteProcessor("secure", acs.CheckSecure)
	ecs = ecommerse.CreateService(acs)
	wePlayDate.AddService("ecommerse", ecs)
}