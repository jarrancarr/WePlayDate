package main

import (
	"net/http"
	//"fmt"
	"io/ioutil"
	"os"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/ecommerse"
)

var (
	wePlayDate *website.Site
	acs *website.AccountService
	ecs *ecommerse.ECommerseService
	logger *website.Log
)

func main() {
	website.ResourceDir = ".."
	website.DataDir = "../data"
	//logger = website.NewLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr, os.Stdout)
	logger = website.NewLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout, os.Stdout)

	setup()
	http.HandleFunc("/js/", website.ServeResource)
	http.HandleFunc("/css/", website.ServeResource)
	http.HandleFunc("/img/", website.ServeResource)
	http.ListenAndServe(":8070", nil)
}

func setup() {
	logger.Trace.Println("setup()")
	wePlayDate = website.CreateSite("WePlayDate", "localhost:8070")
	addScripts();
	addMenus();
	addServices();
	addPages();
	addTemplatePages();
}