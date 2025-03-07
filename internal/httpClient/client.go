package httpClient

import (
	"context"
	"encoding/json"
	"linkSh/pkg/httpkit"
	"log"
	"net/http"
)

type linker interface {
	ShortLink(ctx context.Context, data string) (string, error)
	LongLink(ctx context.Context, data string) (string, error)
	GiveInfoAboutLink(ctx context.Context) ([]string, error)
}

type Client struct {
	linkShortenerSvc linker
}

func New(svc linker) *Client {
	return &Client{
		linkShortenerSvc: svc,
	}
}

func (s *Client) ShortenerLink(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Url string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	link, err := s.linkShortenerSvc.ShortLink(r.Context(), requestData.Url)
	if err != nil {
		httpkit.HTTPResponse(w, "no success", http.StatusInternalServerError)
		return
	}

	httpkit.HTTPResponse(w, link, http.StatusOK)
}

func (s *Client) OriginalLink(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Url string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	link, err := s.linkShortenerSvc.LongLink(r.Context(), requestData.Url)
	if err != nil {
		httpkit.HTTPResponse(w, "no success", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	httpkit.HTTPResponse(w, link, http.StatusOK)

}
func (s *Client) GiveAboutLink(w http.ResponseWriter, r *http.Request) {
	links, err := s.linkShortenerSvc.GiveInfoAboutLink(r.Context())
	if err != nil {
		return
	}

	if links == nil {
		httpResponse(w, "not links", http.StatusNotImplemented)
	}

	httpResponse(w, links, http.StatusOK)
}

func httpResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
