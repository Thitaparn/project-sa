// Code generated by entc, DO NOT EDIT.

package medicalcare

const (
	// Label holds the string label denoting the medicalcare type in the database.
	Label = "medical_care"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMedicalcareName holds the string denoting the medicalcare_name field in the database.
	FieldMedicalcareName = "medicalcare_name"

	// EdgePatients holds the string denoting the patients edge name in mutations.
	EdgePatients = "patients"

	// Table holds the table name of the medicalcare in the database.
	Table = "medical_cares"
	// PatientsTable is the table the holds the patients relation/edge.
	PatientsTable = "patients"
	// PatientsInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientsInverseTable = "patients"
	// PatientsColumn is the table column denoting the patients relation/edge.
	PatientsColumn = "medicalcare_id"
)

// Columns holds all SQL columns for medicalcare fields.
var Columns = []string{
	FieldID,
	FieldMedicalcareName,
}

var (
	// MedicalcareNameValidator is a validator for the "medicalcare_name" field. It is called by the builders before save.
	MedicalcareNameValidator func(string) error
)