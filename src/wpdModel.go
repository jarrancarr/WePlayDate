package main

import (
	//"fmt"
	"time"
	"github.com/jarrancarr/website"
)

type Person struct {
	Name	[]string
	DOB		time.Time
	Male	bool
	Email	string
	Admin	bool
	Questions []Challenge
	ICan	map[*Skill]int8
	IDo		map[*Activity]int8
}

type Group struct {
	member []Family
	Circle []Group
	Permission map[string]bool
}

type Challenge struct {
	Phrase, Reply  string
}

type Family struct {
	Login	*website.Account
	Parent	[]*Person
	Child	[]*Person
	Outer	*Group
	Zip		string
}

type Activity struct {
	What	string
	Component []*Activity
	Required map[int]map[*Skill]int
}

type Skill struct {
	What	string
	Experience map[int8]string
}

var (
	Run = Skill{"Runner",map[int8]string{4:"Jogger", 9:"Runner", 16:"5k", 25:"13.1", 49:"26.2"}}
	Jump = Skill{"Jump", map[int8]string{1:"hop", 4:"leap", 9:"dive", 25:"expert", 49:"pro"}}
	Sprint = Skill{"Sprint", map[int8]string{1:"sprint 100", 4:"sprint 500", 9:"sprint 2000", 25:"expert", 49:"pro"}}
	Hike = Skill{"Hiker",map[int8]string{1:"1-2 Mile", 4:"4-6 Mile", 9:"10 Mile", 16:"Appalatian Trail", 25:"World Traveler"}}
	Kick = Skill{"Kick", map[int8]string{1:"soccer ball ~10ft", 4:"dropkick ~40ft", 9:"launch ~80ft to a single person", 25:"punter ~100ft"}}
	BallHandle = Skill{"Handle a Ball", map[int8]string{1:"beginner", 4:"fair", 9:"good", 25:"expert", 49:"pro"}}
	Shoot = Skill{"Shoot the ball", map[int8]string{1:"beginner", 4:"fair", 9:"good", 25:"expert", 49:"pro"}}
	Attack = Skill{"Attack the ball", map[int8]string{1:"beginner", 4:"fair", 9:"good", 25:"expert", 49:"pro"}}
	
	
	Goalie = Activity{"Goalie", nil, map[int]map[*Skill]int{
		1:map[*Skill]int{&Jump:4, &Kick:4, &BallHandle:4}, 
		4:map[*Skill]int{&Jump:9, &Kick:9, &BallHandle:9, &Run:4}}}
	Striker = Activity{"Striker", nil, map[int]map[*Skill]int{
		1:map[*Skill]int{&Run:9, &Kick:4, &BallHandle:4},
		4:map[*Skill]int{&Run:16, &Kick:9, &BallHandle:9, &Shoot:4, &Sprint:4}}}
	Fullback = Activity{"Fullback", nil, map[int]map[*Skill]int{
		1:map[*Skill]int{&Run:4, &Kick:9, &BallHandle:9},
		4:map[*Skill]int{&Run:9, &Kick:16, &BallHandle:9, &Attack:4}}}
	Halfback = Activity{"Halfback", nil, map[int]map[*Skill]int{
		1:map[*Skill]int{&Run:9, &Kick:9, &BallHandle:9},
		4:map[*Skill]int{&Run:16, &Kick:16, &BallHandle:9, &Attack:4,&Sprint:4}}}
	Soccer = Activity{"Soccer", []*Activity{&Goalie, &Striker, &Fullback, &Halfback}, nil}
	
	Families = map[string]*Family{"Carr":&Family{
		&website.Account{[]string{"Carr"},"CarrFamily","jcarr","jcarr@novetta.com", 
			 []*website.Role{website.StandardRoles["basic"],} , false, time.Now()},
		[]*Person{
			&Person{[]string{"Jarran","Carr"}, time.Date(1971,8,4,0,0,0,0,time.UTC), true, "jarran.carr@gmail.com", true, nil, nil, nil, },
			&Person{[]string{"Jamie","Carr"}, time.Date(1972,2,12,0,0,0,0,time.UTC), true, "jamiesgems@bellsouth.com", true, nil, nil, nil, },
		},
		[]*Person{ &Person{[]string{"Logan","Carr"}, time.Date(2015,5,19,0,0,0,0,time.UTC), true, "", true, nil, nil, nil, }, },
		nil,
		"20720",
	},	}
)