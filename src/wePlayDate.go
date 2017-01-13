package main

import (
	"net/http"
	//"fmt"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/ecommerse"
)

var wePlayDate *website.Site
var acs *website.AccountService
var ecs *ecommerse.ECommerseService

func main() {
	website.ResourceDir = ".."
	website.DataDir = "../data"
	setup()

	http.HandleFunc("/js/", website.ServeResource)
	http.HandleFunc("/css/", website.ServeResource)
	http.HandleFunc("/img/", website.ServeResource)
	http.ListenAndServe(":8070", nil)
}

func setup() {
	wePlayDate = website.CreateSite("WePlayDate", "localhost:8070")
	addScripts();
	addMenus();
	addServices();
	addPages();
	addTemplatePages();
}