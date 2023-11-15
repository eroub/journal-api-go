package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eroub/journal-api-go/internal/app"
	"github.com/eroub/journal-api-go/internal/app/model"
	"github.com/gorilla/mux"
)

func GetEquityAmountsForUser(w http.ResponseWriter, r *http.Request) {
    // TODO: Retrieve userID from authenticated user context
    userID := 1 // Placeholder

    var accounts []model.Account
    result := app.DB.Where("user_id = ?", userID).Find(&accounts)
    if result.Error != nil {
        http.Error(w, "Error retrieving equity amounts for the user's accounts", http.StatusInternalServerError)
        return
    }

    // Respond with the accounts
    json.NewEncoder(w).Encode(accounts)
}

func UpdateAccountEquity(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    accountID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid account ID", http.StatusBadRequest)
        return
    }

    var account model.Account
    if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    result := app.DB.Model(&model.Account{}).Where("account_id = ?", accountID).Update("equity", account.Equity)
    if result.Error != nil {
        http.Error(w, "Error updating account equity", http.StatusInternalServerError)
        return
    }

    if result.RowsAffected == 0 {
        http.Error(w, "Account not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Account equity was updated successfully."))
}
