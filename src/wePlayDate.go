package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"os"
	"crypto/rand"
	"encoding/base64"
	"time"
	//"sync"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/service"
	"github.com/jarrancarr/website/ecommerse"
)

var (
	weePlayDate *website.Site
	acs *website.AccountService
	ecs *ecommerse.ECommerseService
	mss *service.MessageService
	logger *website.Log
)

func main() {
	initData()
	website.ResourceDir = ".."
	website.DataDir = "../data"
	//logger = website.NewLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr, os.Stdout)
	logger = website.NewLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout, os.Stdout)
	service.Logger = logger

	setup()
	
	go simulateCommunity(mss)
	
	http.HandleFunc("/js/", website.ServeResource)
	http.HandleFunc("/css/", website.ServeResource)
	http.HandleFunc("/img/", website.ServeResource)
	http.HandleFunc("/audio/", website.ServeResource)
	http.ListenAndServe(":8070", nil)
}

func setup() {
	logger.Trace.Println("setup()")
	weePlayDate = website.CreateSite("WePlayDate", "localhost:8070")
	acs := website.CreateAccountService()
	weePlayDate.AddService("account", acs)
	weePlayDate.AddSiteProcessor("secure", acs.CheckSecure) // check for logged in user
	addScripts();
	addMenus();
	addServices();
	addPages();
	addTemplatePages();
	
	website.Users = []website.Account{*(Families["jjlcarr"].Login), *(Families["adaknight"].Login),}
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
	
	website.Users = append(website.Users, website.Account{[]string{"Logan", "J", "Carr"}, userName, base64.URLEncoding.EncodeToString(secret), 
		email, []*website.Role{website.StandardRoles["basic"]}, false, time.Now().Add(time.Minute*15)})
	
	// email user the key to log in.
	logger.Info.Println("Log in key is: "+base64.URLEncoding.EncodeToString(secret))
	
	return r.Form.Get("redirect"), nil
}
func LoginPostHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Debug.Println("WeePlayDate.LoginPostHandler(w http.ResponseWriter, r *http.Request, session<"+s.GetId()+">, page<"+p.Title+">)")
	userName := r.Form.Get("UserName")
	password := r.Form.Get("Password")
	
	fam := Families[userName]
	
	if fam != nil && fam.Login.Password == password {
		logger.Debug.Println("Family: "+userName+" logging in")
		s.Data["name"] = fam.Parent[0].Name[0]
		if len(fam.Parent)>1 {
			s.Data["name"] += ", "+fam.Parent[1].Name[0]
		}
		for _, ch := range(fam.Child) {
			s.Data["name"] += ", "+ch.Name[0]
		}
		s.Data["name"] += " " + fam.Parent[0].Name[1]
		s.Data["userName"] = userName
		s.Item["family"] = fam
		acs.Active[userName] = s
		for _, z := range(fam.Zip) {
			mss.Execute([]string{"addRoom", z, ""}, s, p)
		}
		return r.Form.Get("redirect"), nil
	}
	return acs.FailLoginPage, nil
}
func SelectFamilyMember(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Debug.Println("WeePlayDate.SelectFamilyMember(w http.ResponseWriter, r *http.Request, session<"+s.GetId()+">, page<"+p.Title+">)")
	for k,v := range(r.Form) {
		logger.Debug.Println(k+"::"+strings.Join(v,"//"))
	}
	name := strings.Split(r.Form["parent"][0]," ")[0]
	for _, nm := range(r.Form["parent"][1:]) {
		name += ", " + strings.Split(nm, " ")[0]
	}
	for _, nm := range(r.Form["child"]) {
		name += ", " + strings.Split(nm, " ")[0]
	}
	name += strings.Split(r.Form["parent"][0]," ")[1]
	s.Data["name"] = name
	
	return r.Form.Get("redirect"), nil
}
func WhoseThereAjaxHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	w.Write([]byte(`["me", "myself", "Eye"]`))
	return "ok", nil
}