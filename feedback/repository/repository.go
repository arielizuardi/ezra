package repository

import (
	"github.com/arielizuardi/ezra/class"
	"github.com/arielizuardi/ezra/feedback"
)

// Repository ...
type Repository interface {
	FetchFacilitatorFeedbacks(facilitatorID int64, c *class.Class) ([]*feedback.FacilitatorFeedback, error)
	FetchPresenterFeedbacks(presenterID int64, c *class.Class, s *class.Session) ([]*feedback.PresenterFeedback, error)
	StorePresenterFeedbacks(feedbacks []*feedback.PresenterFeedback) error
	FetchAllFeedbackFields() ([]*feedback.Field, error)
}
