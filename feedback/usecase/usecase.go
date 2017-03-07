package usecase

import "github.com/arielizuardi/ezra/feedback"

type FeedbackUsecase interface {
	FetchRatings() ([]*feedback.Rating, error)
	StoreBatchFacilitatorFeedback(fb *feedback.BatchFeedbackFacilitator) error
	StoreSessionPresenterFeedback(fb *feedback.SessionFeedbackPresenter) error
}
