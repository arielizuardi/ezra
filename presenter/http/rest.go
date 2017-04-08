package http

import (
	"net/http"

	presenterusecase "github.com/arielizuardi/ezra/presenter/usecase"
	"github.com/labstack/echo"
)

// ResponseError wraps json error response
type ResponseError struct {
	Message string `json:"error"`
}

type PresenterHTTPHandler struct {
	PresenterUsecase presenterusecase.PresenterUsecase
}

func (h *PresenterHTTPHandler) HandleFetchAllPresenters(c echo.Context) error {
	presenters, err := h.PresenterUsecase.FetchAllPresenters()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	return c.JSON(http.StatusOK, presenters)
}

func Init(e *echo.Echo, p presenterusecase.PresenterUsecase) {
	h := &PresenterHTTPHandler{p}
	e.GET(`/presenter`, h.HandleFetchAllPresenters)
}
