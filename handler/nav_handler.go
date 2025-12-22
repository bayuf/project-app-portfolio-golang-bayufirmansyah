package handler

import "github.com/bayuf/project-app-portfolio-golang-bayufirmansyah/model"

func BuildNav(current string) []model.NavItem {
	return []model.NavItem{
		{
			Title:  "Home",
			URL:    "/",
			Active: current == "Home",
		},
		{
			Title:  "About",
			URL:    "/about",
			Active: current == "About",
		},
		{
			Title:  "Services",
			URL:    "/services",
			Active: current == "Services",
		},
		{
			Title:  "Portofolio",
			URL:    "/portofolio",
			Active: current == "Portofolio",
		},
		{
			Title:  "Contact",
			URL:    "/contact",
			Active: current == "Contact",
		},
	}
}
