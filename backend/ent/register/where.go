// Code generated by entc, DO NOT EDIT.

package register

import (
	"time"

	"github.com/TP/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RegisterPatientID applies equality check predicate on the "register_Patient_id" field. It's identical to RegisterPatientIDEQ.
func RegisterPatientID(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterName applies equality check predicate on the "register_name" field. It's identical to RegisterNameEQ.
func RegisterName(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterName), v))
	})
}

// RegisterCardID applies equality check predicate on the "register_cardID" field. It's identical to RegisterCardIDEQ.
func RegisterCardID(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterCardID), v))
	})
}

// RegisterAddress applies equality check predicate on the "register_address" field. It's identical to RegisterAddressEQ.
func RegisterAddress(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterAddress), v))
	})
}

// RegisterBirthday applies equality check predicate on the "register_birthday" field. It's identical to RegisterBirthdayEQ.
func RegisterBirthday(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterBirthday), v))
	})
}

// RegisterTel applies equality check predicate on the "register_tel" field. It's identical to RegisterTelEQ.
func RegisterTel(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterTel), v))
	})
}

// RegisterHight applies equality check predicate on the "register_hight" field. It's identical to RegisterHightEQ.
func RegisterHight(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterHight), v))
	})
}

// RegisterWeight applies equality check predicate on the "register_weight" field. It's identical to RegisterWeightEQ.
func RegisterWeight(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterWeight), v))
	})
}

// RegisterAge applies equality check predicate on the "register_age" field. It's identical to RegisterAgeEQ.
func RegisterAge(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterAge), v))
	})
}

// RegisterPatientIDEQ applies the EQ predicate on the "register_Patient_id" field.
func RegisterPatientIDEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDNEQ applies the NEQ predicate on the "register_Patient_id" field.
func RegisterPatientIDNEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDIn applies the In predicate on the "register_Patient_id" field.
func RegisterPatientIDIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterPatientID), v...))
	})
}

// RegisterPatientIDNotIn applies the NotIn predicate on the "register_Patient_id" field.
func RegisterPatientIDNotIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterPatientID), v...))
	})
}

// RegisterPatientIDGT applies the GT predicate on the "register_Patient_id" field.
func RegisterPatientIDGT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDGTE applies the GTE predicate on the "register_Patient_id" field.
func RegisterPatientIDGTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDLT applies the LT predicate on the "register_Patient_id" field.
func RegisterPatientIDLT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDLTE applies the LTE predicate on the "register_Patient_id" field.
func RegisterPatientIDLTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDContains applies the Contains predicate on the "register_Patient_id" field.
func RegisterPatientIDContains(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDHasPrefix applies the HasPrefix predicate on the "register_Patient_id" field.
func RegisterPatientIDHasPrefix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDHasSuffix applies the HasSuffix predicate on the "register_Patient_id" field.
func RegisterPatientIDHasSuffix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDEqualFold applies the EqualFold predicate on the "register_Patient_id" field.
func RegisterPatientIDEqualFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterPatientIDContainsFold applies the ContainsFold predicate on the "register_Patient_id" field.
func RegisterPatientIDContainsFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRegisterPatientID), v))
	})
}

// RegisterNameEQ applies the EQ predicate on the "register_name" field.
func RegisterNameEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterName), v))
	})
}

// RegisterNameNEQ applies the NEQ predicate on the "register_name" field.
func RegisterNameNEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterName), v))
	})
}

// RegisterNameIn applies the In predicate on the "register_name" field.
func RegisterNameIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterName), v...))
	})
}

// RegisterNameNotIn applies the NotIn predicate on the "register_name" field.
func RegisterNameNotIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterName), v...))
	})
}

// RegisterNameGT applies the GT predicate on the "register_name" field.
func RegisterNameGT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterName), v))
	})
}

// RegisterNameGTE applies the GTE predicate on the "register_name" field.
func RegisterNameGTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterName), v))
	})
}

// RegisterNameLT applies the LT predicate on the "register_name" field.
func RegisterNameLT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterName), v))
	})
}

// RegisterNameLTE applies the LTE predicate on the "register_name" field.
func RegisterNameLTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterName), v))
	})
}

// RegisterNameContains applies the Contains predicate on the "register_name" field.
func RegisterNameContains(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRegisterName), v))
	})
}

// RegisterNameHasPrefix applies the HasPrefix predicate on the "register_name" field.
func RegisterNameHasPrefix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRegisterName), v))
	})
}

// RegisterNameHasSuffix applies the HasSuffix predicate on the "register_name" field.
func RegisterNameHasSuffix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRegisterName), v))
	})
}

// RegisterNameEqualFold applies the EqualFold predicate on the "register_name" field.
func RegisterNameEqualFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRegisterName), v))
	})
}

// RegisterNameContainsFold applies the ContainsFold predicate on the "register_name" field.
func RegisterNameContainsFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRegisterName), v))
	})
}

// RegisterCardIDEQ applies the EQ predicate on the "register_cardID" field.
func RegisterCardIDEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDNEQ applies the NEQ predicate on the "register_cardID" field.
func RegisterCardIDNEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDIn applies the In predicate on the "register_cardID" field.
func RegisterCardIDIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterCardID), v...))
	})
}

// RegisterCardIDNotIn applies the NotIn predicate on the "register_cardID" field.
func RegisterCardIDNotIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterCardID), v...))
	})
}

// RegisterCardIDGT applies the GT predicate on the "register_cardID" field.
func RegisterCardIDGT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDGTE applies the GTE predicate on the "register_cardID" field.
func RegisterCardIDGTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDLT applies the LT predicate on the "register_cardID" field.
func RegisterCardIDLT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDLTE applies the LTE predicate on the "register_cardID" field.
func RegisterCardIDLTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDContains applies the Contains predicate on the "register_cardID" field.
func RegisterCardIDContains(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDHasPrefix applies the HasPrefix predicate on the "register_cardID" field.
func RegisterCardIDHasPrefix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDHasSuffix applies the HasSuffix predicate on the "register_cardID" field.
func RegisterCardIDHasSuffix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDEqualFold applies the EqualFold predicate on the "register_cardID" field.
func RegisterCardIDEqualFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRegisterCardID), v))
	})
}

// RegisterCardIDContainsFold applies the ContainsFold predicate on the "register_cardID" field.
func RegisterCardIDContainsFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRegisterCardID), v))
	})
}

// RegisterAddressEQ applies the EQ predicate on the "register_address" field.
func RegisterAddressEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressNEQ applies the NEQ predicate on the "register_address" field.
func RegisterAddressNEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressIn applies the In predicate on the "register_address" field.
func RegisterAddressIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterAddress), v...))
	})
}

// RegisterAddressNotIn applies the NotIn predicate on the "register_address" field.
func RegisterAddressNotIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterAddress), v...))
	})
}

// RegisterAddressGT applies the GT predicate on the "register_address" field.
func RegisterAddressGT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressGTE applies the GTE predicate on the "register_address" field.
func RegisterAddressGTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressLT applies the LT predicate on the "register_address" field.
func RegisterAddressLT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressLTE applies the LTE predicate on the "register_address" field.
func RegisterAddressLTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressContains applies the Contains predicate on the "register_address" field.
func RegisterAddressContains(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressHasPrefix applies the HasPrefix predicate on the "register_address" field.
func RegisterAddressHasPrefix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressHasSuffix applies the HasSuffix predicate on the "register_address" field.
func RegisterAddressHasSuffix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressEqualFold applies the EqualFold predicate on the "register_address" field.
func RegisterAddressEqualFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRegisterAddress), v))
	})
}

// RegisterAddressContainsFold applies the ContainsFold predicate on the "register_address" field.
func RegisterAddressContainsFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRegisterAddress), v))
	})
}

// RegisterBirthdayEQ applies the EQ predicate on the "register_birthday" field.
func RegisterBirthdayEQ(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterBirthday), v))
	})
}

// RegisterBirthdayNEQ applies the NEQ predicate on the "register_birthday" field.
func RegisterBirthdayNEQ(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterBirthday), v))
	})
}

// RegisterBirthdayIn applies the In predicate on the "register_birthday" field.
func RegisterBirthdayIn(vs ...time.Time) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterBirthday), v...))
	})
}

// RegisterBirthdayNotIn applies the NotIn predicate on the "register_birthday" field.
func RegisterBirthdayNotIn(vs ...time.Time) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterBirthday), v...))
	})
}

// RegisterBirthdayGT applies the GT predicate on the "register_birthday" field.
func RegisterBirthdayGT(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterBirthday), v))
	})
}

// RegisterBirthdayGTE applies the GTE predicate on the "register_birthday" field.
func RegisterBirthdayGTE(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterBirthday), v))
	})
}

// RegisterBirthdayLT applies the LT predicate on the "register_birthday" field.
func RegisterBirthdayLT(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterBirthday), v))
	})
}

// RegisterBirthdayLTE applies the LTE predicate on the "register_birthday" field.
func RegisterBirthdayLTE(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterBirthday), v))
	})
}

// RegisterTelEQ applies the EQ predicate on the "register_tel" field.
func RegisterTelEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterTel), v))
	})
}

// RegisterTelNEQ applies the NEQ predicate on the "register_tel" field.
func RegisterTelNEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterTel), v))
	})
}

// RegisterTelIn applies the In predicate on the "register_tel" field.
func RegisterTelIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterTel), v...))
	})
}

// RegisterTelNotIn applies the NotIn predicate on the "register_tel" field.
func RegisterTelNotIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterTel), v...))
	})
}

// RegisterTelGT applies the GT predicate on the "register_tel" field.
func RegisterTelGT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterTel), v))
	})
}

// RegisterTelGTE applies the GTE predicate on the "register_tel" field.
func RegisterTelGTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterTel), v))
	})
}

// RegisterTelLT applies the LT predicate on the "register_tel" field.
func RegisterTelLT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterTel), v))
	})
}

// RegisterTelLTE applies the LTE predicate on the "register_tel" field.
func RegisterTelLTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterTel), v))
	})
}

// RegisterHightEQ applies the EQ predicate on the "register_hight" field.
func RegisterHightEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterHight), v))
	})
}

// RegisterHightNEQ applies the NEQ predicate on the "register_hight" field.
func RegisterHightNEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterHight), v))
	})
}

// RegisterHightIn applies the In predicate on the "register_hight" field.
func RegisterHightIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterHight), v...))
	})
}

// RegisterHightNotIn applies the NotIn predicate on the "register_hight" field.
func RegisterHightNotIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterHight), v...))
	})
}

// RegisterHightGT applies the GT predicate on the "register_hight" field.
func RegisterHightGT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterHight), v))
	})
}

// RegisterHightGTE applies the GTE predicate on the "register_hight" field.
func RegisterHightGTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterHight), v))
	})
}

// RegisterHightLT applies the LT predicate on the "register_hight" field.
func RegisterHightLT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterHight), v))
	})
}

// RegisterHightLTE applies the LTE predicate on the "register_hight" field.
func RegisterHightLTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterHight), v))
	})
}

// RegisterWeightEQ applies the EQ predicate on the "register_weight" field.
func RegisterWeightEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterWeight), v))
	})
}

// RegisterWeightNEQ applies the NEQ predicate on the "register_weight" field.
func RegisterWeightNEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterWeight), v))
	})
}

// RegisterWeightIn applies the In predicate on the "register_weight" field.
func RegisterWeightIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterWeight), v...))
	})
}

// RegisterWeightNotIn applies the NotIn predicate on the "register_weight" field.
func RegisterWeightNotIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterWeight), v...))
	})
}

// RegisterWeightGT applies the GT predicate on the "register_weight" field.
func RegisterWeightGT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterWeight), v))
	})
}

// RegisterWeightGTE applies the GTE predicate on the "register_weight" field.
func RegisterWeightGTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterWeight), v))
	})
}

// RegisterWeightLT applies the LT predicate on the "register_weight" field.
func RegisterWeightLT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterWeight), v))
	})
}

// RegisterWeightLTE applies the LTE predicate on the "register_weight" field.
func RegisterWeightLTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterWeight), v))
	})
}

// RegisterAgeEQ applies the EQ predicate on the "register_age" field.
func RegisterAgeEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisterAge), v))
	})
}

// RegisterAgeNEQ applies the NEQ predicate on the "register_age" field.
func RegisterAgeNEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisterAge), v))
	})
}

// RegisterAgeIn applies the In predicate on the "register_age" field.
func RegisterAgeIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisterAge), v...))
	})
}

// RegisterAgeNotIn applies the NotIn predicate on the "register_age" field.
func RegisterAgeNotIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisterAge), v...))
	})
}

// RegisterAgeGT applies the GT predicate on the "register_age" field.
func RegisterAgeGT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisterAge), v))
	})
}

// RegisterAgeGTE applies the GTE predicate on the "register_age" field.
func RegisterAgeGTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisterAge), v))
	})
}

// RegisterAgeLT applies the LT predicate on the "register_age" field.
func RegisterAgeLT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisterAge), v))
	})
}

// RegisterAgeLTE applies the LTE predicate on the "register_age" field.
func RegisterAgeLTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisterAge), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Gender) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHave applies the HasEdge predicate on the "have" edge.
func HasHave() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HaveTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HaveTable, HaveColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHaveWith applies the HasEdge predicate on the "have" edge with a given conditions (other predicates).
func HasHaveWith(preds ...predicate.MedicalCare) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HaveInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HaveTable, HaveColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwner1 applies the HasEdge predicate on the "owner1" edge.
func HasOwner1() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Owner1Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, Owner1Table, Owner1Column),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwner1With applies the HasEdge predicate on the "owner1" edge with a given conditions (other predicates).
func HasOwner1With(preds ...predicate.Employee) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Owner1InverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, Owner1Table, Owner1Column),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHave1 applies the HasEdge predicate on the "have1" edge.
func HasHave1() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Have1Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, Have1Table, Have1Column),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHave1With applies the HasEdge predicate on the "have1" edge with a given conditions (other predicates).
func HasHave1With(preds ...predicate.Disease) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Have1InverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, Have1Table, Have1Column),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Register) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Register) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Register) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		p(s.Not())
	})
}
