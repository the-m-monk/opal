package api

import (
	"encoding/json"
	"net/http"
	"opal/internal/jfstructs"
)

func EndpointUsersItemsUuidIntros(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !isHeaderAuthTokenValid(r) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	res := jfstructs.CommonItemList{}
	res.Items = make([]jfstructs.CommonItem, 0)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
