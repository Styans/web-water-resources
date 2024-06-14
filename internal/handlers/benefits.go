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
	t, err := h.service.TariffsService.GetAlltariffs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)

		return
	}
	h.templates.Render(w, r, "benefits.page.html", &render.PageData{
		Topic:   "Льготы",
		Tariffs: *t,
	})
}
