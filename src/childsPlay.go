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
	Article			map[string]*Post
	Lock 			sync.Mutex
}

func (cps *ChildsPlayService) Status() string {
	return "good"
}

func (cps *ChildsPlayService) Execute(data []string, s *website.Session, p *website.Page) string {
	logger.Trace.Println("ChildsPlayService.Execute("+data[0]+", page<"+p.Title+">)")
	switch data[0] {
		case "localPosts": return "hi"
	}
	return ""
}

func (cps *ChildsPlayService) Get(p *website.Page, s *website.Session, data []string) website.Item {
	logger.Debug.Println("ChildsPlayService.Get(page<"+p.Title+">, session<"+s.GetUserName()+">, "+data[0]+")")
	return []website.Item{
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Big Dog","Check out this huge dog","10","10","90","120","bigDog.jpg","bd","0"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Cute Kitten","Super Awe....","10","30","130","90","cuteCat.jpg","ck","1"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Ice Cream","Making ice cream, easy","10","50","130","90","iceCream.jpg","ic","1"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Black Horse","This is a beautiful horse","10","70","110","110","blackHorse.jpg","bh","1"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Funny Kid","Kid trips over pillow","30","70","140","80","funnyKid.jpg","fk","2"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Drone Spying","Using drones to spy on neighbors","50","70","100","80","Drone.jpg","ds","2"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Night Bridge","Beautiful pics of bridges at night","70","70","100","80","Bridge.jpg","nb","2"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Ice Sculpting","There will be a class on ice sculpting in Monroe park","70","50","100","80","IceSculpt.jpg","is","2"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Butterfly","The butterfly park in Bowie has opened","70","30","100","80","Butterfly.jpg","bf","2"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Dance","Salsa Dance club","70","10","100","80","Dance.png","sd","2"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Playhouse","Playhouse givaway","50","10","100","80","Playhouse.jpg","ph","2"},
			struct { Title, Desc, X, Y, W, H, JPG, Link, Age string } {"Petting Zoo","Friendship gardens are hosting a petting zoo.","30","10","100","80","PettingZoo.jpg","pz","2"},}
}

func CreateChildsPlayService() *ChildsPlayService {
	logger.Debug.Println("CreateChildsPlayService()")
	cps := ChildsPlayService{Lock: sync.Mutex{}, Article:make(map[string]*Post)}
	keys := make([]string, len(Families))
	i := 0
	for k, _ := range Families {
		keys[i] = k
	    i++
	}
	cps.Article["bd"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"bigDog.jpg", Text:"Check out this huge dog!"}
	cps.Article["ck"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"cuteCat.jpg", Text:"Super Awe...."}
	cps.Article["ic"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"iceCream.jpg", Text:"Making ice cream, easy"}
	cps.Article["bh"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"blackHorse.jpg", Text:"This is a beautiful horse"}
	cps.Article["fk"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"funnyKid.jpg", Text:"Kid trips over pillow"}
	cps.Article["ds"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"Drone.jpg", Text:"Using drones to spy on neighbors"}
	cps.Article["nb"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"Bridge.jpg", Text:"Beautiful pics of bridges at night"}
	cps.Article["is"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"IceSculpt.jpg", Text:"There will be a class on ice sculpting in Monroe park"}
	cps.Article["bf"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"Butterfly.jpg", Text:"The butterfly park in Bowie has opened"}
	cps.Article["sd"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"Dance.jpg", Text:"Salsa Dance club"}
	cps.Article["ph"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"Playhouse.jpg", Text:"Playhouse givaway"}
	cps.Article["pz"] = &Post{Author: Families[keys[rand.Intn(len(keys))]].Parent[0], Pic:"PettingZoo.jpg", Text:"Friendship gardens are hosting a petting zoo."}
	return &cps
}