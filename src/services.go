package main

import (
	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/service"
	//"github.com/jarrancarr/website/ecommerse"
)

func addServices() {
	logger.Trace.Println("addServices()")
	acs = website.CreateAccountService()
	weePlayDate.AddService("account", acs)
	weePlayDate.AddSiteProcessor("secure", acs.CheckSecure)
	//ecs = ecommerse.CreateService(acs)
	//weePlayDate.AddService("ecommerse", ecs)
	mss = service.CreateMessageService(acs)
	weePlayDate.AddService("message", mss)
	uts = service.CreateUtilityService(acs)
	weePlayDate.AddService("utility", uts)
	cps = CreateChildsPlayService()
	weePlayDate.AddService("childsPlay", cps)
	mss.AddHook(wpdMessageHook)
}
