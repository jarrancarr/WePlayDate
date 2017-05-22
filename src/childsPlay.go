package main

import (
	"fmt"
	//"net/http"
	//"errors"
	//"time"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/jarrancarr/website"
)

type ChildsPlayService struct {
	Places    map[string]*Region
	Metrics   map[string][]int
	Thumbnail map[string][]byte
	Lock      sync.Mutex
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
		down := 0
		cfg := s.Item["family"].(*Family).Item[data[1]].(ModalPlacement)
		wid := cfg.W
		if wid == 0 {
			wid = 400
		}
		hit := cfg.H
		if hit == 0 {
			hit = 270
		}
		wid = (wid - 100) * 100 / wid
		hit, down = (hit-130)*100/hit, 2000/hit
		deg := 36
		arc := 360 / deg
		posX := make([]string, deg)
		posY := make([]string, deg)
		for i := 0; i < deg; i += 1 {
			posX[i] = fmt.Sprintf("%d", int(float64(wid/2)+math.Sin(float64(arc*i))*float64(wid/2)))
			posY[i] = fmt.Sprintf("%d", down+int(float64(hit/2)+math.Cos(float64(arc*i))*float64(hit/2)))
			//posX[i] = fmt.Sprintf("%d",int(float64(wid/2)+math.Sin(float64(arc*i))*float64(wid/4)+math.Sin(float64(3*arc*i))*float64(wid/12)))
			//posY[i] = fmt.Sprintf("%d",int(float64(hit/2)+math.Cos(float64(arc*i))*float64(hit/4)-math.Cos(float64(3*arc*i))*float64(hit/12)))
		}
		articles := []website.Item{}
		if cps.Places[data[1]] == nil {
			return nil
		}
		for n, art := range cps.Places[data[1]].Article {
			i := rand.Intn(len(posX))
			logger.Trace.Println(art)
			articles = append(articles, struct{ Title, Desc, X, Y, W, H, JPG, Link, Age string }{
				art.Title, art.Text, posX[i], posY[i], "100", "100", art.Pic, n, "0"})
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

func (cps *ChildsPlayService) MakePlace(where string, lat, long float32) {
	cps.Places[where] = &Region{Name: where, Article: make(map[string]*Post), Center: Coordinates{lat, long}}
	for name, place := range cps.Places {
		if name != where && dist(cps.Places[where], cps.Places[name]) < 0.2 {
			cps.Places[where].AddNeighbor(place)
		}
	}
}

func (cps *ChildsPlayService) AddThumb(key string, data []byte) {
	if cps.Thumbnail == nil {
		cps.Thumbnail = make(map[string][]byte)
	}
	cps.Thumbnail[key] = data
}

func degreesToRadians(degrees float32) float32 {
	return degrees * 3.14159 / 180
}

func dist(r1, r2 *Region) float32 {
	earthRadiusKm := 6371
	dLat := degreesToRadians(r2.Center.Lat - r1.Center.Lat)
	dLon := degreesToRadians(r2.Center.Long - r1.Center.Long)
	lat1 := degreesToRadians(r1.Center.Lat)
	lat2 := degreesToRadians(r2.Center.Lat)
	a := math.Sin(float64(dLat/2))*math.Sin(float64(dLat/2)) + math.Sin(float64(dLon/2))*math.Sin(float64(dLon/2))*math.Cos(float64(lat1))*math.Cos(float64(lat2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return float32(earthRadiusKm) * float32(c)
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

	cps.MakePlace("20720", 39.005867, -76.789025)
	cps.MakePlace("20721", 38.924936, -76.790291)
	cps.MakePlace("20722", 38.936522, -76.949665)
	cps.MakePlace("20723", 39.141052, -76.864416)
	cps.MakePlace("20724", 39.092186, -76.792573)
	cps.MakePlace("20725", 39.108165, -76.850504)
	cps.MakePlace("20726", 39.093793, -76.858412)
	for _, aName := range []string{"bd", "ck", "ic", "bh", "fk", "ds", "nb"} {
		cps.Places["20720"].Article[aName] = makeArticle()
	}
	for _, aName := range []string{"is", "bf", "sd", "ph", "pz", "cn", "cl"} {
		cps.Places["20726"].Article[aName] = makeArticle()
	}
	cps.Places["20720"].AddPlace("Bowie State University", &Region{Name: "Bowie State University", Center: Coordinates{39.02, -76.76}})
	cps.Places["20720"].AddPlace("Bowie Montessori", &Region{Name: "Bowie Montessori", Center: Coordinates{39.0, -76.755}})
	cps.Places["20720"].AddPlace("Rockledge Elementary", &Region{Name: "Rockledge Elementary", Center: Coordinates{39.00, -76.76}})
	cps.Places["20720"].AddPlace("Whitehall Elementary", &Region{Name: "Whitehall Elementary", Center: Coordinates{38.99, -76.75}})
	cps.Places["20720"].AddPlace("Sanuel Ogle Middle", &Region{Name: "Sanuel Ogle Middle", Center: Coordinates{}})
	cps.Places["20720"].AddPlace("Tulip Grove", &Region{Name: "Tulip Grove", Center: Coordinates{38.98, -76.75}})
	cps.Places["20720"].AddPlace("Bowie High", &Region{Name: "Bowie High", Center: Coordinates{38.98, -76.74}})
	cps.Places["20720"].AddPlace("St Matthews", &Region{Name: "St Matthews", Center: Coordinates{38.977, -76.745}})
	cps.Places["20720"].AddPlace("The Goddard", &Region{Name: "The Goddard", Center: Coordinates{38.97, -76.76}})
	cps.Places["20720"].AddPlace("YMCA", &Region{Name: "YMCA", Center: Coordinates{38.98, -76.75}})
	cps.Places["20720"].MapPhoto = "20720.png"
	cps.Places["20726"].MapPhoto = "20726.png"
	return &cps
}

var pics = []string{"bigDog.jpg", "cuteCat.jpg", "iceCream.jpg", "blackHorse.jpg", "funnyKid.jpg", "Drone.jpg", "Bridge.jpg", "IceSculpt.jpg", "Butterfly.jpg", "Dance.png", "Playhouse.jpg", "PettingZoo.jpg", "City.jpg", "Country.jpg"}

//var pics = []string{"bigDog.bpg", "cuteCat.bpg", "iceCream.bpg", "blackHorse.bpg", "funnyKid.bpg", "Drone.bpg", "Bridge.bpg", "IceSculpt.bpg", "Butterfly.bpg", "Dance.bpg", "Playhouse.bpg", "PettingZoo.bpg", "City.bpg", "Country.bpg"}
var text = []string{"Check out this huge dog!", "Super Awe....", "Making ice cream, easy", "This is a beautiful horse", "Kid trips over pillow", "Using drones to spy on neighbors", "Beautiful pics of bridges at night",
	"There will be a class on ice sculpting in Monroe park", "The butterfly park in Bowie has opened", "Salsa Dance club", "Playhouse givaway", "Friendship gardens are hosting a petting zoo.", "City life at night", "Nothing like country living."}

func makeArticle() *Post {
	rd := rand.Intn(len(famKeys))
	fam := Families[famKeys[rd]]
	if fam == nil {
		os.Exit(1)
	}
	par := rand.Intn(len(fam.Parent))
	art := rand.Intn(len(pics))
	return &Post{Author: fam.Parent[par], User: fam.Login.User, Pic: pics[art], Text: text[art]}
}
