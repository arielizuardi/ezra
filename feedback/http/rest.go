package http

import (
	"net/http"

	feedbackusecase "github.com/arielizuardi/ezra/feedback/usecase"
	"github.com/labstack/echo"
)

// ResponseError wraps json error response
type ResponseError struct {
	Message string `json:"error"`
}

type feedbackHTTPHandler struct {
	FeedbackUsecase feedbackusecase.FeedbackUsecase
}

// PresenterFeedbackRequest ...
type PresenterFeedbackRequest struct {
	ClassID     string                     `json:"class_id"`
	PresenterID int64                      `json:"presenter_id"`
	SessionID   int64                      `json:"session_id"`
	Mappings    []*feedbackusecase.Mapping `json:"mappings"`
	Values      [][]string                 `json:"values"`
}

func (f *feedbackHTTPHandler) HandleStorePresenterFeedbackFromGsheet(c echo.Context) error {
	pf := new(PresenterFeedbackRequest)
	if err := c.Bind(&pf); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &ResponseError{err.Error()})
	}

	if err := f.FeedbackUsecase.StorePresenterFeedbackWithMapping(pf.PresenterID, pf.ClassID, pf.SessionID, pf.Mappings, pf.Values); err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

// Init ...
func Init(e *echo.Echo, feedbackUsecase feedbackusecase.FeedbackUsecase) {
	h := &feedbackHTTPHandler{feedbackUsecase}
	e.POST(`/presenter/:id/feedbackmapping`, h.HandleStorePresenterFeedbackFromGsheet)
}
