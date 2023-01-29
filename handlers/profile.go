package handlers

// import (
// 	"encoding/json"
// 	"housy/models"
// 	"housy/repositories"
// 	"net/http"
// 	"strconv"

// 	"github.com/go-playground/validator/v10/translations/id"
// 	"github.com/gorilla/mux"
// 	"golang.org/x/text/message"
// 	"honnef.co/go/tools/analysis/code"
// )

// type handlerProfile struct {
// 	ProfileRepository repositories.ProfileRepository
// }

// func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
// 	return &handlerProfile{ProfileRepository}
// }

// func (h *handlerProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	id, _ = strconv.Atoi(mux.Vars(r)["id"])

// 	profile := models.Profile
// 	profile, err := h.ProfileRepository.GetProfile(id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{code: http.StatusBadRequest, message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}
// }