package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"poe.market/config"
	"poe.market/database"
)

func NewHandler(pgConfig config.PgConfig) *Handler {
	handler := &Handler{
		Mux:   chi.NewMux(),
		store: database.NewStore(pgConfig),
	}

	handler.Use(middleware.Logger)
	handler.Route("/soldItems", func(r chi.Router) {
		r.Get("/", handler.getSoldItems())
		r.Get("/{itemName}", handler.getItemPrices())
	})

	return handler
}

type Handler struct {
	*chi.Mux

	store *database.Store
}

func slugify(input string) string {
	return strings.ReplaceAll(input, " ", "-")
}

func deslugify(input string) string {
	return strings.ReplaceAll(input, "-", " ")
}

func (h *Handler) getSoldItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		itemNames, err := h.store.SoldItemStore.Items()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for i := range itemNames {
			itemNames[i] = slugify(itemNames[i])
		}

		data, err := json.Marshal(itemNames)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func (h *Handler) getItemPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		itemName := deslugify(chi.URLParam(r, "itemName"))
		itemPrices, err := h.store.SoldItemStore.Prices(itemName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(itemPrices)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}
