package main

import (
	"fmt"
	//"net/http"
	//"errors"
	//"time"
	"math/rand"
	"strconv"
	"strings"
	"sync"

	"github.com/jarrancarr/website"
)

type ChildsPlayService struct {
	Places  map[string]*Region
	Metrics map[string][]int
	Lock    sync.Mutex
}

func (cps *ChildsPlayService) Status() string {
	return "good"
}

func (cps *ChildsPlayService) Execute(data []string, s *website.Session, p *website.Page) string {
	logger.Trace.Println("ChildsPlayService.Execute(" + data[0] + ", page<" + p.Title + ">)")
	switch data[0] {
	case "#families":
		return fmt.Sprintf("%d", len(Families))
	}
	return ""
}

func (cps *ChildsPlayService) Get(p *website.Page, s *website.Session, data []string) website.Item {
	logger.Debug.Println("ChildsPlayService.Get(page<" + p.Title + ">, session<" + s.GetUserName() + ">, " + strings.Join(data, "|") + ")")
	switch data[0] {
	case "posts":
		pos := []string{"10", "20", "30", "40", "50", "60", "70", "80"}
		articles := []website.Item{}
		for n, art := range cps.Places[data[1]].Article {
			articles = append(articles, struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{
				art.Title, art.Text, pos[rand.Intn(len(pos))], pos[rand.Intn(len(pos))], "100", "100", art.Pic, n, "0"})
		}
		return articles
	case "getFamily":
		return Families[data[1]]
	case "getFamilies":
		var answ []interface{}
		page := 1
		page, _ = strconv.Atoi(p.Param["famPage"])
		if page < 1 {
			page = 1
		}

		famPerPage := 10
		famPerPage, _ = strconv.Atoi(p.Param["famPerPage"])
		if famPerPage < 10 {
			famPerPage = 10
		}

		index := 0

		// need to order these
		for name, family := range Families {
			if index < (page-1)*famPerPage || index > page*famPerPage {
				index += 1
				continue
			}
			index += 1
			var kids []string
			for _, k := range family.Child {
				kids = append(kids, fmt.Sprintf("%s-%s", k.FullName(), k.Age()))
			}
			fam := struct {
				UName, SirName, Dad, Mom, Home, Profile, Pic string
				Children                                     int
				Child                                        []string
			}{
				name, family.Login.Name[0], "X", "X", strings.Join(family.Zip, "/"), family.Profile, family.ProfilePic,
				len(family.Child), kids,
			}
			if family.Parent[0].Male {
				fam.Dad = family.Parent[0].Name[0]
			} else {
				fam.Mom = family.Parent[0].Name[0]
			}
			if len(family.Parent) == 2 {
				if family.Parent[1].Male {
					fam.Dad = family.Parent[1].Name[0]
				} else {
					fam.Mom = family.Parent[1].Name[0]
				}
			}
			answ = append(answ, fam)
		}
		return answ
	case "getPerson":
		if len(data) < 3 {
			logger.Warning.Println("length of data is less than 3")
			return struct{ Name, Profile, Pic string }{"Noone", "I am a ghost!", "Blankeroo.jpg"}
		}
		family := Families[data[1]]
		if family == nil {
			logger.Warning.Println("no family found")
			return struct{ Name, Profile, Pic string }{"Noone", "I am a ghost!", "Blankeroo.jpg"}
		}
		person := family.GetFamilyMember(data[2])
		if person == nil {
			logger.Warning.Println("no person in that family found")
			return struct{ Name, Profile, Pic string }{"Noone", "I am a ghost!", "Blankeroo.jpg"}
		}
		return struct{ Name, Profile, Pic string }{person.FullName(), person.Profile, person.ProfilePic}
	}
	return nil
}

func (cps *ChildsPlayService) Metric(what ...string) int {
	switch what[0] {
	case "#families":
		return len(Families)
	}
	return 0
}

func CreateChildsPlayService() *ChildsPlayService {
	logger.Debug.Println("CreateChildsPlayService()")
	cps := ChildsPlayService{Lock: sync.Mutex{}, Places: make(map[string]*Region), Metrics: make(map[string][]int)}
	keys := make([]string, len(Families))
	i := 0
	for k, _ := range Families {
		keys[i] = k
		i++
	}
	cps.Places["20720"] = &Region{Name: "20720", Article: make(map[string]*Post)}
	cps.Places["20726"] = &Region{Name: "20726", Article: make(map[string]*Post)}
	for _, aName := range []string{"bd", "ck", "ic", "bh", "fk", "ds", "nb"} {
		cps.Places["20720"].Article[aName] = makeArticle()
	}
	for _, aName := range []string{"is", "bf", "sd", "ph", "pz", "cn", "cl"} {
		cps.Places["20726"].Article[aName] = makeArticle()
	}
	return &cps
}

var pics = []string{"bigDog.jpg", "cuteCat.jpg", "iceCream.jpg", "blackHorse.jpg", "funnyKid.jpg", "Drone.jpg", "Bridge.jpg", "IceSculpt.jpg", "Butterfly.jpg", "Dance.png", "Playhouse.jpg", "PettingZoo.jpg", "City.jpg", "Country.jpg"}
var text = []string{"Check out this huge dog!", "Super Awe....", "Making ice cream, easy", "This is a beautiful horse", "Kid trips over pillow", "Using drones to spy on neighbors", "Beautiful pics of bridges at night",
	"There will be a class on ice sculpting in Monroe park", "The butterfly park in Bowie has opened", "Salsa Dance club", "Playhouse givaway", "Friendship gardens are hosting a petting zoo.", "City life at night", "Nothing like country living."}

func makeArticle() *Post {
	fam := Families[famKeys[rand.Intn(len(famKeys))]]
	par := rand.Intn(len(fam.Parent))
	art := rand.Intn(len(pics))
	return &Post{Author: fam.Parent[par], User: fam.Login.User, Pic: pics[art], Text: text[art]}
}
