package redirect

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"

	resp "urlshortner/internal/lib/api/response"
	"urlshortner/internal/lib/logger/sl"
	"urlshortner/internal/storage"
)

// URLGetter is an interface for getting url by alias.
//
//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=URLGetter
type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const operation = "handlers.url.redirect.New"

		log := log.With(
			slog.String("operation", operation),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, resp.ERROR("invalid request"))

			return
		}

		resURL, err := urlGetter.GetURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)

			render.JSON(w, r, resp.ERROR("not found"))

			return
		}
		if err != nil {
			log.Error("failed to get url", sl.Err(err))

			render.JSON(w, r, resp.ERROR("internal error"))

			return
		}

		log.Info("got url", slog.String("url", resURL))

		// redirect to found url
		http.Redirect(w, r, resURL, http.StatusFound)
	}
}
