package handlers

import (
	"net/http"
	"test/internal/render"
)

func (h *Handler) usermenu(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/usermenu" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	users, err := h.service.UserService.GetAllUsers()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.templates.Render(w, r, "usermenu.page.html", &render.PageData{
		Topic: "Абонент",
		Users: users,
	})
}
