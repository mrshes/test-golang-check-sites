package server

import (
	"encoding/json"
	"net/http"
	"test-golang-check-sites/internal/models"
)

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	data := s.Services.SiteParserService.ListSites
	// Return all info
	if !r.URL.Query().Has("d") {
		s.newResponse(w, data)
		return
	}

	domain := r.URL.Query().Get("d")
	item, ok := data[domain]
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		s.newResponse(w, map[string]string{"error": "site not found"})
		return
	}
	s.Statistics.IncSingle()
	s.newResponse(w, models.ListSites{domain: item})
}

func (s *Server) getMin(writer http.ResponseWriter, request *http.Request) {
	data := s.Services.SiteParserService.ListSites
	lastIndex := ""
	result := make(models.ListSites, 1)
	for domain, item := range data {
		if len(result) == 0 {
			result[domain] = item
			lastIndex = domain
		} else {
			if result[lastIndex].Time > item.Time && item.Status == true {
				result[domain] = item
				delete(result, lastIndex)
				lastIndex = domain
			}
		}
	}
	s.Statistics.IncMin()
	s.newResponse(writer, result)
}

func (s *Server) getMax(writer http.ResponseWriter, request *http.Request) {
	data := s.Services.SiteParserService.ListSites
	lastIndex := ""
	result := make(models.ListSites, 1)
	for domain, item := range data {
		if len(result) == 0 {
			result[domain] = item
			lastIndex = domain
		} else {
			if result[lastIndex].Time < item.Time && item.Status == true {
				result[domain] = item
				delete(result, lastIndex)
				lastIndex = domain
			}
		}
	}
	s.Statistics.IncMax()
	s.newResponse(writer, result)
}

func (s *Server) getStatistics(writer http.ResponseWriter, request *http.Request) {
	stat := s.Statistics
	s.newResponse(writer, stat)
}

func (s *Server) newResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
