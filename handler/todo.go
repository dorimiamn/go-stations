package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

// ServeHTTP implements http.Handler interface.
func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		v := r.URL.Query()
		if v == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, _ := strconv.Atoi(v.Get("prev_id"))
		size, _ := strconv.Atoi(v.Get("size"))

		todos, err := h.svc.ReadTODO(r.Context(), int64(id), int64(size))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		convertedTodos := make([]model.TODO,0)

		for _, todo := range todos {
			convertedTodos = append(convertedTodos, *todo)
		}
		fmt.Println(convertedTodos)
		var res = model.ReadTODOResponse{TODOs: convertedTodos}
		json.NewEncoder(w).Encode(&res)

	case http.MethodPost:
		var req model.CreateTODORequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todo, _ := h.svc.CreateTODO(r.Context(), req.Subject, req.Description)
		var res = model.CreateTODOResponse{ TODO: *todo }
		json.NewEncoder(w).Encode(&res)
	case http.MethodPut:
		var req model.UpdateTODORequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todo, err := h.svc.UpdateTODO(r.Context(), req.ID, req.Subject, req.Description)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var res = model.UpdateTODOResponse{ TODO: *todo }
		json.NewEncoder(w).Encode(&res)
	}
}
