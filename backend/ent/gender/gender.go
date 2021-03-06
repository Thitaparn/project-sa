// Code generated by entc, DO NOT EDIT.

package gender

const (
	// Label holds the string label denoting the gender type in the database.
	Label = "gender"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGenderName holds the string denoting the gender_name field in the database.
	FieldGenderName = "gender_name"

	// EdgePatients holds the string denoting the patients edge name in mutations.
	EdgePatients = "patients"

	// Table holds the table name of the gender in the database.
	Table = "genders"
	// PatientsTable is the table the holds the patients relation/edge.
	PatientsTable = "patients"
	// PatientsInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientsInverseTable = "patients"
	// PatientsColumn is the table column denoting the patients relation/edge.
	PatientsColumn = "gender_id"
)

// Columns holds all SQL columns for gender fields.
var Columns = []string{
	FieldID,
	FieldGenderName,
}

var (
	// GenderNameValidator is a validator for the "gender_name" field. It is called by the builders before save.
	GenderNameValidator func(string) error
)
