package scanner

import (
	"strings"
	"sync"

	"github.com/microcosm-cc/bluemonday"

	models "github.com/temathc/news-aggregator/models"
)

type Scanner struct {
	publicationsRepo publicationsRepository
	parser           parser
}

func NewScanner(scannerRepo publicationsRepository, parser parser) Scanner {
	scanner := Scanner{
		publicationsRepo: scannerRepo,
		parser:           parser,
	}
	return scanner
}

func (c *Scanner) ScanRss(link string, wg *sync.WaitGroup) error {
	defer wg.Done()
	feed, err := c.parser.ParseURL(link)
	if err != nil {
		return err
	}
	p := bluemonday.StripTagsPolicy()
	publications := make([]models.Publications, 0, len(feed.Items))
	for _, item := range feed.Items {
		publ := models.Publications{}
		publ.GUID = &item.GUID
		publ.Title = item.Title
		publ.Description = strings.ReplaceAll(p.Sanitize(item.Description), "\n", "")
		publ.PubTime = *item.PublishedParsed
		publ.Link = item.Link
		publications = append(publications, publ)
	}

	return c.publicationsRepo.AddPublications(publications)
}
