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
		pos := []string{"10", "20", "30", "40", "50", "60", "70"}
		articles := []website.Item{}
		for n, art := range(cps.Article) {
			articles = append(articles, struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{
				art.Title, art.Text, pos[rand.Intn(len(pos))], pos[rand.Intn(len(pos))], "100", "100", art.Pic, n, "0"}) 
		}
		return articles
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
