package main

func addTemplatePages() {
	logger.Trace.Println("addTemplatePages()")
	weePlayDate.AddPage("", "head", "")
	weePlayDate.AddPage("header", "header", "")
	weePlayDate.AddPage("", "footer", "")
	weePlayDate.AddPage("", "banner", "")
	weePlayDate.AddPage("", "family", "")
	weePlayDate.AddPage("", "dashboard", "")
	weePlayDate.AddPage("", "wallchart", "")
	weePlayDate.AddPage("", "modal/person", "")
	weePlayDate.AddPage("", "modal/family", "")
	weePlayDate.AddPage("", "modal/welcome", "")
	weePlayDate.AddPage("", "modal/info", "")
	weePlayDate.AddPage("", "modal/article", "")
	weePlayDate.AddPage("", "modal/apply", "")
	weePlayDate.AddPage("", "modal/login", "")
	weePlayDate.AddPage("", "modal/errors", "")
	weePlayDate.AddPage("", "modal/album", "")
	weePlayDate.AddPage("", "modal/alert", "")
	weePlayDate.AddPage("", "modal/schedual", "")
	weePlayDate.AddPage("", "wallpaper", "")
}
