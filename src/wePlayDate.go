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
	Date_Format = "MM/dd/yyyy"
	Date_Format_GL = "01/02/2006"
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
	logger.Trace.Println("MainInitProcessor(w http.ResponseWriter, r *http.Request, website.Session<<"+s.GetId()+">>, p *website.Page)");
	if s.Item["numParents"] == nil {
		s.Item["numParents"] = 0
	}
	if s.Item["numChildren"] == nil {
		s.Item["numChildren"] = 0
	} 
	return "ok", nil
}
func RegisterPostHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Trace.Println("RegisterPostHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	userName := r.Form.Get("userName")
	email := r.Form.Get("email")
	zip := r.Form.Get("zip")

	logger.Info.Println("userName: " + userName + "   Email:" + email + "  ")
	child := "data"
	children := make([]*Person, 0)
	for i := 0; child != ""; i++ {
		child = ""
		childSpecs := strings.Split(r.Form.Get(fmt.Sprintf("child%d", i)), "|")
		if len(childSpecs)==3 {
			dob, _ := time.Parse(childSpecs[1], Date_Format)
			children = append(children, &Person{Name: []string{childSpecs[0]}, DOB: dob, Male: (childSpecs[2] == "Boy"), Admin: false})
			child = childSpecs[0]
			logger.Info.Println("child: " + child + ", "+children[i].DOB.Format(Date_Format_GL))
			s.Item["numChildren"] = i+1
		}
	}
	parent := "data"
	parents := make([]*Person, 0)
	for i := 0; parent != ""; i++ {
		parent = ""
		parentSpecs := strings.Split(r.Form.Get(fmt.Sprintf("parent%d", i)), "|")
		if len(parentSpecs)==2 {			
			parents = append(parents, &Person{Name: []string{parentSpecs[0]}, Male: (parentSpecs[1] == "Dad"), Admin: false})
			parent = parentSpecs[0]
			logger.Info.Println("parent: " + parent)
			s.Item["numParents"] = i+1
		}
	}
	s.Data["retry"] = "#applyModal"
	s.Data["error"] = "A user already exists with that user name."
	s.Data["zip"] = zip
	s.Data["email"] = email
	type data struct { Name, MOB, Sex string }
	pData := []data{}
	for _, p := range(parents) {
		if p.Male {
			pData = append(pData, data{p.Name[0],"","Dad"})
		} else {
			pData = append(pData, data{p.Name[0],"","Mom"})
		}
	}
	s.Item["parentData"] = pData
	cData := []data{}
	for _, c := range(children) {
		if c.Male {
			cData = append(cData, data{c.Name[0],c.DOB.Format(Date_Format_GL),"Boy"})
		} else {
			cData = append(cData, data{c.Name[0],c.DOB.Format(Date_Format_GL),"Girl"})
		}
	}
	s.Item["childData"] = cData
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

	if Families[userName] != nil { // username is already in use
		return "#errorModal", errors.New("user already exists")
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
	name := ""
	if r.Form["parent"]!=nil {
		name = strings.Split(r.Form["parent"][0], " ")[0]
		for _, nm := range r.Form["parent"][1:] {
			name += "/" + strings.Split(nm, " ")[0]
		}		
	} else {
		return r.Form.Get("redirect"), nil
	}
	if r.Form["child"]!=nil {
		for _, nm := range r.Form["child"] {
			name += "/" + strings.Split(nm, " ")[0]
		}
	}
	name += " " + strings.Split(r.Form["parent"][0], " ")[1]
	s.Data["name"] = name

	return r.Form.Get("redirect"), nil
}
func WhoseThereAjaxHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	httpData, _ :=ioutil.ReadAll(r.Body)
	if (httpData == nil || len(httpData) == 0) {
		return "", errors.New("No Data")
	}
	dataList := strings.Split(string(httpData),"&")
	roomName := strings.Split(dataList[0],"=")[1]
	room, err := mss.GetRoom(roomName)
	if err != nil {
		return "error", err
	}
	w.Write([]byte(room.WhoseThere()))
	return "ok", nil
}
