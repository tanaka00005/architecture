package handler

import (
	"architecture/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type DiaryHandler struct {
	svc service.DiaryService
}

func NewDiaryHandler(svc service.DiaryService) *DiaryHandler {
	return &DiaryHandler{svc: svc}
}

func (h *DiaryHandler) GetUserWithPosts(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	userID,err := strconv.ParseUint(idStr,10,64)

	if err != nil {
		http.Error(w,"invalid user ID",http.StatusBadRequest)
		return 
	}

	userWithPosts, err := h.svc.GetUserWithPosts(uint(userID))
	if err != nil {
		http.Error(w,"internal server error",http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type","application/json")
	if err := json.NewEncoder(w).Encode(userWithPosts); err != nil {
		http.Error(w, "Faild to encode response",http.StatusInternalServerError)
	}

}
