package handlers

import (
	"net/http"
	"test/internal/render"
)

func (h *Handler) benefits(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/benefits" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	h.templates.Render(w, r, "benefits.page.html", &render.PageData{
		Topic: "Льготы",
	})
}
