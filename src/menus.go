package main

import (
	"net/http"

	"github.com/jarrancarr/website"
)

func addMenus() {
	wePlayDate.AddParamTriggerHandler("language", chooseLanguage)
	wePlayDate.Data["languages"] = []string{"en", "cs", "fr", "sp"}
}

func chooseLanguage(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	s.Data["language"] = p.Param["language"]
	http.Redirect(w, r, s.Data["navigation"], 302)
	return "", nil
}
