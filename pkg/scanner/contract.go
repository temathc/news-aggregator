package scanner

import (
	"github.com/mmcdole/gofeed"

	"github.com/temathc/news-aggregator/models"
)

//go:generate mockgen -source=contract.go -destination contract_mocks_test.go -package $GOPACKAGE

type publicationsRepository interface {
	AddPublications([]models.Publications) error
}

type parser interface {
	ParseURL(feedURL string) (feed *gofeed.Feed, err error)
}
