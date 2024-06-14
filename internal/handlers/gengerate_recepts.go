package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"test/internal/models"
	"time"
)

func (h *Handler) generateRecrpts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/generateRecrpts" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	userId, err := strconv.Atoi(r.FormValue("UserId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	last, err := strconv.Atoi(r.FormValue("Last"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	past, err := strconv.Atoi(r.FormValue("Past"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	substract, err := strconv.Atoi(r.FormValue("Substract"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sum, err := strconv.Atoi(r.FormValue("Sum"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.service.UserService.GetUserByID(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accr := models.AccrualsDTO{
		Past:      past,
		Last:      last,
		Substract: substract,
		Sum:       sum,
		Addr:      user.Addr,
		Name:      user.Name,
	}
	err = h.service.ReceptsService.CreateAccruals(accr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/usermenu?id=%d", userId), http.StatusSeeOther)
}

func (h *Handler) generateRes(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/generateRes" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	userId, err := strconv.Atoi(r.FormValue("UserId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sum, err := strconv.Atoi(r.FormValue("Sum"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.service.UserService.GetUserByID(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dateOnly := r.FormValue("Date")

	date, err := time.Parse("2006-01-02", dateOnly)
	pay := models.PaymentsDTO{
		Sum:  sum,
		Date: date,
		Addr: user.Addr,
		Name: user.Name,
	}
	err = h.service.ReceptsService.CreateRes(pay)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/usermenu?id=%d", userId), http.StatusSeeOther)
}
