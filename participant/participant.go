package participant

type Participant struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	DOB         string `json:"dob"`
	PhoneNumber string `json:"phone_number"`
}
