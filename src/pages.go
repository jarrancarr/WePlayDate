package main

func addPages() {
	main := wePlayDate.AddPage("WePlayDate", "main", "/")
	main.AddBypassSiteProcessor("secure")
}