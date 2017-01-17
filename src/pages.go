package main

import (
	"github.com/jarrancarr/website/html"
)

func addPages() {
	main := wePlayDate.AddPage("WePlayDate", "main", "/")
	main.AddBypassSiteProcessor("secure")
	main.Html.Add("circleMenuItem", html.NewTag("circle id==login cx==${CX} cy==10 r==50 fill==#93e stroke==#222 stroke-width==1 fill-opacity==0.8"))
	main.Html.Add("circleMenuItem", html.NewTag("text x==${TX} y==25 font-family==Verdana font-size==26 fill==#ee9 ${LABEL}"))
}
