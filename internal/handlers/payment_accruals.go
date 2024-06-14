package handlers

import (
	"fmt"
	"strconv"
	"test/internal/models"
	"time"

	"net/http"
)

func (h *Handler) paymentAccruals(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/paymentAccruals" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum, err := strconv.Atoi(r.URL.Query().Get("Sum"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idAccruals, err := strconv.Atoi(r.URL.Query().Get("idAccruals"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payment := &models.PaymentsDTO{
		Date:   time.Now(),
		Sum:    sum,
		UserId: userId,
	}
	err = h.service.AccountsService.CreatePayment(payment, idAccruals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/usermenu?id=%d", userId), http.StatusSeeOther)

}
