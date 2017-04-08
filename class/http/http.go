package http

import (
	"net/http"

	classusecase "github.com/arielizuardi/ezra/class/usecase"
	"github.com/labstack/echo"
)

// ResponseError wraps json error response
type ResponseError struct {
	Message string `json:"error"`
}

type ClassHTTPHandler struct {
	ClassUsecase classusecase.ClassUsecase
}

func (h *ClassHTTPHandler) HandleFetchAllClasses(c echo.Context) error {
	classes, err := h.ClassUsecase.FetchAllClasses()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	return c.JSON(http.StatusOK, classes)
}

func (h *ClassHTTPHandler) HandleFetchAllSessions(c echo.Context) error {
	sessions, err := h.ClassUsecase.FetchAllSessions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	return c.JSON(http.StatusOK, sessions)
}

func Init(e *echo.Echo, c classusecase.ClassUsecase) {
	h := &ClassHTTPHandler{c}
	e.GET(`/class`, h.HandleFetchAllClasses)
	e.GET(`/session`, h.HandleFetchAllSessions)
}
