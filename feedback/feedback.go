package feedback

import (
	"time"

	"github.com/arielizuardi/ezra/class"
	"github.com/arielizuardi/ezra/facilitator"
	"github.com/arielizuardi/ezra/participant"
	"github.com/arielizuardi/ezra/presenter"
)

var PresenterRatingKey = []string{
	`Penguasaan materi`,
	`Sistematika penyajian`,
	`Gaya atau metode penyajian`,
	`Pengaturan waktu`,
	`Penggunaan alat bantu`,
	`Nilai keseluruhan`,
}

var FacilitatorRatingKey = []string{
	`Mampu menjelaskan tujuan dan manfaat kelas ini dengan baik`,
	`Membangun hubungan baik dengan saya`,
	`Mampu mengajak peserta untuk berdiskusi`,
	`Mampu membuat proses diskusi berjalan dengan baik`,
	`Mampu menjawab pertanyaan / concern yang ada selama diskusi kelompok & memberikan feedback yang bermanfaat`,
	`Memiliki kedalaman materi yang dibutuhkan`,
	`Bersikap profesional, berbusana rapi serta berperilaku & bertutur kata sopan`,
}

// Rating represents rating name and rating score
// e.g Rating{1,` Mampu menjelaskan tujuan dan manfaat kelas ini dengan baik`, 5}
type Rating struct {
	Key         string
	Description string
	Score       int64
}

// NewRating returns new instance of rating
func NewRating(key string, description string, score int64) *Rating {
	return &Rating{key, description, score}
}

// PresenterFeedback represents presenter feedback
type PresenterFeedback struct {
	Class              *class.Class
	Session            int64
	Presenter          *presenter.Presenter
	Participant        *participant.Participant
	Ratings            []*Rating
	PositiveComment    string
	ImprovementComment string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// FacilitatorFeedback represents facilitator feedback
type FacilitatorFeedback struct {
	Class       *class.Class
	Facilitator *facilitator.Facilitator
	Participant *participant.Participant
	Ratings     []*Rating
	Comment     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
