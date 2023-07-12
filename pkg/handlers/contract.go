package handlers

import (
	"github.com/temathc/news-aggregator/models"
)

//go:generate mockgen -source=contract.go -destination contract_mocks_test.go -package $GOPACKAGE

type publicationsRepository interface {
	GetPublicationsWithLimit(int) ([]models.Publications, error)
}
