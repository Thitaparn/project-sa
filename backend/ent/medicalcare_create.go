// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/TP/app/ent/medicalcare"
	"github.com/TP/app/ent/patient"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// MedicalCareCreate is the builder for creating a MedicalCare entity.
type MedicalCareCreate struct {
	config
	mutation *MedicalCareMutation
	hooks    []Hook
}

// SetMedicalcareName sets the medicalcare_name field.
func (mcc *MedicalCareCreate) SetMedicalcareName(s string) *MedicalCareCreate {
	mcc.mutation.SetMedicalcareName(s)
	return mcc
}

// AddPatientIDs adds the patients edge to Patient by ids.
func (mcc *MedicalCareCreate) AddPatientIDs(ids ...int) *MedicalCareCreate {
	mcc.mutation.AddPatientIDs(ids...)
	return mcc
}

// AddPatients adds the patients edges to Patient.
func (mcc *MedicalCareCreate) AddPatients(p ...*Patient) *MedicalCareCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mcc.AddPatientIDs(ids...)
}

// Mutation returns the MedicalCareMutation object of the builder.
func (mcc *MedicalCareCreate) Mutation() *MedicalCareMutation {
	return mcc.mutation
}

// Save creates the MedicalCare in the database.
func (mcc *MedicalCareCreate) Save(ctx context.Context) (*MedicalCare, error) {
	if _, ok := mcc.mutation.MedicalcareName(); !ok {
		return nil, &ValidationError{Name: "medicalcare_name", err: errors.New("ent: missing required field \"medicalcare_name\"")}
	}
	if v, ok := mcc.mutation.MedicalcareName(); ok {
		if err := medicalcare.MedicalcareNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "medicalcare_name", err: fmt.Errorf("ent: validator failed for field \"medicalcare_name\": %w", err)}
		}
	}
	var (
		err  error
		node *MedicalCare
	)
	if len(mcc.hooks) == 0 {
		node, err = mcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalCareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mcc.mutation = mutation
			node, err = mcc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mcc.hooks) - 1; i >= 0; i-- {
			mut = mcc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mcc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mcc *MedicalCareCreate) SaveX(ctx context.Context) *MedicalCare {
	v, err := mcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mcc *MedicalCareCreate) sqlSave(ctx context.Context) (*MedicalCare, error) {
	mc, _spec := mcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mcc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	mc.ID = int(id)
	return mc, nil
}

func (mcc *MedicalCareCreate) createSpec() (*MedicalCare, *sqlgraph.CreateSpec) {
	var (
		mc    = &MedicalCare{config: mcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: medicalcare.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalcare.FieldID,
			},
		}
	)
	if value, ok := mcc.mutation.MedicalcareName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicalcare.FieldMedicalcareName,
		})
		mc.MedicalcareName = value
	}
	if nodes := mcc.mutation.PatientsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalcare.PatientsTable,
			Columns: []string{medicalcare.PatientsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return mc, _spec
}
