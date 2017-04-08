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

// Field ...
// e.g Field{1,` Mampu menjelaskan tujuan dan manfaat kelas ini dengan baik`, 5}
// e.g Field{3,` Strength Statement`, `Suka sekali memasak`}
type Field struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description, omitempty"`
	Value       interface{} `json:"value, omitempty"`
}

// PresenterFeedback represents presenter feedback
type PresenterFeedback struct {
	Class       *class.Class
	Session     *class.Session
	Presenter   *presenter.Presenter
	Participant *participant.Participant
	Fields      []*Field
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// FacilitatorFeedback represents facilitator feedback
type FacilitatorFeedback struct {
	Class       *class.Class
	Facilitator *facilitator.Facilitator
	Participant *participant.Participant
	Fields      []*Field
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
