package http

import (
	"net/http"
	"strconv"

	reportusecase "github.com/arielizuardi/ezra/feedback/report/usecase"
	"github.com/labstack/echo"
)

// ResponseError wraps json error response
type ResponseError struct {
	Message string `json:"error"`
}

type ReportHTTPHandler struct {
	ReportUsecase reportusecase.ReportUsecase
}

func (h *ReportHTTPHandler) HandleGeneratePresenterReport(c echo.Context) error {
	strPresenterID := c.Param(`presenter_id`)
	presenterID, err := strconv.Atoi(strPresenterID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{`'presenter_id' field must be a number`})
	}

	classID := c.QueryParam(`class_id`)
	if classID == `` {
		return c.JSON(http.StatusBadRequest, &ResponseError{`'class_id' field required`})
	}

	strSessionID := c.QueryParam(`session_id`)
	sessionID, err := strconv.Atoi(strSessionID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{`'session_id' field must be a number`})
	}

	report, err := h.ReportUsecase.GeneratePresenterReport(int64(presenterID), classID, int64(sessionID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	return c.JSON(http.StatusOK, report)
}

func Init(e *echo.Echo, reportUsecase reportusecase.ReportUsecase) {
	h := &ReportHTTPHandler{reportUsecase}
	e.GET(`/presenter/:presenter_id/report`, h.HandleGeneratePresenterReport)

}
