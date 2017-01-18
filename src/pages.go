package main

import (
	"github.com/jarrancarr/website/html"
)

func addPages() {
	logger.Trace.Println("addPages()")
	main := wePlayDate.AddPage("WePlayDate", "main", "/")
	main.AddBypassSiteProcessor("secure")
	main.Html.Add("circleMenuItem", html.NewTag("circle id==login cx==${CX} cy==10 r==50 fill==#${FILL} stroke==#222 stroke-width==1 fill-opacity==0.8"))
	main.Html.Add("circleMenuItem", html.NewTag("a xlink:href==#${MODAL}").AppendChild(
		html.NewTag("text x==${TX} y==25 font-family==Verdana font-size==26 fill==#ee9 ${LABEL}")))
	
	
	main.Html.Add("pictures", html.NewTag("defs").AppendChild(
		html.NewTag("pattern id==${ID} x==0 y==0 height==1 width==1").AppendChild(
		html.NewTag("image x==${XX} y==${YY} height==${HEIGHT} width==${WIDTH} xlink:href==/img/${ID}.jpg"))))
	main.Html.Add("pictures", html.NewTag("circle id==${ID}-circle cx==${CX}% cy==${CY}% r==${R} fill==url(#${ID}) stroke==#39e stoke-width==16px stroke-opacity==0.5").AppendChild(
		html.NewTag("animate id==fadein-${ID} attributeName==opacity values==${FADE} dur==${DUR}s begin==${ITERATOR}s repeatCount=indefinite")))
}
