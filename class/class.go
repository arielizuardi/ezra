package class

const (
	// COB const for Community of Believers
	COB = `COB`
	// COL const for Community of Leaders
	COL = `COL`
	// COC const for Community of Councellors
	COC = `COC`
)

// Class represents class name
type Class struct {
	Name  string
	Batch int64
	Year  int64
}
