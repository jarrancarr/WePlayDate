package main

import (
	"net/http"
	//"io/ioutil"
	//"errors"
	//"strings"
	"fmt"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/html"
)

func addPages() {
	logger.Trace.Println("addPages()")
	main := weePlayDate.AddPage("WePlayDate", "main", "/")
	main.AddBypassSiteProcessor("secure")
	main.AddInitProcessor(MainInitProcessor)
	main.AddPostHandler("login", LoginPostHandler)
	main.AddPostHandler("apply", RegisterPostHandler)
	acs.FailLoginPage = "/main#failedLogin"
	acs.LogoutPage = "/main#goodBye"

	main.Html.Add("circleMenuItem", html.NewTag("circle id==login cx==${CX} cy==10 r==50 fill==#${FILL} stroke==#222 stroke-width==1 fill-opacity==0.8"))
	main.Html.Add("circleMenuItem", html.NewTag("a xlink:href==#${MODAL}").AppendChild(
		html.NewTag("text x==${TX} y==25 font-family==Verdana font-size==26 fill==#ee9 ${LABEL}")))

	main.Html.Add("pictures", html.NewTag("defs").AppendChild(
		html.NewTag("pattern id==${ID} x==0 y==0 height==1 width==1").AppendChild(
			html.NewTag("image x==${XX} y==${YY} height==${HEIGHT} width==${WIDTH} xlink:href==/img/${ID}.jpg"))))
	main.Html.Add("pictures", html.NewTag("circle id==${ID}-circle cx==${CX}% cy==${CY}% r==${R} fill==url(#${ID}) stroke==#39e stoke-width==16px stroke-opacity==0.5").AppendChild(
		html.NewTag("animate id==fadein-${ID} attributeName==opacity values==${FADE} dur==${DUR}s begin==${ITERATOR}s repeatCount=indefinite")))
	main.AddParam("DateFormat", Date_Format)

	home := weePlayDate.AddPage("home", "home", "/home")
	home.AddAJAXHandler("newRoom", mss.CreateRoomAJAXHandler)
	home.AddAJAXHandler("talks", mss.GetConversationsAJAXHandler)
	home.AddAJAXHandler("message", mss.MessageAJAXHandler)
	home.AddAJAXHandler("exitRoom", mss.ExitRoomAJAXHandler)
	home.AddAJAXHandler("whoseThere", WhoseThereAjaxHandler)
	home.AddAJAXHandler("familyProfile", GetFamilyProfileAjaxHandler)
	home.AddAJAXHandler("personProfile", GetPersonProfileAjaxHandler)
	home.AddAJAXHandler("article", GetArticleAjaxHandler)
	home.AddAJAXHandler("editUpdate", UpdateFieldAjaxHandler)
	home.AddAJAXHandler("getMap", GetMapAjaxHandler)

	home.AddPostHandler("logout", acs.LogoutPostHandler)
	home.AddPostHandler("selectFamilyMember", SelectFamilyMember)
	home.AddPostHandler("edit", EditDataPostHandler)

	home.Html.Add("circleMenuItem", html.NewTag("circle id==login cx==${CX} cy==10 r==50 fill==#${FILL} stroke==#222 stroke-width==1 fill-opacity==0.8"))
	home.Html.Add("circleMenuItem", html.NewTag("a xlink:href==#${MODAL}").AppendChild(
		html.NewTag("text x==${TX} y==25 font-family==Verdana font-size==26 fill==#ee9 ${LABEL}")))
	home.AddParam("newRoomSetup", "setup: newRoomName = $('#newRoom-name').val(); newRoomPass = $('#newRoom-pass').val(); ")
	home.AddParam("newRoomSuccess", `success:
		var ul = $( "<ul/>", {"class": "ptButton"}); 
		var obj = JSON.parse(data); 
		$("#roomList").empty(); 
		$("#roomList").append(ul); 
		$.each(obj, function(val, i) { 
			item = $(document.createElement('button')).text( val + '  ' + i + ' occupance' ); 
			item.attr("class", "ptButton"); 
			item.attr("onclick","enterRoom('"+val+"')"); 
			ul.append( item ); 
			ul.append( '<br/>' ); 
		});
	`)

	admin := weePlayDate.AddPage("", "admin", "/admin")
	admin.AddBypassSiteProcessor("secure")
	admin.AddInitProcessor(InitAdminPage)
}

func InitAdminPage(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	if s.Data["metrics"] == "" {
		s.Data["metrics"] = "{'famPage'=0,'roomPage'=0,'userPage'=0}"
	}
	p.Param["#families"] = fmt.Sprintf("%d", len(Families))
	return "ok", nil
}
