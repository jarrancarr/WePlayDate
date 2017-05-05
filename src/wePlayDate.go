package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"encoding/json"
	//"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
	//"sync"

	"github.com/jarrancarr/website"
	//"github.com/jarrancarr/website/ecommerse"
	"github.com/jarrancarr/website/service"
)

var (
	weePlayDate *website.Site
	acs         *website.AccountService
	//ecs         	*ecommerse.ECommerseService
	mss            *service.MessageService
	cps            *ChildsPlayService
	logger         *website.Log
	Date_Format    = "MM/dd/yyyy"
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
	go collectMetrics(cps, mss, acs, weePlayDate)

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

	//website.Users = []website.Account{*(Families["jjlcarr"].Login), *(Families["adaknight"].Login)}
}

func MainInitProcessor(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Trace.Println("MainInitProcessor(w http.ResponseWriter, r *http.Request, website.Session<<" + s.GetId() + ">>, p *website.Page)")
	http.SetCookie(w, &http.Cookie{"testCookie", "accepted", "/", p.Site.Url,
		time.Now().Add(time.Minute * 2), "", 50000, false, true, "none=none", []string{"none=none"}})
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
	familyName := r.Form.Get("familyName")
	email := r.Form.Get("email")
	zip := r.Form.Get("zip")

	logger.Info.Println("userName: " + userName + "   Email:" + email + "  ")
	child := "data"
	children := make([]*Person, 0)
	for i := 0; child != ""; i++ {
		child = ""
		childSpecs := strings.Split(r.Form.Get(fmt.Sprintf("child%d", i)), "|")
		if len(childSpecs) == 3 {
			dob, _ := time.Parse(childSpecs[1], Date_Format)
			children = append(children, &Person{Name: []string{childSpecs[0], familyName}, DOB: dob, Male: (childSpecs[2] == "Boy"), Admin: false})
			child = childSpecs[0]
			logger.Info.Println("child: " + child + ", " + children[i].DOB.Format(Date_Format_GL))
			s.Item["numChildren"] = i + 1
		}
	}
	parent := "data"
	parents := make([]*Person, 0)
	for i := 0; parent != ""; i++ {
		parent = ""
		parentSpecs := strings.Split(r.Form.Get(fmt.Sprintf("parent%d", i)), "|")
		if len(parentSpecs) == 2 {
			parents = append(parents, &Person{Name: []string{parentSpecs[0], familyName}, Male: (parentSpecs[1] == "Dad"), Admin: false})
			parent = parentSpecs[0]
			logger.Info.Println("parent: " + parent)
			s.Item["numParents"] = i + 1
		}
	}
	s.Data["retry"] = "#applyModal"
	s.Data["error"] = "A user already exists with that user name."
	s.Data["zip"] = zip
	s.Data["email"] = email
	type data struct{ Name, MOB, Sex string }
	pData := []data{}
	for _, p := range parents {
		if p.Male {
			pData = append(pData, data{p.Name[0], "", "Dad"})
		} else {
			pData = append(pData, data{p.Name[0], "", "Mom"})
		}
	}
	s.Item["parentData"] = pData
	cData := []data{}
	for _, c := range children {
		if c.Male {
			cData = append(cData, data{c.Name[0], c.DOB.Format(Date_Format_GL), "Boy"})
		} else {
			cData = append(cData, data{c.Name[0], c.DOB.Format(Date_Format_GL), "Girl"})
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

	Families[userName] = &Family{Login: &website.Account{[]string{""}, userName, base64.URLEncoding.EncodeToString(secret),
		email, []*website.Role{website.StandardRoles["basic"]}, false, time.Now().Add(time.Minute * 15)},
		Parent: parents, Child: children, Zip:[]string{zip}}

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
		if len(fam.Parent) > 0 {
			s.Data["name"] = fam.Parent[0].Name[0]
			if len(fam.Parent) > 1 {
				s.Data["name"] += ", " + fam.Parent[1].Name[0]
			}
			for _, ch := range fam.Child {
				s.Data["name"] += ", " + ch.Name[0]
			}
		} else {
			s.Data["name"] = fam.Child[0].Name[0]
			for _, ch := range fam.Child[1:] {
				s.Data["name"] += ", " + ch.Name[0]
			}
		}
		s.Data["name"] += " " + fam.Parent[0].Name[1]
		s.Data["userName"] = userName
		s.Item["family"] = fam
		acs.Active[userName] = s
		for _, z := range fam.Zip {
			mss.Execute([]string{"addRoom", z, ""}, s, p)
		}
		if !s.Cookie {
			return r.Form.Get("redirect") + "?_id=" + s.Data["id"], nil
		}
		return r.Form.Get("redirect"), nil
	}
	s.Data["retry"] = "#loginModal"
	s.Data["error"] = "We do not recognized that user name and password"
	return "#errorModal", errors.New("failed login")
}
func SelectFamilyMember(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Trace.Println("WeePlayDate.SelectFamilyMember(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	name := ""
	first := true
	if r.Form["parent"] != nil {
		for _, nm := range r.Form["parent"] {
			if first {
				first = false
			} else {
				name += "/"
			}
			name += strings.Split(nm, " ")[0]
		}
	}
	if r.Form["child"] != nil {
		for _, nm := range r.Form["child"] {
			if first {
				first = false
			} else {
				name += "/"
			}
			name += strings.Split(nm, " ")[0]
		}
	}
	fm, good := s.Item["family"].(*Family)
	if good {
		name += " " + fm.Parent[0].Name[1]
	} else {
		logger.Warning.Println("Could not find family in session! ")
	}
	s.Data["name"] = name

	return r.Form.Get("redirect"), nil
}
func WhoseThereAjaxHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Trace.Println("WeePlayDate.WhoseThereAjaxHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	httpData, _ := ioutil.ReadAll(r.Body)
	if httpData == nil || len(httpData) == 0 {
		return "", errors.New("No Data")
	}
	dataList := strings.Split(string(httpData), "&")
	roomName := strings.Split(dataList[0], "=")[1]
	room, err := mss.GetRoom(roomName)
	if err != nil {
		return "error", err
	}
	w.Write([]byte(room.WhoseThere()))
	return "ok", nil
}
func GetFamilyProfileAjaxHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Trace.Println("WeePlayDate.GetFamilyProfileAjaxHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	data := pullData(r)
	fam := Families[data["user"]]
	if fam == nil {
		return "", errors.New("No faminly found")
	}
	var thisPerson *Person
	for _, fm := range fam.Parent {
		if fm.Name[0] == strings.Split(data["name"], "+")[0] {
			thisPerson = fm
		}
	}
	if thisPerson == nil {
		for _, fm := range fam.Child {
			if fm.Name[0] == strings.Split(data["name"], "+")[0] {
				thisPerson = fm
			}
		}
	}
	if thisPerson == nil {
		return "", errors.New("No Family by that name")
	}
	w.Write([]byte(`{"name":"` + thisPerson.FullName() + `", "age":"` + thisPerson.Age() + `", "sex":"` + thisPerson.Sex() + `", "profile":"` + thisPerson.Profile +
		`", "likes":"` + strings.Join(thisPerson.Likes, "|") + `", "user":"` + data["user"] + `"}`))
	return "ok", nil
}
func GetPersonProfileAjaxHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Debug.Println("WeePlayDate.GetPersonProfileAjaxHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	data := pullData(r)
	logger.Debug.Println(data["user"] + " " + data["name"])
	family := Families[data["user"]]
	name := strings.Split(data["name"], " ")[0]
	if family == nil {
		return "", errors.New("No such user family in system")
	}
	var thisPerson *Person
	for _, fm := range family.Parent {
		logger.Debug.Println("comparing "+fm.Name[0]+"=="+ name)
		if fm.Name[0] == name {
			thisPerson = fm
		}
	}
	if thisPerson == nil {
		for _, fm := range family.Child {
			logger.Debug.Println("comparing "+fm.Name[0]+"=="+ name)
			if fm.Name[0] == name {
				thisPerson = fm
			}
		}
	}
	if thisPerson == nil {
		return "", errors.New("No such person")
	}
	w.Write([]byte(`{"name":"` + thisPerson.FullName() + `", "age":"` + thisPerson.Age() + `", "pic":"` + thisPerson.ProfilePic +
		`", "sex":"` + thisPerson.Sex() + `", "profile":"` + thisPerson.Profile +
		`", "likes":"` + strings.Join(thisPerson.Likes, "|") + `"}`))
	return "ok", nil
}
func GetArticleAjaxHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Trace.Println("WeePlayDate.GetArticleAjaxHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	info := pullData(r)
	if cps.Places[info["place"]] == nil {
		return "", errors.New("No such place as " + info["place"])
	}
	if cps.Places[info["place"]].Article[info["articleName"]] == nil {
		return "", errors.New("No article by name " + info["articleName"] + " in this place")
	}
	article := cps.Places[info["place"]].Article[info["articleName"]]
	w.Write([]byte(`{"title":"` + article.Title + `", "author":"` + article.Author.FullName() + `", "text":"` + article.Text + `", "pic":"` + article.Pic + `", "user":"` + article.User + `"}`))
	return "ok", nil
}
func EditDataPostHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Debug.Println("EditDataPostHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	userName := r.Form.Get("user")
	text := r.Form.Get("edit")
	field := r.Form.Get("field")
	logger.Debug.Printf("inputs: %s,%s,%s", userName, field, text)
	fam := Families[userName]
	if fam == nil {
		return "", errors.New("No family by that user id")
	}
	switch field {
	case "familyProfile":
		fam.Profile = text
		break
	}
	return r.Form.Get("redirect"), nil
}
func UpdateFieldAjaxHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Trace.Println("UpdateFieldAjaxHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	data := pullData(r)
	for k, v := range(data) {
		logger.Debug.Println(k+" = "+v)
	}
	fam := Families[s.GetUserName()]
	if fam == nil {
		return "", errors.New("No family by that user id")
	}
	field := strings.Split(data["field"],":")
	if len(field)<2 {
		return "", errors.New("Not enough data found.")
	}
	
	htmlProfile := ""
	
	switch (field[0]) {
		case "personProfile" : 
			person := fam.GetFamilyMember(field[1])
			if person != nil {
				if len(field) == 3 {
					if field[2] == "undo" {
						s.AddData("redo", person.Profile)
						person.Profile = s.GetData("undo")
						htmlProfile += `<a title="Redo" class="modalButton undo" onclick="update('personProfile:`+field[1]+`:redo', '', $(this).parent())">R</a>`
					}
					if field[2] == "redo" {
						person.Profile = s.GetData("redo")
						htmlProfile += `<a title="Undo" class="modalButton undo" onclick="update('personProfile:`+field[1]+`:undo', '', $(this).parent())">U</a>`
					}
				} else {
					s.AddData("undo",person.Profile)
					person.Profile = data["data"]
					htmlProfile += `<a title="Undo" class="modalButton undo" onclick="update('personProfile:`+field[1]+`:undo', '', $(this).parent())">U</a>`
				}
				for _, line := range(strings.Split(person.Profile,"\n")) {
					htmlProfile += "<p>"+line+"</p>"
				}
			}
			htmlProfile += `<a title="Edit" class="modalButton edit peekaboo" onclick="input('New profile for `+
				field[1]+`', this, 'personProfile:`+field[0]+`:`+field[1]+`', 4, 180, 0,0,0,0);">E</a>`
			htmlProfile += `<a title="Comment" class="modalButton comment" style="display:none;">C</a>`
			break
		case "familyProfile" :
			if len(field) == 3 {
				if field[2] == "undo" {
					s.AddData("redo", fam.Profile)
					fam.Profile = s.GetData("undo")
					htmlProfile += `<a title="Redo" class="modalButton undo" onclick="update('familyProfile:`+field[1]+`:redo', '', $(this).parent())">R</a>`
				}
				if field[2] == "redo" {
					fam.Profile = s.GetData("redo")
					htmlProfile += `<a title="Undo" class="modalButton undo" onclick="update('familyProfile:`+field[1]+`:undo', '', $(this).parent())">U</a>`
				}
			} else {
				s.AddData("undo", fam.Profile)
				fam.Profile = data["data"]
				htmlProfile += `<a title="Undo" class="modalButton undo" onclick="update('familyProfile:`+field[1]+`:undo', '', $(this).parent())">U</a>`
			}
			for _, line := range(strings.Split(fam.Profile,"\n")) {
				htmlProfile += "<p>"+line+"</p>"
			}
			htmlProfile += `<a title="Edit" class="modalButton edit peekaboo" onclick="input('New profile for `+
				field[1]+`', this, 'familyProfile:`+field[0]+`:`+field[1]+`', 4, 180, 0,0,0,0);">E</a>`
			htmlProfile += `<a title="Comment" class="modalButton comment" style="display:none;">C</a>`
			break
		
		case "personProfileComment":
			otherFam := Families[field[1]]
			person := fam.Parent[0] // need to get which family member
			otherPerson := otherFam.GetFamilyMember(field[2])
			if otherPerson != nil {
				logger.Debug.Println("personProfileComment: "+otherPerson.FullName())
				otherPerson.CommentOn(person, "personProfileComment", data["data"])
			}
			htmlProfile += `<a title="Comment" class="modalButton comment peekaboo" onclick="input('comment on profile',this, 'personProfileComment:`+field[1]+`:`+field[2]+`', 12, 65, 0,0,0,0);">C</a>`
			htmlProfile += `<a title="Edit" class="modalButton edit" style="display:none;">E</a>`
			for _, line := range(strings.Split(otherPerson.Profile,"\n")) {
				htmlProfile += "<p>"+line+"</p>"
			}
			break
		
		case "personProfilePicComment":
			otherFam := Families[field[1]]
			person := fam.Parent[0] // need to get which family member
			otherPerson := otherFam.GetFamilyMember(field[2])
			if otherPerson != nil {
				logger.Debug.Println("personProfilePicComment: "+otherPerson.FullName())
				otherPerson.CommentOn(person, "personProfilePicComment", data["data"])
			}
			htmlProfile += `<img src="../img/profile/`+otherPerson.ProfilePic+`">`
			htmlProfile += `<a title="Comment" class="modalButton comment peekaboo" onclick="input('comment on profile Pic',this, 'personProfilePicComment:`+field[1]+`:`+field[2]+`', 12, 65, 0,0,0,0);">C</a>`
			htmlProfile += `<a title="Edit" class="modalButton edit" style="display:none;">E</a>`
			break
	}
	
	w.Write([]byte(htmlProfile))
	return "ok", nil
}
func GetMapAjaxHandler(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	logger.Debug.Println("GetMapAjaxHandler(w http.ResponseWriter, r *http.Request, session<" + s.GetId() + ">, page<" + p.Title + ">)")
	data := pullData(r)
	pl := cps.Places[data["place"]]
	htmlModal := `<div style="background: url('../img/map/`+pl.MapPhoto+`'); width: 800px; height: 800px; background-size: 100% 100%;"><a href="#closeModal" title="Close" class="closeModal modalButton">X</a><h2>Map</h2>`
	htmlModal += `<p>opening map:`+data["place"]+`</p>`
	if pl == nil {
		htmlModal += `<p>sorry, that place is not found.</p>`
	} else {
		htmlModal += `<p>`+fmt.Sprintf("%d", len(pl.Activities))+` activities</p>`
		htmlModal += `<p>`+fmt.Sprintf("%d", len(pl.Article))+` articles</p>`
		htmlModal += `<p>lat:`+fmt.Sprintf("%f", pl.Center.Lat)+` long:`+fmt.Sprintf("%f", pl.Center.Long)+`</p>`
	}
	htmlModal += `</div>`
	w.Write([]byte(htmlModal))
	return "ok", nil
}
func pullData(r *http.Request) map[string]string {
	logger.Trace.Println("pullData(r *http.Request)")
	httpData, _ := ioutil.ReadAll(r.Body)
	if httpData == nil || len(httpData) == 0 { return nil }
	reader := bytes.NewReader(httpData)
	mapData := new(map[string]string)
    if err := json.NewDecoder(reader).Decode(mapData); err != nil {
		logger.Error.Println("error decoding json string: "+string(httpData))
        return nil
    }
	return *mapData
}