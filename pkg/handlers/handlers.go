package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Publication struct {
	publicationRepo publicationsRepository
}

func NewPublication(publicationRepo publicationsRepository) Publication {
	return Publication{publicationRepo}
}

func (c *Publication) ListPublication(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	publications, err := c.publicationRepo.GetPublicationsWithLimit(limit)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	publicationsJson, err := json.Marshal(publications)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(publications) != 0 {
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(publicationsJson)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
