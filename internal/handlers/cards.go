package handlers

import (
	"github.com/cslemes/hc-web/internal/views"
	"github.com/labstack/echo/v4"
)

type Card struct {
	Race  string
	Image string
}

func Cards() echo.HandlerFunc {
	return func(c echo.Context) error {

		cards := []Card{
			{Race: "Dwarf Mage ", Image: "/static/DwarfMage.png"},
			{Race: "Elf Warrior", Image: "/static/ElfWarrior.png"},
			{Race: "Human Cleric", Image: "/static/HumanCleric.png"},
			{Race: "Elf Warrior", Image: "/static/ElfWarriorFem.png"},
			{Race: "Poney Thief", Image: "/static/PoneyThief.png"},
			{Race: "Hobbit Bard", Image: "/static/HobbitBard.png"},
		}

		data := map[string]interface{}{
			"Cards": cards,
		}

		if err := views.NewTemplates().Render(c.Response(), "cards", data, c); err != nil {
			return c.String(500, "Internal Server Error")
		}

		return nil
	}
}
