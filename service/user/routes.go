package user

import (
	"encoding/json"
	"net/http"

	"github.com/Rothiii/golang-marketplace/types"
	"github.com/Rothiii/golang-marketplace/utils"
	"github.com/gorilla/mux"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.LoginHandler).Methods("POST")
	router.HandleFunc("/register", h.LoginHandler).Methods("POST")
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
}
func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
}
