// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/TP/app/ent/disease"
	"github.com/TP/app/ent/employee"
	"github.com/TP/app/ent/gender"
	"github.com/TP/app/ent/medicalcare"
	"github.com/TP/app/ent/patient"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PatientCreate is the builder for creating a Patient entity.
type PatientCreate struct {
	config
	mutation *PatientMutation
	hooks    []Hook
}

// SetPatientID sets the patient_ID field.
func (pc *PatientCreate) SetPatientID(s string) *PatientCreate {
	pc.mutation.SetPatientID(s)
	return pc
}

// SetPatientName sets the patient_name field.
func (pc *PatientCreate) SetPatientName(s string) *PatientCreate {
	pc.mutation.SetPatientName(s)
	return pc
}

// SetPatientCardID sets the patient_cardID field.
func (pc *PatientCreate) SetPatientCardID(s string) *PatientCreate {
	pc.mutation.SetPatientCardID(s)
	return pc
}

// SetPatientAddress sets the patient_address field.
func (pc *PatientCreate) SetPatientAddress(s string) *PatientCreate {
	pc.mutation.SetPatientAddress(s)
	return pc
}

// SetPatientBirthday sets the patient_birthday field.
func (pc *PatientCreate) SetPatientBirthday(t time.Time) *PatientCreate {
	pc.mutation.SetPatientBirthday(t)
	return pc
}

// SetPatientTel sets the patient_tel field.
func (pc *PatientCreate) SetPatientTel(s string) *PatientCreate {
	pc.mutation.SetPatientTel(s)
	return pc
}

// SetPatientAge sets the patient_age field.
func (pc *PatientCreate) SetPatientAge(i int) *PatientCreate {
	pc.mutation.SetPatientAge(i)
	return pc
}

// SetGenderID sets the gender edge to Gender by id.
func (pc *PatientCreate) SetGenderID(id int) *PatientCreate {
	pc.mutation.SetGenderID(id)
	return pc
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (pc *PatientCreate) SetNillableGenderID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetGenderID(*id)
	}
	return pc
}

// SetGender sets the gender edge to Gender.
func (pc *PatientCreate) SetGender(g *Gender) *PatientCreate {
	return pc.SetGenderID(g.ID)
}

// SetMedicalcareID sets the medicalcare edge to MedicalCare by id.
func (pc *PatientCreate) SetMedicalcareID(id int) *PatientCreate {
	pc.mutation.SetMedicalcareID(id)
	return pc
}

// SetNillableMedicalcareID sets the medicalcare edge to MedicalCare by id if the given value is not nil.
func (pc *PatientCreate) SetNillableMedicalcareID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetMedicalcareID(*id)
	}
	return pc
}

// SetMedicalcare sets the medicalcare edge to MedicalCare.
func (pc *PatientCreate) SetMedicalcare(m *MedicalCare) *PatientCreate {
	return pc.SetMedicalcareID(m.ID)
}

// SetEmployeeID sets the employee edge to Employee by id.
func (pc *PatientCreate) SetEmployeeID(id int) *PatientCreate {
	pc.mutation.SetEmployeeID(id)
	return pc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (pc *PatientCreate) SetNillableEmployeeID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetEmployeeID(*id)
	}
	return pc
}

// SetEmployee sets the employee edge to Employee.
func (pc *PatientCreate) SetEmployee(e *Employee) *PatientCreate {
	return pc.SetEmployeeID(e.ID)
}

// SetDiseaseID sets the disease edge to Disease by id.
func (pc *PatientCreate) SetDiseaseID(id int) *PatientCreate {
	pc.mutation.SetDiseaseID(id)
	return pc
}

// SetNillableDiseaseID sets the disease edge to Disease by id if the given value is not nil.
func (pc *PatientCreate) SetNillableDiseaseID(id *int) *PatientCreate {
	if id != nil {
		pc = pc.SetDiseaseID(*id)
	}
	return pc
}

// SetDisease sets the disease edge to Disease.
func (pc *PatientCreate) SetDisease(d *Disease) *PatientCreate {
	return pc.SetDiseaseID(d.ID)
}

// Mutation returns the PatientMutation object of the builder.
func (pc *PatientCreate) Mutation() *PatientMutation {
	return pc.mutation
}

// Save creates the Patient in the database.
func (pc *PatientCreate) Save(ctx context.Context) (*Patient, error) {
	if _, ok := pc.mutation.PatientID(); !ok {
		return nil, &ValidationError{Name: "patient_ID", err: errors.New("ent: missing required field \"patient_ID\"")}
	}
	if v, ok := pc.mutation.PatientID(); ok {
		if err := patient.PatientIDValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_ID", err: fmt.Errorf("ent: validator failed for field \"patient_ID\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PatientName(); !ok {
		return nil, &ValidationError{Name: "patient_name", err: errors.New("ent: missing required field \"patient_name\"")}
	}
	if v, ok := pc.mutation.PatientName(); ok {
		if err := patient.PatientNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_name", err: fmt.Errorf("ent: validator failed for field \"patient_name\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PatientCardID(); !ok {
		return nil, &ValidationError{Name: "patient_cardID", err: errors.New("ent: missing required field \"patient_cardID\"")}
	}
	if v, ok := pc.mutation.PatientCardID(); ok {
		if err := patient.PatientCardIDValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_cardID", err: fmt.Errorf("ent: validator failed for field \"patient_cardID\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PatientAddress(); !ok {
		return nil, &ValidationError{Name: "patient_address", err: errors.New("ent: missing required field \"patient_address\"")}
	}
	if v, ok := pc.mutation.PatientAddress(); ok {
		if err := patient.PatientAddressValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_address", err: fmt.Errorf("ent: validator failed for field \"patient_address\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PatientBirthday(); !ok {
		return nil, &ValidationError{Name: "patient_birthday", err: errors.New("ent: missing required field \"patient_birthday\"")}
	}
	if _, ok := pc.mutation.PatientTel(); !ok {
		return nil, &ValidationError{Name: "patient_tel", err: errors.New("ent: missing required field \"patient_tel\"")}
	}
	if v, ok := pc.mutation.PatientTel(); ok {
		if err := patient.PatientTelValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_tel", err: fmt.Errorf("ent: validator failed for field \"patient_tel\": %w", err)}
		}
	}
	if _, ok := pc.mutation.PatientAge(); !ok {
		return nil, &ValidationError{Name: "patient_age", err: errors.New("ent: missing required field \"patient_age\"")}
	}
	if v, ok := pc.mutation.PatientAge(); ok {
		if err := patient.PatientAgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "patient_age", err: fmt.Errorf("ent: validator failed for field \"patient_age\": %w", err)}
		}
	}
	var (
		err  error
		node *Patient
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PatientCreate) SaveX(ctx context.Context) *Patient {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientCreate) sqlSave(ctx context.Context) (*Patient, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PatientCreate) createSpec() (*Patient, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patient{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patient.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PatientID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientID,
		})
		pa.PatientID = value
	}
	if value, ok := pc.mutation.PatientName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientName,
		})
		pa.PatientName = value
	}
	if value, ok := pc.mutation.PatientCardID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientCardID,
		})
		pa.PatientCardID = value
	}
	if value, ok := pc.mutation.PatientAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientAddress,
		})
		pa.PatientAddress = value
	}
	if value, ok := pc.mutation.PatientBirthday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: patient.FieldPatientBirthday,
		})
		pa.PatientBirthday = value
	}
	if value, ok := pc.mutation.PatientTel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldPatientTel,
		})
		pa.PatientTel = value
	}
	if value, ok := pc.mutation.PatientAge(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: patient.FieldPatientAge,
		})
		pa.PatientAge = value
	}
	if nodes := pc.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MedicalcareIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.MedicalcareTable,
			Columns: []string{patient.MedicalcareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalcare.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.EmployeeTable,
			Columns: []string{patient.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.DiseaseTable,
			Columns: []string{patient.DiseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}
