package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	//"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	//"sync"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/ecommerse"
	"github.com/jarrancarr/website/service"
)

var (
	weePlayDate *website.Site
	acs         *website.AccountService
	ecs         *ecommerse.ECommerseService
	mss         *service.MessageService
	logger      *website.Log
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
	addScripts()
	addMenus()
	addServices()
	addPages()
	addTemplatePages()

	website.Users = []website.Account{*(Families["jjlcarr"].Login), *(Families["adaknight"].Login)}
}

func MainInitProcessor(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	s.Data["numParents"] = "0"
	s.Data["numchildren"] = "0"
	return "ok", nil
}
func RegisterPostHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Debug.Println("AccountService.RegisterPostHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	userName := r.Form.Get("userName")
	email := r.Form.Get("email")
	zip := r.Form.Get("zip")

	if email == "" {
		s.Data["retry"] = "#applyModal"
		s.Data["error"] = "A user account must have a valid email address"
		return "#errorModal", errors.New("invalid data")
	}

	if zip == "" {
		s.Data["retry"] = "#applyModal"
		s.Data["error"] = "A user account must have a zip code"
		return "#errorModal", errors.New("invalid data")
	}

	if Families[userName] != nil {
		s.Data["retry"] = "#applyModal"
		s.Data["error"] = "A user already exists with that user name."
		s.Data["zip"] = zip
		s.Data["email"] = email
		type data struct {
			Name, MOB, Sex string
			Parent         bool
		}
		s.Data["numParents"] = "1"
		s.Data["numchildren"] = "2"
		s.Item["parentData"] = []data{
			data{"Mary", "No", "Mom", true}}
		s.Item["parentData"] = []data{
			data{"text", "5/2012", "Boy", false},
			data{"toto", "11/2013", "Girl", false}}
		return "#errorModal", errors.New("user already exists")
	}

	logger.Info.Println("userName: " + userName + "   Email:" + email + "  ")
	child := "data"
	children := make([]*Person, 0)
	for i := 0; child != ""; i++ {
		childSpecs := strings.Split(r.Form.Get(fmt.Sprintf("child%d", i)), "|")
		dob, _ := time.Parse(childSpecs[1], "2014-05")
		children = append(children, &Person{Name: []string{childSpecs[0]}, DOB: dob, Male: (childSpecs[2] == "Boy"), Admin: false})
		child = childSpecs[0]
		logger.Info.Println("child: " + child)
		s.Data["numChildren"] = fmt.Sprintf("%d", i)
	}
	parent := "data"
	parents := make([]*Person, 0)
	for i := 0; parent != ""; i++ {
		parentSpecs := strings.Split(r.Form.Get(fmt.Sprintf("parent%d", i)), "|")
		parents = append(parents, &Person{Name: []string{parentSpecs[0]}, Male: (parentSpecs[1] == "Dad"), Admin: false})
		parent = parentSpecs[0]
		logger.Info.Println("parent: " + parent)
		s.Data["numParents"] = fmt.Sprintf("%d", i)
	}

	secret := make([]byte, 16)
	rand.Read(secret)

	Families[userName] = &Family{&website.Account{[]string{""}, userName, base64.URLEncoding.EncodeToString(secret),
		email, []*website.Role{website.StandardRoles["basic"]}, false, time.Now().Add(time.Minute * 15)},
		parents, children, nil, []string{zip}, nil, nil, ""}

	// email user the key to log in.
	logger.Info.Println("Log in key is: " + base64.URLEncoding.EncodeToString(secret))

	return r.Form.Get("redirect"), nil
}
func LoginPostHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Debug.Println("WeePlayDate.LoginPostHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	userName := r.Form.Get("UserName")
	password := r.Form.Get("Password")

	fam := Families[userName]

	if fam != nil && fam.Login.Password == password {
		logger.Debug.Println("Family: " + userName + " logging in")
		s.Data["name"] = fam.Parent[0].Name[0]
		if len(fam.Parent) > 1 {
			s.Data["name"] += ", " + fam.Parent[1].Name[0]
		}
		for _, ch := range fam.Child {
			s.Data["name"] += ", " + ch.Name[0]
		}
		s.Data["name"] += " " + fam.Parent[0].Name[1]
		s.Data["userName"] = userName
		s.Item["family"] = fam
		acs.Active[userName] = s
		for _, z := range fam.Zip {
			mss.Execute([]string{"addRoom", z, ""}, s, p)
		}
		return r.Form.Get("redirect"), nil
	}
	s.Data["retry"] = "#loginModal"
	s.Data["error"] = "We do not recognized that user name and password"
	return "#errorModal", errors.New("failed login")
}
func SelectFamilyMember(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Debug.Println("WeePlayDate.SelectFamilyMember(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	for k, v := range r.Form {
		logger.Debug.Println(k + "::" + strings.Join(v, "//"))
	}
	name := strings.Split(r.Form["parent"][0], " ")[0]
	for _, nm := range r.Form["parent"][1:] {
		name += "/" + strings.Split(nm, " ")[0]
	}
	for _, nm := range r.Form["child"] {
		name += "/" + strings.Split(nm, " ")[0]
	}
	name += strings.Split(r.Form["parent"][0], " ")[1]
	s.Data["name"] = " " + name

	return r.Form.Get("redirect"), nil
}
func WhoseThereAjaxHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	w.Write([]byte(`["me", "myself", "Eye"]`))
	return "ok", nil
}
