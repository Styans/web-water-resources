package handlers

import (
	"net/http"
	"test/internal/render"
)

func (h *Handler) tariffs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tariffs" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	h.templates.Render(w, r, "tariffs.page.html", &render.PageData{
		Topic: "Тарифы",
	})
}
