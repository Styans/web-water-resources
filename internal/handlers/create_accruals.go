package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"test/internal/models"
	"time"
)

func (h *Handler) createAccruals(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createaccual" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Post", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(r.FormValue("UserId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// layout := "2006-01-02 15:04:05"

	// date := time.Now().Format("2006-01-02")

	past, err := strconv.Atoi(r.FormValue("past"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	last, err := strconv.Atoi(r.FormValue("last"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	accrual := &models.Accounts{
		Date:   time.Now(),
		Past:   past,
		Last:   last,
		UserId: userId,
	}

	err = h.service.AccountsService.CreateAccruals(accrual)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	// user, err := h.service.UserService.GetUserByID(userId)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println("8")
	// payments, err := h.service.AccountsService.GetPaymentsByUserID(userId)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println("9")
	// accruals, err := h.service.AccountsService.GetAccrualsByUserID(userId)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	http.Redirect(w, r, fmt.Sprintf("/usermenu?id=%d", userId), http.StatusSeeOther)
	fmt.Println("10")
}
