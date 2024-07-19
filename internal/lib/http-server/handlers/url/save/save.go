package save

import (
	"errors"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"log/slog"
	"net/http"
	"urlshortner/internal/lib/api/response"
	"urlshortner/internal/lib/logger/sl"
	"urlshortner/internal/lib/random"
	"urlshortner/internal/storage"
)

const aliasLength = 8

type URLSaver interface {
	SaveURL(urlToSave, alias string) (int64, error)
}

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}
type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		const operation = "handlers.url.save.New"

		log = log.With(
			slog.String("op", operation),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		var request Request

		err := render.DecodeJSON(r.Body, &request)

		if err != nil {
			log.Error("failed to decode request", sl.Err(err))

			render.JSON(w, r, response.ERROR("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", request))

		if err := validator.New().Struct(request); err != nil {

			validateError := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, response.ValidationError(validateError))

			return
		}
		alias := request.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLength)
		}
		id, err := urlSaver.SaveURL(request.URL, alias)
		if errors.Is(err, storage.ErrURLExists) {
			log.Info("url already exists", slog.String("url", request.URL))

			render.JSON(w, r, response.ERROR("url already exists"))

			return
		}
		if err != nil {
			log.Error("failed to save url", sl.Err(err))
			render.JSON(w, r, response.ERROR("failed to save url"))
			return
		}
		log.Info("url added", slog.Int64("id", id))

		render.JSON(w, r, Response{
			Response: response.OK(),
			Alias:    alias,
		})
	}
}
