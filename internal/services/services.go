package services

import (
	"context"
	"test-golang-check-sites/internal/storage"
)

type Services struct {
	SiteParserService *siteParserService
}

type Deps struct {
	Ctx      context.Context
	Storage  *storage.Storage
	Interval int
}

func NewServices(deps Deps) *Services {
	return &Services{
		SiteParserService: NewSiteParserService(deps.Ctx, deps.Storage.Domains, deps.Interval),
	}
}
