package handlers

import (
	"net/http"

	"github.com/cslemes/heroes-cube-web/internal/templates"
)

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := templates.RenderHome(w); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
