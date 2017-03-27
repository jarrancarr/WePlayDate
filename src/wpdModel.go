package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/service"
)

type Person struct {
	Name       []string
	DOB        time.Time
	Male       bool
	Email      string
	Admin      bool
	Questions  []Challenge
	ICan       map[*Skill]int8
	IDo        map[*Activity]int8
	Likes      []string
	Dislikes   []string
	Profile    string
	ProfilePic string
	Picture    []string
}

type Group struct {
	member     []*Family
	Circle     map[string]*Group
	Permission map[string]bool
}

type Event struct {
	Title, Details string
	Attendees      []*Family
	Host           *Family
	Time           time.Time
	Duration       time.Duration
	Where          string
}

type Region struct {
	Name       string
	Lat, Long  float32
	Lounge     *service.Room
	Article    map[string]*Post
	Activities []Event
}

type Challenge struct {
	Phrase, Reply string
}

type Message struct {
	From          *Family
	CC            []*Family
	Subject, Body string
}

type Comment struct {
	From *Person
	Text string
}

type Post struct {
	Author      *Person
	User        string
	Pic         string
	Title, Text string
	Blog        []*Post
}

type Family struct {
	Login                  *website.Account
	Parent, Child          []*Person
	Outer                  *Group
	Zip, Buzzword, Turnoff []string
	Profile, ProfilePic    string
	MailBox                map[string][]Message
	Album                  map[string]string
	Comments               map[string][]Comment
}

type Activity struct {
	What      string
	Component []*Activity
	Required  map[int]map[*Skill]int
}

type Skill struct {
	What       string
	Experience map[int8]string
}

func (p *Person) Sex() string {
	age := time.Since(p.DOB)
	if age.Hours() > float64(24*365*18) {
		if p.Male {
			return "Dad"
		}
		return "Mom"
	}
	if p.Male {
		return "Boy"
	}
	return "Girl"
}

func (p *Person) Age() string {
	age := time.Since(p.DOB)
	if age.Hours() > float64(24*365*3) {
		return fmt.Sprintf("%d", int(age.Hours()/(24*365)))
	}
	return fmt.Sprintf("%d months", int(age.Hours()/(24*30)))
}

func (p *Person) FullName() string {
	return strings.Join(p.Name, " ")
}

func (f *Family) Name() string {
	name := f.Parent[0].Name[0]
	for _, p := range f.Parent[1:] {
		name += ", " + p.Name[0]
	}
	for _, c := range f.Child {
		name += ", " + c.Name[0]
	}
	name += " " + f.Parent[0].Name[1]
	return name
}

func (f *Family) String() string {
	return f.Name() + ": " + f.Login.User
}

var (
	Run        = Skill{"Runner", map[int8]string{4: "Jogger", 9: "Runner", 16: "5k", 25: "13.1", 49: "26.2"}}
	Jump       = Skill{"Jump", map[int8]string{1: "hop", 4: "leap", 9: "dive", 25: "expert", 49: "pro"}}
	Sprint     = Skill{"Sprint", map[int8]string{1: "sprint 100", 4: "sprint 500", 9: "sprint 2000", 25: "expert", 49: "pro"}}
	Hike       = Skill{"Hiker", map[int8]string{1: "1-2 Mile", 4: "4-6 Mile", 9: "10 Mile", 16: "Appalatian Trail", 25: "World Traveler"}}
	Kick       = Skill{"Kick", map[int8]string{1: "soccer ball ~10ft", 4: "dropkick ~40ft", 9: "launch ~80ft to a single person", 25: "punter ~100ft"}}
	BallHandle = Skill{"Handle a Ball", map[int8]string{1: "beginner", 4: "fair", 9: "good", 25: "expert", 49: "pro"}}
	Shoot      = Skill{"Shoot the ball", map[int8]string{1: "beginner", 4: "fair", 9: "good", 25: "expert", 49: "pro"}}
	Attack     = Skill{"Attack the ball", map[int8]string{1: "beginner", 4: "fair", 9: "good", 25: "expert", 49: "pro"}}

	Goalie = Activity{"Goalie", nil, map[int]map[*Skill]int{
		1: map[*Skill]int{&Jump: 4, &Kick: 4, &BallHandle: 4},
		4: map[*Skill]int{&Jump: 9, &Kick: 9, &BallHandle: 9, &Run: 4}}}
	Striker = Activity{"Striker", nil, map[int]map[*Skill]int{
		1: map[*Skill]int{&Run: 9, &Kick: 4, &BallHandle: 4},
		4: map[*Skill]int{&Run: 16, &Kick: 9, &BallHandle: 9, &Shoot: 4, &Sprint: 4}}}
	Fullback = Activity{"Fullback", nil, map[int]map[*Skill]int{
		1: map[*Skill]int{&Run: 4, &Kick: 9, &BallHandle: 9},
		4: map[*Skill]int{&Run: 9, &Kick: 16, &BallHandle: 9, &Attack: 4}}}
	Halfback = Activity{"Halfback", nil, map[int]map[*Skill]int{
		1: map[*Skill]int{&Run: 9, &Kick: 9, &BallHandle: 9},
		4: map[*Skill]int{&Run: 16, &Kick: 16, &BallHandle: 9, &Attack: 4, &Sprint: 4}}}
	Soccer = Activity{"Soccer", []*Activity{&Goalie, &Striker, &Fullback, &Halfback}, nil}

	Jarran = Person{Name: []string{"Jarran", "Carr"}, DOB: time.Date(1971, 8, 4, 0, 0, 0, 0, time.UTC), Male: true,
		Email: "jarran.carr@gmail.com", Admin: true}
	Jamie = Person{Name: []string{"Jamie", "Carr"}, DOB: time.Date(1972, 2, 12, 0, 0, 0, 0, time.UTC), Male: false,
		Email: "jamiesgems@bellsouth.com"}
	Logan = Person{Name: []string{"Logan", "Carr"}, DOB: time.Date(2015, 5, 19, 0, 0, 0, 0, time.UTC), Male: true,
		Email: ""}
	Andy = Person{Name: []string{"Andy", "Knight"}, DOB: time.Date(1972, 3, 26, 0, 0, 0, 0, time.UTC), Male: true,
		Email: ""}
	Deanna = Person{Name: []string{"Deanna", "Knight"}, DOB: time.Date(1963, 3, 24, 0, 0, 0, 0, time.UTC), Male: false,
		Email: ""}
	AJ = Person{Name: []string{"Andy", "Knight", "Jr."}, DOB: time.Date(2000, 11, 24, 0, 0, 0, 0, time.UTC), Male: true,
		Email: ""}
	maleNames   = []string{"Alexader", "Andrew", "Anthony", "Adam", "Aaron", "Brian", "Bill", "Brandon", "Benjamin", "Cameron", "Charles", "Christopher", "Corwin", "Debra", "Damon", "Donald", "Daniel", "David", "Dennis", "Douglas", "Edward", "Eric", "Frank", "Fred", "Greggory", "Gary", "George", "Henry", "Hank", "Ivan", "Jacob", "Jack", "Jason", "Jerry", "Jeffery", "Joseph", "Joshua", "James", "John", "Jose", "Kyle", "Kevin", "Larry", "Luke", "Mark", "Michael", "Matthew", "Ned", "Nicholas", "Oliver", "Oscar", "Patrick", "Peter", "Paul", "Phillip", "Quinn", "Raymond", "Richard", "Ronald", "Robbert", "Roman", "Ryan", "Steven", "Sean", "Samuel", "Scott", "Theodore", "Todd", "Thad", "Thomas", "Timothy", "Tyler", "Udel", "Victor", "William", "Walter", "Y", "Zachary", "Zed"}
	femaleNames = []string{"Alice", "Anne", "Ashley", "Amanda", "Amy", "Anna", "Angela", "Barbara", "Brenda", "Betty", "Carolyn", "Cheryl", "Catherine", "Christine", "Doris", "Cynthia", "Deborah", "Donna", "Edith", "Emma", "Evelyn", "Elizabeth", "Emily", "Fay", "Gloria", "Helen", "Janet", "Jean", "Jessica", "Joyce", "Julie", "Joan", "Judith", "Jennifer", "Kimberly", "Karen", "Kathleen", "Kelly", "Lauren", "Laura", "Lisa", "Linda", "Laura", "Margaret", "Michelle", "Maria", "Melissa", "Mary", "Megan", "Nancy", "Olivia", "Patricia", "Pamela", "Rachel", "Rebecca", "Ruth", "Samantha", "Sandra", "Susan", "Sarah", "Stephenie", "Sharon", "Shirley", "Theresa", "Tammy", "Tiffany", "Virgina", "Vallery", "Vivian", "Victoria", "Venus", "Wendy", "Wanda", "Yvette"}
	familyNames = []string{"Smith", "Murphy", "Lam", "Martin", "Brown", "Roy", "Tremblay", "Lee", "Johnson", "Williams", "Jones", "Miller", "Davis", "Garcia", "Rodriguez", "Wilson", "Martinez", "Anderson", "Taylor", "Thomas", "Hernandez", "Moore", "Jackson", "Thompson", "White", "Lopez", "Gonzolez", "Harris", "Clark", "Lewis", "Robinson", "Walker", "Perez", "Hall", "Young", "Allen", "Sanchez", "Write", "King", "Scott", "Green", "Baker", "Adams", "Nelson", "Hill", "Ramirez", "Campbell", "Mitchell", "Roberts", "Carter", "Phillips", "Evans", "Turner", "Torres", "Parker", "Collins", "Edwards", "Stewart", "Florez", "Morris", "Nguyen", "Rivera", "Cook", "Rodgers", "Morgan", "Peterson", "Cooper", "Reed", "Bailey", "Bell", "Gomez", "Kelly", "Howard", "Ward", "Cox", "Diaz", "Richardson", "Wood", "Watson", "Brooks", "Bennett", "Gray", "James", "Reyes", "Cruz", "Hughes", "Price", "Myers", "Long", "Foster", "Sanders", "Ross", "Morales", "Powell", "Sullivan", "Russell", "Ortiz", "Jenkins", "Gutierrez", "Perry", "Butler", "Barnes", "Fisher", "Saim", "Chan"}

	Families = map[string]*Family{
		"jjlcarr": &Family{Login: &website.Account{[]string{"Carr"}, "jjlcarr", "jcarr48", "jcarr@novetta.com", []*website.Role{website.StandardRoles["basic"]},
			false, time.Now()}, Parent: []*Person{&Jarran, &Jamie}, Child: []*Person{&Logan}, Zip: []string{"20720", "20726"}, Buzzword: []string{"hi", "help"}, Turnoff: []string{"hate"}},
		"adaknight": &Family{Login: &website.Account{[]string{"Knight"}, "adaknight", "aknight96", "", []*website.Role{website.StandardRoles["basic"]},
			false, time.Now()}, Parent: []*Person{&Andy, &Deanna}, Child: []*Person{&AJ}, Zip: []string{"20720"}, Buzzword: []string{"Hi", "Help"}, Turnoff: []string{"hate"}},
	}
	famKeys      []string
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
		"what about a roller skate party?", "The movie was too scary", "brb... baby crying."}
	likes = []string{"Horseback riding", "Mountain bikes", "Sports cars", "Ecological farming", "Urban farming", "Cross country skiing",
		"Roller blading", "Archery", "Hunting/Fishing", "Romance books", "Science fiction", "Historical fiction", "Theatre arts", "Painting",
		"Home maintainance", "Woodworking", "Cabinatery", "Astronomy", "Hot rods and muscle cars", "Photograhy", "Computer games", "Hiking",
		"Bowling", "Beaches", "Fine dining", "Cooking", "Cross-fit", "Chess", "Jousting", "Renessaince festival", "Running", "Dogs and cats",
		"Diving and swimming", "Brewing and winemaking", "Science experiments", "Model trains", "RC planes/heli", "Surfing/water sports",
		"Fantasy Role playing", "Decorating", "Knitting", "Arts and crafts", "Hot air ballooning", "Motorcycles", "Camping", "Boating",
		"Acting", "Amateur radio", "Baton twirling", "Board games", "Book restoration", "Cabaret", "Calligraphy", "Candle making",
		"Coffee roasting", "Coloring", "Computer programming", "Cooking", "Cosplaying", "Couponing", "Creative writing", "Crocheting",
		"Cross-stitch", "Crossword puzzles", "Cryptography", "Dance", "Deep web", "Digital arts", "Do it yourself", "Drama", "Drawing",
		"Electronics", "Embroidery", "Fantasy Sports", "Fashion", "Fishkeeping", "Flower arranging", "Foreign language learning", "Gaming",
		"Genealogy", "Glassblowing", "Gunsmithing", "Homebrewing", "Ice skating", "Jewelry making", "Jigsaw puzzles", "Juggling", "Knapping",
		"Knife making", "Knitting", "Kombucha Brewing", "Lacemaking", "Lapidary", "Leather crafting", "Lego building", "Listening to music",
		"Lockpicking", "Machining", "Macrame", "Magic", "Metalworking", "Model building", "Origami", "Painting", "Pet", "Philately",
		"Plastic embedding", "Playing musical instruments", "Poi", "Pottery", "Puzzles", "Quilling", "Quilting", "Reading", "Scrapbooking",
		"Sculpting", "Sewing", "Singing", "Sketching", "Soapmaking", "Stand-up comedy", "Table tennis", "Tatting", "Taxidermy", "Video gaming",
		"Watching movies", "Watching television", "Web surfing", "Whittling", "Wikipedia editing", "Worldbuilding", "Writing", "Yo-yoing",
		"Air sports", "BASE jumping", "Baseball", "Basketball", "Beekeeping", "Bird watching", "Blacksmithing", "Board sports", "Bodybuilding",
		"Brazilian jiu-jitsu", "Dowsing", "Driving", "Flag football", "Flying", "Flying disc", "Foraging", "Freestyle football", "Gardening",
		"Geocaching", "Handball", "High-Powered Rocketry", "Hooping", "Inline skating", "Kayaking", "Kite flying", "Kitesurfing", "Letterboxing",
		"Metal detecting", "Mountaineering", "Mushroom hunting/Mycology", "Netball", "Orienteering", "Paintball", "Parkour", "Polo", "Rafting",
		"Rappelling", "Rock climbing", "Rugby", "Sailing", "Sand art", "Scouting", "Sculling or Rowing", "Topiary", "Skateboarding",
		"Skimboarding", "Skydiving", "Slacklining", "Snowboarding", "Soccer", "Stone skipping", "Taekwondo", "Urban exploration",
		"Vehicle restoration", "Walking", "Action figure", "Antiquing", "Art collecting", "Book collecting", "Coin collecting",
		"Comic book collecting", "Deltiology", "Die-cast toy", "Element collecting", "Movie and movie memorabilia collecting", "Record collecting",
		"Flower collecting and pressing", "Fossil hunting", "Insect collecting", "Metal detecting", "Stone collecting", "Rock balancing",
		"Sea glass collecting", "Seashell collecting", "Badminton", "Billiards", "Boxing", "Bridge", "Cheerleading", "Color guard", "Curling",
		"Darts", "Debate", "Fencing", "Gymnastics", "Kabaddi", "Marbles", "Martial arts", "Mahjong", "Poker", "Slot car racing", "Table football",
		"Volleyball", "Weightlifting", "Airsoft", "Beach Volleyball", "Breakdancing", "Cricket", "Disc golf", "Exhibition drill", "Field hockey",
		"Figure skating", "Footbag", "Golfing", "Handball", "Judo", "Kart racing", "Knife throwing", "Lacrosse"}
)

func initData() {
	factor := 10
	for i := 0; i < 10*factor; i++ {
		familyName := familyNames[rand.Intn(len(familyNames))]
		mom := Person{Name: []string{femaleNames[rand.Intn(len(femaleNames))], familyName}, DOB: time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28), 0, 0, 0, 0, time.UTC),
			Male: false, Email: "test@email.com"}
		dad := Person{Name: []string{maleNames[rand.Intn(len(maleNames))], familyName}, DOB: time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28), 0, 0, 0, 0, time.UTC),
			Male: true, Email: "test@email.com"}
		var children []*Person
		for i := 0; i < 2; i = rand.Intn(3) {
			children = append(children, &Person{Name: []string{maleNames[rand.Intn(len(maleNames))], familyName},
				DOB:  time.Date(2010+rand.Intn(6), time.Month(rand.Intn(12)), 1+rand.Intn(28), 0, 0, 0, 0, time.UTC),
				Male: i%2 == 0, Email: "test@email.com"})
		}
		userName := strings.ToLower(dad.Name[0][:1] + mom.Name[0][:1] + familyName)
		if Families[userName] != nil {
			userName += fmt.Sprintf("%d", rand.Intn(10000))
		}
		Families[userName] = &Family{Login: &website.Account{[]string{familyName}, userName, "password", userName + "@childsplay.com",
			[]*website.Role{website.StandardRoles["basic"]}, false, time.Now()}, Parent: []*Person{&dad, &mom}, Child: children,
			Zip: []string{fmt.Sprintf("%d", 20710+rand.Intn(20))}, Buzzword: []string{"love"}, Turnoff: []string{"hate"}}
		website.Users = append(website.Users, *Families[userName].Login)
	}
	for i := 0; i < 4*factor; i++ { // single moms
		familyName := familyNames[rand.Intn(len(familyNames))]
		mom := Person{Name: []string{femaleNames[rand.Intn(len(femaleNames))], familyName},
			DOB: time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28), 0, 0, 0, 0, time.UTC), Male: false, Email: "test@email.com"}
		var children []*Person
		for i := 0; i < 2; i = rand.Intn(3) {
			children = append(children, &Person{Name: []string{maleNames[rand.Intn(len(maleNames))], familyName},
				DOB:  time.Date(2010+rand.Intn(6), time.Month(rand.Intn(12)), 1+rand.Intn(28), 0, 0, 0, 0, time.UTC),
				Male: i%2 == 0, Email: "test@email.com"})
		}
		userName := strings.ToLower(mom.Name[0][:2] + familyName)
		if Families[userName] != nil {
			userName += fmt.Sprintf("%d", rand.Intn(10000))
		}
		Families[userName] = &Family{Login: &website.Account{[]string{familyName}, userName, "password", userName + "@childsplay.com",
			[]*website.Role{website.StandardRoles["basic"]}, false, time.Now()}, Parent: []*Person{&mom}, Child: children,
			Zip: []string{fmt.Sprintf("%d", 20710+rand.Intn(20))}, Buzzword: []string{"love"}, Turnoff: []string{"hate"}}
		website.Users = append(website.Users, *Families[userName].Login)
	}
	for i := 0; i < 2*factor; i++ { // single dads
		familyName := familyNames[rand.Intn(len(familyNames))]
		dad := Person{Name: []string{maleNames[rand.Intn(len(femaleNames))], familyName},
			DOB: time.Date(1965+rand.Intn(35), time.Month(rand.Intn(12)), 1+rand.Intn(28), 0, 0, 0, 0, time.UTC), Male: true, Email: "test@email.com"}
		var children []*Person
		for i := 0; i < 2; i = rand.Intn(3) {
			children = append(children, &Person{Name: []string{maleNames[rand.Intn(len(maleNames))], familyName},
				DOB:  time.Date(2010+rand.Intn(6), time.Month(rand.Intn(12)), 1+rand.Intn(28), 0, 0, 0, 0, time.UTC),
				Male: i%2 == 0, Email: "test@email.com"})
		}
		userName := strings.ToLower(dad.Name[0][:2] + familyName)
		if Families[userName] != nil {
			userName += fmt.Sprintf("%d", rand.Intn(10000))
		}
		Families[userName] = &Family{Login: &website.Account{[]string{familyName}, userName, "password", userName + "@childsplay.com",
			[]*website.Role{website.StandardRoles["basic"]}, false, time.Now()}, Parent: []*Person{&dad}, Child: children,
			Zip: []string{fmt.Sprintf("%d", 20710+rand.Intn(20))}, Buzzword: []string{"love"}, Turnoff: []string{"hate"}}
		website.Users = append(website.Users, *Families[userName].Login)
	}

	famKeys = make([]string, len(Families))
	i := 0
	for k, f := range Families {
		if len(f.Parent) == 2 {
			f.ProfilePic = "mf"
		} else {
			if f.Parent[0].Male {
				f.ProfilePic = "f"
			} else {
				f.ProfilePic = "m"
			}
		}
		for _, p := range f.Parent {
			addLikes(p)
			p.ProfilePic = "blank.jpg"
			p.Profile = story()
		}
		for _, p := range f.Child {
			addLikes(p)
			p.ProfilePic = "blank.jpg"
			if p.Male {
				f.ProfilePic += "b"
			} else {
				f.ProfilePic += "g"
			}
			p.Profile = story()
		}
		f.ProfilePic += fmt.Sprintf("%d.jpg", rand.Intn(10))
		famKeys[i] = k
		f.Profile = story()
		i++
	}
}

func addLikes(p *Person) {
	for i := 0; i < 4; i = rand.Intn(5) {
		p.Likes = append(p.Likes, likes[rand.Intn(len(likes))])
	}
}

func simulateCommunity(mss *service.MessageService) {
	logger.Trace.Println()
	for i := 0; i < 10; i++ {
		go activeUser(Families[famKeys[rand.Intn(len(Families))]], mss)
	}
	for {
		go activeUser(Families[famKeys[rand.Intn(len(Families))]], mss)
		time.Sleep(time.Millisecond * 2000)
	}
}

func activeUser(fm *Family, mss *service.MessageService) {
	logger.Trace.Println()
	userSession := website.Session{make(map[string]interface{}), make(map[string]string), true}
	userSession.Data["name"] = fm.Parent[0].FullName()
	userSession.Data["userName"] = fm.Login.User
	userSession.Item["family"] = fm
	acs.Lock.Lock()
	acs.Active[fm.Login.User] = &userSession
	acs.Lock.Unlock()
	mss.Execute([]string{"addRoom", fm.Zip[0]}, &userSession, nil)
	for i := 0; i < 100; i = rand.Intn(101) {
		if rand.Intn(5) == 0 {
			mss.Execute([]string{"post", fm.Zip[0], fm.Login.User, Conversation[rand.Intn(len(Conversation))]}, &userSession, nil)
		} else {
			comment := ""
			if rand.Intn(5) == 0 {
				comment = sentense()
			} else {
				comment = phrase()
			}
			mss.Execute([]string{"post", fm.Zip[0], fm.Login.User, comment}, &userSession, nil)
		}
		time.Sleep(time.Duration(rand.Int31n(10000)) * time.Millisecond)
	}
	mss.Execute([]string{"exitRoom", fm.Zip[0]}, &userSession, nil)
	acs.Lock.Lock()
	delete(acs.Active, fm.Login.User)
	acs.Lock.Unlock()
}

func collectMetrics(cps *ChildsPlayService, mss *service.MessageService, acs *website.AccountService, wpd *website.Site) {
	for {
		cps.Metrics["rooms"] = append(cps.Metrics["rooms"], mss.Metrics("rooms"))
		// cps.Metrics[""] = append(cps.Metrics[""], )
		time.Sleep(time.Millisecond * 5000)
	}
}

func word() string {
	precons := "bcdfghlmnprstwjkvwyz"
	postcons := "bdghmnprtx"
	digraphs := "thchshwhquckph"
	ccons := "bcdfghlmnprstwbcdfghlmnprstwbcdfghlmnprstwjkvwxyz"
	vowel := "aeiou"
	word := ""
	for i := 0; i < 4; {
		i = rand.Intn(13)
		dg := rand.Intn(7)
		switch rand.Intn(25) {
		case 0:
			word += string(vowel[rand.Intn(5)])
		case 1:
			word += string(vowel[rand.Intn(5)]) + "?"
		case 2:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + "?"
		case 3:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)])
		case 4:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + "?"
		case 5:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + string(vowel[rand.Intn(5)]) + "?"
		case 6:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + "?"
		case 7:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + string(ccons[rand.Intn(49)]) + "?"
		case 8:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + string(precons[rand.Intn(20)]) + "e"
		case 9:
			word += string(digraphs[dg*2:dg*2+2]) + string(vowel[rand.Intn(5)]) + "?"
		case 10:
			word += string(digraphs[dg*2:dg*2+2]) + string(vowel[rand.Intn(5)])
		case 11:
			word += string(digraphs[dg*2:dg*2+2]) + string(vowel[rand.Intn(5)]) + "?"
		case 12:
			word += string(digraphs[dg*2:dg*2+2]) + string(vowel[rand.Intn(5)]) + string(vowel[rand.Intn(5)]) + "?"
		case 13:
			word += string(digraphs[dg*2:dg*2+2]) + string(vowel[rand.Intn(5)]) + "?"
		case 14:
			word += string(digraphs[dg*2:dg*2+2]) + string(vowel[rand.Intn(5)]) + string(ccons[rand.Intn(49)]) + "?"
		case 15:
			word += string(digraphs[dg*2:dg*2+2]) + string(vowel[rand.Intn(5)]) + string(precons[rand.Intn(20)]) + "e"
		case 16:
			word += string(vowel[rand.Intn(5)])
		case 17:
			word += string(vowel[rand.Intn(5)]) + "?"
		case 18:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + "?"
		case 19:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)])
		case 20:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + "?"
		case 21:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + string(vowel[rand.Intn(5)]) + "?"
		case 22:
			word += string(precons[rand.Intn(20)]) + string(ccons[rand.Intn(49)]) + string(vowel[rand.Intn(5)]) + "?"
		case 23:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + string(ccons[rand.Intn(49)]) + "?"
		case 24:
			word += string(precons[rand.Intn(20)]) + string(vowel[rand.Intn(5)]) + string(precons[rand.Intn(20)]) + "e"
		}
		if i < 4 && word[len(word)-1:] == "?" {
			word = word[:len(word)-1] + string(ccons[rand.Intn(49)])
		} else {
			word = word[:len(word)-1] + string(postcons[rand.Intn(10)])
		}
	}
	return word
}

func phrase() string {
	phrase := ""
	for i := 0; i < 10; i += rand.Intn(9) {
		phrase += word() + " "
	}
	return phrase[:len(phrase)-1]
}

func sentense() string {
	sentense := ""
	for i := 0; i < 30; i += rand.Intn(9) {
		sentense += word() + " "
	}
	sentense = strings.ToUpper(sentense[:1]) + sentense[1:]
	return sentense[:len(sentense)-1] + "."
}

func story() string {
	story := ""
	for i := 0; i < 20; i += rand.Intn(12) {
		story += sentense() + "  "
	}
	return story[:len(story)-2]
}
