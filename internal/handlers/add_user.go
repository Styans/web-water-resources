package handlers

import (
	"net/http"
	"test/internal/models"
	"test/internal/render"
)

func (h *Handler) newUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/add_user" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		h.templates.Render(w, r, "newuser.page.html", &render.PageData{
			Topic: "Льготы",
		})
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := &models.User{
			Name:       r.FormValue("name"),
			Secondname: r.FormValue("secondname"),
			Patronymic: r.FormValue("patronymic"),
			Benefits:   r.FormValue("benefits"),

			Districts: r.FormValue("districts"),
			Addr:      r.FormValue("addr"),
		}

		err = h.service.UserService.CreateUser(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.templates.Render(w, r, "newuser.page.html", &render.PageData{
			Topic: "Льготы",
		})
	}

}
