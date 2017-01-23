package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"crypto/rand"
	"encoding/base64"
	"time"

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
	http.HandleFunc("/audio/", website.ServeResource)
	http.ListenAndServe(":8070", nil)
}

func setup() {
	logger.Trace.Println("setup()")
	wePlayDate = website.CreateSite("WePlayDate", "localhost:8070")
	acs := website.CreateAccountService()
	wePlayDate.AddService("account", acs)
	wePlayDate.AddSiteProcessor("secure", acs.CheckSecure) // check for logged in user
	addScripts();
	addMenus();
	addServices();
	addPages();
	addTemplatePages();
}



func RegisterPostHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Debug.Println("AccountService.RegisterPostHandler(w http.ResponseWriter, r *http.Request, session<"+s.GetId()+">, page<"+p.Title+">)")
	userName := r.Form.Get("userName")
	email := r.Form.Get("email")
	
	logger.Info.Println("userName: "+userName+"   Email:"+email+"  ")
	child := "data"
	for i := 0; child != "" ; i++ {
		child = r.Form.Get(fmt.Sprintf("child%d",i))
		logger.Info.Println("child: "+child)
	}
	parent := "data"
	for i := 0; parent != "" ; i++ {
		parent = r.Form.Get(fmt.Sprintf("parent%d",i))
		logger.Info.Println("parent: "+parent)
	}
	
	secret := make([]byte, 16)
	rand.Read(secret)
	
	website.Users = append(website.Users, website.Account{"", "Logan", "J", "Carr", "", "lcarr", "LCarr48", 
		base64.URLEncoding.EncodeToString(secret), []*website.Role{website.StandardRoles["basic"]}, false, time.Now().Add(time.Minute*15)})
	
	// email user the key to log in.
	logger.Info.Println("Log in key is: "+base64.URLEncoding.EncodeToString(secret))
	
	return r.Form.Get("redirect"), nil
}