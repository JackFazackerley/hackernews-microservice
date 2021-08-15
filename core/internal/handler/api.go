package handler

import (
	"encoding/json"
	"net/http"

	"core/internal/store"
)

type API struct {
	reader store.Reader
}

func New(reader store.Reader) *API {
	return &API{
		reader: reader,
	}
}

func (a *API) All(w http.ResponseWriter, r *http.Request) error {
	items, err := a.reader.All()
	if err != nil {
		return HTTPError{
			Message: err,
			Code:    http.StatusInternalServerError,
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		return HTTPError{
			Message: err,
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (a *API) Jobs(w http.ResponseWriter, r *http.Request) error {
	items, err := a.reader.Jobs()
	if err != nil {
		return HTTPError{
			Message: err,
			Code:    http.StatusInternalServerError,
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		return HTTPError{
			Message: err,
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (a *API) Stories(w http.ResponseWriter, r *http.Request) error {
	items, err := a.reader.Stories()
	if err != nil {
		return HTTPError{
			Message: err,
			Code:    http.StatusInternalServerError,
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		return HTTPError{
			Message: err,
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
