package http_test

import (
	nethttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arielizuardi/ezra/feedback"
	"github.com/arielizuardi/ezra/feedback/http"
	usecase "github.com/arielizuardi/ezra/feedback/usecase/mocks"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	presenterFeedbackJSON = `{
        "class_id": "COL-B1-2016",
        "presenter_id":1,
        "session_id":1,
        "mappings":[
            {"header_id":0,"field_id":1},
            {"header_id":1,"field_id":2}
        ],
        "values":[
            ["123","456"],
            ["123","456"]
        ]
	}`
)

func TestStorePresenterFeedbackFromGsheet(t *testing.T) {
	u := new(usecase.FeedbackUsecase)
	u.On(`StorePresenterFeedbackWithMapping`,
		int64(1),
		`COL-B1-2016`,
		int64(1),
		mock.AnythingOfType(`[]*usecase.Mapping`),
		mock.AnythingOfType(`[][]string`),
	).Return([]*feedback.PresenterFeedback{}, nil)

	// Setup
	e := echo.New()

	req, err := nethttp.NewRequest(echo.POST, "/presenter/1/feedbackmapping", strings.NewReader(presenterFeedbackJSON))
	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := new(http.FeedbackHTTPHandler)

		h.FeedbackUsecase = u

		if assert.NoError(t, h.HandleStorePresenterFeedbackFromGsheet(c)) {
			assert.Equal(t, nethttp.StatusCreated, rec.Code)
		}
	}
}

func TestStorePresenterFeedbackFromGsheetWithInvalidParams(t *testing.T) {
	u := new(usecase.FeedbackUsecase)
	u.On(`StorePresenterFeedbackWithMapping`,
		int64(1),
		`COL-B1-2016`,
		int64(1),
		mock.AnythingOfType(`[]*usecase.Mapping`),
		mock.AnythingOfType(`[][]string`),
	).Return([]*feedback.PresenterFeedback{}, nil)

	// Setup
	e := echo.New()

	req, err := nethttp.NewRequest(echo.POST, "/presenter/1/feedbackmapping", strings.NewReader(`{}`))
	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := new(http.FeedbackHTTPHandler)

		h.FeedbackUsecase = u

		if assert.NoError(t, h.HandleStorePresenterFeedbackFromGsheet(c)) {
			assert.Equal(t, nethttp.StatusBadRequest, rec.Code)
		}
	}
}
