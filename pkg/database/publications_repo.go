package database

import (
	models "github.com/temathc/news-aggregator/models"
)

func (r RepoDB) GetPublicationsWithLimit(limit int) ([]models.Publications, error) {
	publList := make([]models.Publications, 0, limit)

	query := `select title, description, pubtime, link
		  from publications order by pubtime DESC limit $1`

	if err := r.db.Select(&publList, query, limit); err != nil {
		return nil, err
	}
	return publList, nil
}

func (r RepoDB) AddPublications(pubList []models.Publications) error {
	query, args, err := r.db.BindNamed(
		`insert into publications (
						  guid,
						  title,
						  description,
						  pubtime,
						  link)
		values (
				  :guid,
				  :title,
				  :description,
				  :pubtime,
				  :link)
		on conflict (guid) do nothing`, pubList)
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
