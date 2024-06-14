package handlers

import (
	"net/http"
	"strconv"
	"test/internal/render"
)

func (h *Handler) usermenu(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/usermenu" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {

		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.UserService.GetUserByID(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payments, err := h.service.AccountsService.GetPaymentsByUserID(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accruals, err := h.service.AccountsService.GetAccrualsByUserID(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum, err := h.service.TariffsService.GetSumByName(user.Benefits)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	h.templates.Render(w, r, "usermenu.page.html", &render.PageData{
		Topic:      "Абонент",
		User:       user,
		Paymants:   payments,
		Accruals:   accruals,
		TariffsSum: int(sum),
		Cost:       28,
	})
}
