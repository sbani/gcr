package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"github.com/sbani/amycr/config"
	"github.com/sbani/amycr/http"
	"github.com/sbani/amycr/pkg"
	"github.com/sbani/amycr/storage"
)

// Handler holds all other handlers and prepares them for routing
type Handler struct {
	e           *echo.Echo
	ContentType *http.ContentTypeHandler
	Record      *http.RecordHandler
	Storage     *http.StorageHandler
}

// Start the handler and bootrap all others
func (h *Handler) Start(c *config.Config, e *echo.Echo) {
	storage, err := storage.NewManager(c)
	if err != nil {
		pkg.Must(errors.Wrap(err, "Storage"))
	}

	e.Use(middleware.Logger())

	h.ContentType = newContentTypeHandler(c, e, storage.ContentType())
	h.Record = newRecordHandler(c, e, storage)
	h.Storage = newStorageHandler(c, e, storage)
}
