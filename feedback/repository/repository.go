package repository

import "github.com/arielizuardi/ezra/feedback"

// Repository ...
type Repository interface {
	FetchFacilitatorFeedbacks(facilitatorID int64, batch int64, year int64) ([]*feedback.FacilitatorFeedback, error)
	FetchPresenterFeedbacks(presenterID int64, session int64, batch int64, year int64) ([]*feedback.PresenterFeedback, error)
	StorePresenterFeedbacks(feedbacks []*feedback.PresenterFeedback) error
	FetchAllFeedbackFields() ([]*feedback.Field, error)
}
