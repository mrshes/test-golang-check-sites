package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"test-golang-check-sites/internal/models"
	"test-golang-check-sites/pkg"
	"time"
)

const (
	SCHEME = "http://"
)

type siteParserService struct {
	sync.Mutex
	domains   []string
	ListSites models.ListSites `json:"ListSites"`
}

func NewSiteParserService(ctx context.Context, domains []string, interval int) *siteParserService {
	service := &siteParserService{
		domains:   domains,
		ListSites: make(map[string]*models.SiteStatus),
	}
	go func() {
		service.Run(ctx, interval)
	}()
	return service
}

func (s *siteParserService) Run(ctx context.Context, interval int) {
	log.Println("start service ")
	for {
		select {
		case <-ctx.Done():
			log.Println("siteParserService Run closed")
			return
			break
		default:
			log.Println("launching the parser")
			s.checkSites(ctx)
			time.Sleep(time.Duration(interval) * time.Minute)
		}
	}
}

func (s *siteParserService) checkSites(ctx context.Context) {
	for _, domain := range s.domains {
		select {
		case <-ctx.Done():
			fmt.Println("checkSites closed loop!")
			return
		default:
			go func(domain string) {
				status := s.checkSiteStatus(SCHEME + domain)
				s.Lock()
				s.ListSites[domain] = status
				s.Unlock()
			}(domain)
		}
	}
	log.Println("END OF THE VERIFICATION CYCLE")
}

func (s *siteParserService) checkSiteStatus(domain string) (result *models.SiteStatus) {
	log.Printf("check status for %s", domain)
	var resp *http.Response
	var err error
	result = new(models.SiteStatus)
	timeVal := pkg.Stopwatch(func() {
		resp, err = s.makeRequest(domain)
	})
	if err != nil {
		log.Println("makeRequest error:", err)
		return result
	}
	if resp.StatusCode == http.StatusOK {
		result.Status = true
	}
	result.Time = time.Duration(timeVal)
	return result
}

func (s *siteParserService) makeRequest(domain string) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: func() func(req *http.Request, via []*http.Request) error {
			redirects := 0
			return func(req *http.Request, via []*http.Request) error {
				if redirects > 12 {
					return errors.New("after 12 redirects: ")
				}
				redirects++
				return nil
			}
		}(),
	}
	req, err := http.NewRequest("GET", domain, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
