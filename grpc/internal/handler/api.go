package handler

import (
	"encoding/json"
	"net/http"

	pb "grpc/internal/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type API struct {
	client pb.APIClient
}

func New(client pb.APIClient) *API {
	return &API{
		client: client,
	}
}

func (a *API) All(w http.ResponseWriter, r *http.Request) error {
	items, err := a.client.All(r.Context(), &emptypb.Empty{})
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
	items, err := a.client.Jobs(r.Context(), &emptypb.Empty{})
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
	items, err := a.client.Stories(r.Context(), &emptypb.Empty{})
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
