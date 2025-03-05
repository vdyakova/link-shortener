package handler

import (
	"encoding/json"
	"fmt"
	"linkSh/httpkit"
	"linkSh/internal/service"
	"net/http"
)

type ShortenerHandler struct {
	service service.ShortenerService
}

func NewShortenerHandler(svc service.ShortenerService) *ShortenerHandler {
	return &ShortenerHandler{
		service: svc,
	}
}

func (s *ShortenerHandler) ShortenerLink(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Url string `json:"url"`
	}
	ctx := r.Context()
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	link, err := s.service.ShortLink(ctx, requestData.Url)
	if err != nil {
		httpkit.HTTPResponse(w, "no success", http.StatusInternalServerError)
		return
	}
	httpkit.HTTPResponse(w, link, http.StatusOK)

}

func (s *ShortenerHandler) OriginalLink(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Url string `json:"url"`
	}
	ctx := r.Context()
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	link, err := s.service.LongLink(ctx, requestData.Url)
	if err != nil {
		httpkit.HTTPResponse(w, "no success", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	httpkit.HTTPResponse(w, link, http.StatusOK)

}
func (s *ShortenerHandler) GiveAboutLink(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	links, err := s.service.GiveInfoAboutLink(ctx)
	if links == nil {
		httpkit.HTTPResponse(w, "not links", http.StatusNotImplemented)
	}
	if err != nil {
		return
	}
	httpkit.HTTPResponse(w, links, http.StatusOK)
}
