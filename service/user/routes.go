package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/Types"
	"github.com/sikozonpc/ecom/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get JSON payload
	var payload Types.RegisterUserPayload
	if err := utils.ParseJSON(r.Body, &payload); err != nil {
		utils.WriteError(w.(http.ResponseWriter), http.StatusBadRequest, err)
	}

	// Check if the user exists

	// If exists, return error
	// If not, create user
}
