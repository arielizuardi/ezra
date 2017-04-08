package http

import (
	"net/http"

	feedbackusecase "github.com/arielizuardi/ezra/feedback/usecase"
	"github.com/arielizuardi/ezra/validator"
	"github.com/labstack/echo"
)

// ResponseError wraps json error response
type ResponseError struct {
	Message string `json:"error"`
}

type FeedbackHTTPHandler struct {
	FeedbackUsecase feedbackusecase.FeedbackUsecase
}

// PresenterFeedbackRequest ...
type PresenterFeedbackRequest struct {
	ClassID     string                     `json:"class_id" validate:"required"`
	PresenterID int64                      `json:"presenter_id" validate:"required"`
	SessionID   int64                      `json:"session_id" validate:"required"`
	Mappings    []*feedbackusecase.Mapping `json:"mappings" validate:"required"`
	Values      [][]string                 `json:"values" validate:"required"`
}

func (f *FeedbackHTTPHandler) HandleStorePresenterFeedbackFromGsheet(c echo.Context) error {
	// TODO rest_test and change the way to generate report
	pf := new(PresenterFeedbackRequest)
	if err := c.Bind(&pf); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ResponseError{err.Error()})
	}

	vld := validator.NewRequestValidator()
	if err := vld.Validate(pf); err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{err.Error()})
	}

	if _, err := f.FeedbackUsecase.StorePresenterFeedbackWithMapping(pf.PresenterID, pf.ClassID, pf.SessionID, pf.Mappings, pf.Values); err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func (f *FeedbackHTTPHandler) HandleFetchAllFeedbackFields(c echo.Context) error {
	feedbackFields, err := f.FeedbackUsecase.FetchAllFeedbackFields()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	return c.JSON(http.StatusOK, feedbackFields)
}

// Init ...
func Init(e *echo.Echo, feedbackUsecase feedbackusecase.FeedbackUsecase) {
	h := &FeedbackHTTPHandler{feedbackUsecase}
	e.GET(`/feedback/field`, h.HandleFetchAllFeedbackFields)
	e.POST(`/exportfeedback`, h.HandleStorePresenterFeedbackFromGsheet)
}
