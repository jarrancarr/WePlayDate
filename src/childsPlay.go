package main

import (
	//"fmt"
	//"net/http"
	//"errors"
	//"time"
	//"strings"
	"math/rand"
	"sync"

	"github.com/jarrancarr/website"
)

type ChildsPlayService struct {
	Article map[string]*Post
	Lock    sync.Mutex
}

func (cps *ChildsPlayService) Status() string {
	return "good"
}

func (cps *ChildsPlayService) Execute(data []string, s *website.Session, p *website.Page) string {
	logger.Trace.Println("ChildsPlayService.Execute(" + data[0] + ", page<" + p.Title + ">)")
	switch data[0] {
	case "localPosts":
		return "hi"
	}
	return ""
}

func (cps *ChildsPlayService) Get(p *website.Page, s *website.Session, data []string) website.Item {
	logger.Debug.Println("ChildsPlayService.Get(page<" + p.Title + ">, session<" + s.GetUserName() + ">, " + data[0] + ")")
	switch data[0] {
	case "localPosts":
		return []website.Item{
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Big Dog", "Check out this huge dog", "10", "10", "90", "120", "bigDog.jpg", "bd", "0"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Cute Kitten", "Super Awe....", "10", "30", "130", "90", "cuteCat.jpg", "ck", "1"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Ice Cream", "Making ice cream, easy", "10", "50", "130", "90", "iceCream.jpg", "ic", "1"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Black Horse", "This is a beautiful horse", "10", "70", "110", "110", "blackHorse.jpg", "bh", "1"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Funny Kid", "Kid trips over pillow", "30", "70", "140", "80", "funnyKid.jpg", "fk", "2"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Drone Spying", "Using drones to spy on neighbors", "50", "70", "100", "80", "Drone.jpg", "ds", "2"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Night Bridge", "Beautiful pics of bridges at night", "70", "70", "100", "80", "Bridge.jpg", "nb", "2"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Ice Sculpting", "There will be a class on ice sculpting in Monroe park", "70", "50", "100", "80", "IceSculpt.jpg", "is", "2"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Butterfly", "The butterfly park in Bowie has opened", "70", "30", "100", "80", "Butterfly.jpg", "bf", "2"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Dance", "Salsa Dance club", "70", "10", "100", "80", "Dance.png", "sd", "2"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Playhouse", "Playhouse givaway", "50", "10", "100", "80", "Playhouse.jpg", "ph", "2"},
			struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{"Petting Zoo", "Friendship gardens are hosting a petting zoo.", "30", "10", "100", "80", "PettingZoo.jpg", "pz", "2"}}
	}
	return nil
}

func CreateChildsPlayService() *ChildsPlayService {
	logger.Debug.Println("CreateChildsPlayService()")
	cps := ChildsPlayService{Lock: sync.Mutex{}, Article: make(map[string]*Post)}
	keys := make([]string, len(Families))
	i := 0
	for k, _ := range Families {
		keys[i] = k
		i++
	}
	for _, aName := range []string{"bd", "ck", "ic", "bh", "fk", "ds", "nb", "is", "bf", "sd", "ph", "pz", "cn", "cl"} {
		cps.Article[aName] = makeArticle()
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
