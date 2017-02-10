package main

import (
	"fmt"
	"time"
	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/service"
	"math/rand"
	"strings"
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
	member []*Family
	Circle map[string]*Group
	Permission map[string]bool
}

type Challenge struct {
	Phrase, Reply  string
}

type Message struct {
	From 	*Family
	CC		[]*Family
	Subject, Body string
}

type Family struct {
	Login	*website.Account
	Parent	[]*Person
	Child	[]*Person
	Outer	*Group
	Zip		[]string
	Buzzword []string
	Turnoff []string
}

type Region struct {
	Name		string
	Lat, Long	float32
	room		*service.Room
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

func (p *Person) FullName() string {
	return strings.Join(p.Name, " ")
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
	
	Jarran = Person{[]string{"Jarran","Carr"}, time.Date(1971,8,4,0,0,0,0,time.UTC), true, "jarran.carr@gmail.com", true, nil, nil, nil, }
	Jamie = Person{[]string{"Jamie","Carr"}, time.Date(1972,2,12,0,0,0,0,time.UTC), false, "jamiesgems@bellsouth.com", true, nil, nil, nil, }
	Logan = Person{[]string{"Logan","Carr"}, time.Date(2015,5,19,0,0,0,0,time.UTC), true, "", true, nil, nil, nil, }
	Andy = Person{[]string{"Andy","Knight"}, time.Date(1972,3,26,0,0,0,0,time.UTC), true, "jarran.carr@gmail.com", true, nil, nil, nil, }
	Deanna = Person{[]string{"Deanna","Knight"}, time.Date(1963,3,24,0,0,0,0,time.UTC), false, "jamiesgems@bellsouth.com", true, nil, nil, nil, }
	AJ = Person{[]string{"Andy","Knight", "Jr."}, time.Date(2000,11,24,0,0,0,0,time.UTC), true, "", true, nil, nil, nil, }
	maleNames = []string{"Alexader", "Andrew", "Anthony", "Adam", "Aaron", "Brian", "Bill", "Brandon", "Benjamin", "Cameron", "Charles", "Christopher", "Debra", "Damon", "Donald", "Daniel", "David", "Dennis", "Douglas", "Edward", "Eric", "Frank", "Fred", "Greggory", "Gary", "George", "Henry", "Ivan", "Jacob", "Jack", "Jason", "Jerry", "Jeffery", "Joseph", "Joshua", "James", "John", "Jose", "Kyle", "Kevin", "Larry", "Mark", "Michael", "Matthew", "Ned", "Nicholas", "Oliver", "Patrick", "Peter", "Paul", "Quinn", "Raymond", "Richard", "Ronald", "Robbert", "Ryan", "Steven", "Sean", "Samuel", "Scott", "Todd", "Thad", "Thomas", "Timothy", "Tyler", "Udel", "Victor", "William", "Walter", "Y", "Zachary", "Zed" }
	femaleNames = []string{"Alice", "Anne", "Ashley", "Amanda", "Amy", "Anna", "Angela", "Barbara", "Brenda", "Betty", "Carolyn", "Cheryl", "Catherine", "Christine", "Doris", "Cynthia", "Deborah", "Donna", "Edith", "Emma", "Evelyn", "Elizabeth", "Emily", "Fay", "Gloria", "Helen", "Janet", "Jean", "Jessica", "Joyce", "Julie", "Joan", "Judith", "Jennifer", "Kimberly","Karen","Kathleen", "Kelly", "Lauren", "Laura", "Lisa","Linda","Laura", "Margaret","Michelle", "Maria", "Melissa", "Mary","Megan",	"Nancy", "Olivia", "Patricia", "Pamela", "Rachel","Rebecca","Ruth", "Samantha", "Sandra","Susan","Sarah","Stephenie","Sharon","Shirley", "Theresa", "Tammy", "Tiffany", "Virgina","Vallery","Vivian","Victoria","Venus","Wendy","Wanda","Yvette", }
	familyNames = []string{"Smith", "Murphy", "Lam", "Martin", "Brown", "Roy", "Tremblay", "Lee", "Johnson", "Williams", "Jones", "Miller", "Davis", "Garcia", "Rodriguez", "Wilson", "Martinez", "Anderson", "Taylor", "Thomas", "Hernandez", "Moore", "Jackson", "Thompson", "White", "Lopez", "Gonzolez", "Harris", "Clark", "Lewis", "Robinson", "Walker", "Perez", "Hall", "Young", "Allen", "Sanchez", "Write", "King", "Scott", "Green", "Baker", "Adams", "Nelson", "Hill", "Ramirez", "Campbell", "Mitchell", "Roberts", "Carter", "Phillips", "Evans", "Turner", "Torres", "Parker", "Collins", "Edwards", "Stewart", "Florez", "Morris", "Nguyen", "Rivera", "Cook", "Rodgers", "Morgan", "Peterson", "Cooper", "Reed", "Bailey", "Bell", "Gomez", "Kelly", "Howard", "Ward", "Cox", "Diaz", "Richardson", "Wood", "Watson", "Brooks", "Bennett", "Gray", "James", "Reyes", "Cruz", "Hughes", "Price", "Myers", "Long", "Foster", "Sanders", "Ross", "Morales", "Powell", "Sullivan", "Russell", "Ortiz", "Jenkins", "Gutierrez", "Perry", "Butler", "Barnes", "Fisher", "Saim", "Chan", }
	
	Families = map[string]*Family{
		"jjlcarr":&Family{&website.Account{[]string{"Carr"},"jjlcarr","jcarr48","jcarr@novetta.com", []*website.Role{website.StandardRoles["basic"], },
		false, time.Now()}, []*Person{ &Jarran,	&Jamie,	}, []*Person{ &Logan, }, nil, []string{"20720"}, []string{"hi", "help"}, []string{"hate"}},
		"adaknight":&Family{&website.Account{[]string{"Knight"},"adaknight","aknight96","", []*website.Role{website.StandardRoles["basic"], }, 
		false, time.Now()},	[]*Person{ &Andy,	&Deanna,	}, []*Person{ &AJ, }, nil, []string{"20720"}, []string{"Hi", "Help"}, []string{"hate"}},	
	}
	
	Conversation = []string{"Hello", "Nice kids", "Be right back", "see you later", "what time", "the park was nice", 
		"tomorrow is better", "I don't know", "Wait till I call you", "there are more at home", "what did you find there", 
		"how was surfing", "were sailing tomorrow", "taking the dogs for a walk", "playing with the cat", "shopping for a dress", 
		"have to do some home repair", "You can't tell him that!", "Were going camping", "cooking on the BBQ", "trying a vegan recipie", 
		"where did you go to school", "how are the schools in that area", "she is the best teacher", "I'm not so sure", 
		"my parents are in town", "there is a swim meet", "I'll see you at hockey practice", "Is you car fixed yet?", "You gave me that apron", 
		"He's talking politics", "whatever, I can make it.", "teaching my kids about life", "Is there a playground at that park?", 
		"where was that persian restaurant?", "Who sang that song?", "What were you planning this weekend?", "When do you have free?", 
		"Why don't you meet me at that place?", "How would you punish your kids for that?", "I can try.", "Well, I have a few tools.", 
		"You'll have to pardon the mess.", "I got all these ingredients...", "I'll knit you another pair.", "That just happened", 
		"I'll show you how to build that", "The plummer is here", "The fix-it guy is at the door.", "I need to learn how to work on my car", 
		"oops... baby is crying", "I got plenty of extra firewood", "I'm trying to decide what color to paint that room", 
		"were having family night", "We had a lot of fun at your house", "The kids have all of those toys", "He is taking a potter class", 
		"If the weather is nice, we'll have a picnic.", "make sure you bring a jacket", "hiking the trails", "bicycling to the hills", 
		"what about a roller skate party?", "The movie was too scary", }
)

func initData() {
	for i:=0; i<1000; i++ {
		familyName := familyNames[rand.Intn(len(familyNames))]
		mom := Person{[]string{femaleNames[rand.Intn(len(femaleNames))], familyName, }, 
			time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28),0,0,0,0,time.UTC), false, "test@email.com", true, nil, nil, nil, }
		dad := Person{[]string{maleNames[rand.Intn(len(maleNames))], familyName, }, 
			time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28),0,0,0,0,time.UTC), true, "test@email.com", true, nil, nil, nil, }
		var children []*Person
		for i:=0;i<2;i = rand.Intn(3) {
			children = append(children,&Person{[]string{maleNames[rand.Intn(len(maleNames))], familyName, }, 
				time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28),0,0,0,0,time.UTC), 
				i%2==0, "test@email.com", true, nil, nil, nil, })
		}
		userName := dad.Name[0][:1] + mom.Name[0][:1] + strings.ToLower(familyName)
		Families[userName] = &Family{&website.Account{[]string{familyName,}, userName, "password", userName+"@childsplay.com", 
			[]*website.Role{website.StandardRoles["basic"], }, false, time.Now()}, []*Person{ &dad, &mom}, children, nil, 
			[]string{fmt.Sprintf("%d",20710+rand.Intn(20))}, []string{"love"}, []string{"hate"}}
		website.Users = append(website.Users,*Families[userName].Login)
	}
	for i:=0; i<400; i++ { // single moms
		familyName := familyNames[rand.Intn(len(familyNames))]
		mom := Person{[]string{femaleNames[rand.Intn(len(femaleNames))], familyName, }, 
			time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28),0,0,0,0,time.UTC), false, "test@email.com", true, nil, nil, nil, }
		var children []*Person
		for i:=0;i<2;i = rand.Intn(3) {
			children = append(children,&Person{[]string{maleNames[rand.Intn(len(maleNames))], familyName, }, 
				time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28),0,0,0,0,time.UTC), 
				i%2==0, "test@email.com", true, nil, nil, nil, })
		}
		userName := mom.Name[0][:2] + strings.ToLower(familyName)
		Families[userName] = &Family{&website.Account{[]string{familyName,}, userName, "password", userName+"@childsplay.com", 
			[]*website.Role{website.StandardRoles["basic"], }, false, time.Now()}, []*Person{ &mom}, children, nil, 
			[]string{fmt.Sprintf("%d",20710+rand.Intn(20))}, []string{"love"}, []string{"hate"}}
		website.Users = append(website.Users,*Families[userName].Login)
	}
	for i:=0; i<200; i++ { // single dads
		familyName := familyNames[rand.Intn(len(familyNames))]
		dad := Person{[]string{femaleNames[rand.Intn(len(femaleNames))], familyName, }, 
			time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28),0,0,0,0,time.UTC), false, "test@email.com", true, nil, nil, nil, }
		var children []*Person
		for i:=0;i<2;i = rand.Intn(3) {
			children = append(children,&Person{[]string{maleNames[rand.Intn(len(maleNames))], familyName, }, 
				time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28),0,0,0,0,time.UTC), 
				i%2==0, "test@email.com", true, nil, nil, nil, })
		}
		userName := dad.Name[0][:2] + strings.ToLower(familyName)
		Families[userName] = &Family{&website.Account{[]string{familyName,}, userName, "password", userName+"@childsplay.com", 
			[]*website.Role{website.StandardRoles["basic"], }, false, time.Now()}, []*Person{ &dad}, children, nil, 
			[]string{fmt.Sprintf("%d",20710+rand.Intn(20))}, []string{"love"}, []string{"hate"}}
		website.Users = append(website.Users,*Families[userName].Login)
	}
}

func simulateCommunity(mss *service.MessageService) {
	logger.Trace.Println()
	famKeys := make([]string, len(Families))
	i := 0
	for k := range(Families) {
		famKeys[i] = k
		i++
	}
	for i := 0;i<10;i++{
		go activeUser(Families[famKeys[rand.Intn(len(Families))]], mss)
	}
	for {
		go activeUser(Families[famKeys[rand.Intn(len(Families))]], mss)
		time.Sleep(time.Millisecond*2000)
	}
}

func activeUser(fm *Family, mss *service.MessageService) {
	logger.Trace.Println()
	userSession := website.Session{make(map[string]interface{}), make(map[string]string)}
	userSession.Data["name"] = fm.Parent[0].FullName()
	userSession.Data["userName"] = fm.Login.User
	userSession.Item["family"] = fm
	acs.Lock.Lock()
	acs.Active[fm.Login.User] = &userSession
	acs.Lock.Unlock()
	mss.Execute([]string{"addRoom","zip-"+fm.Zip[0]}, &userSession, nil)
	for i := 0; i<100; i = rand.Intn(101) {
		mss.Execute([]string{"post", "zip-"+fm.Zip[0], fm.Parent[0].FullName(), "("+fm.Login.User+"):"+Conversation[rand.Intn(len(Conversation))]}, &userSession ,nil)
		time.Sleep(time.Duration(rand.Int31n(10000))*time.Millisecond)
	}
	mss.Execute([]string{"exitRoom","zip-"+fm.Zip[0]}, &userSession, nil)
}