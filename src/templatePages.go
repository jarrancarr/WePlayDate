package main

func addTemplatePages() {
	logger.Trace.Println("addTemplatePages()")
	wePlayDate.AddPage("", "head", "")
	wePlayDate.AddPage("header", "header", "")
	wePlayDate.AddPage("", "footer", "")
	wePlayDate.AddPage("", "banner", "")
}