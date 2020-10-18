package controllers

import (
	"context"
	"time"

	"github.com/TP/app/ent/disease"
	"github.com/TP/app/ent/employee"
	"github.com/TP/app/ent/gender"
	"github.com/TP/app/ent/medicalcare"
	"github.com/TP/app/ent/patient"

	//  "fmt"
	"strconv"

	"github.com/TP/app/ent"
	"github.com/gin-gonic/gin"
)

// PatientController defines the struct for the Patient controller
type PatientController struct {
	client *ent.Client
	router gin.IRouter
}

type Patient struct {
	PatientID       string
	PatientName     string
	PatientCardID   int
	Gender          int
	PatientAddress  string
	PatientTel      int
	PatientBirthday string
	PatientAge      int
	Disease         int
	MedicalCare     int
	Employee        int
}

// PatientCreate handles POST requests for adding Patient entities
// @Summary Create Patient
// @Description Create Patient
// @ID create-Patient
// @Accept   json
// @Produce  json
// @Param Patient body ent.Patient true "Patient entity"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [post]
func (ctl *PatientController) PatientCreate(c *gin.Context) {
	obj := Patient{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Patient binding failed",
		})
		return
	}

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Gender not found",
		})
		return
	}

	
	d, err := ctl.client.Disease.
		Query().
		Where(disease.IDEQ(int(obj.Disease))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Disease not found",
		})
		return
	}

	m, err := ctl.client.MedicalCare.
		Query().
		Where(medicalcare.IDEQ(int(obj.MedicalCare))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "MedicalCare not found",
		})
		return
	}

	e, err := ctl.client.Employee.
		Query().
		Where(employee.IDEQ(int(obj.Employee))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Employee not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.PatientBirthday)
	p, err := ctl.client.Patient.
		Create().
		SetPatientID(obj.PatientID).
		SetPatientName(obj.PatientName).
		SetPatientCardID(obj.PatientCardID).
		SetGender(g).
		SetPatientAddress(obj.PatientAddress).
		SetPatientTel(obj.PatientTel).
		SetPatientBirthday(time).
		SetPatientAge(obj.PatientAge).
		SetDisease(d).
		SetMedicalcare(m).
		SetEmployee(e).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   p,
	})

}

// GetPatient handles GET requests to retrieve a Patient entity
// @Summary Get a Patient entity by ID
// @Description get Patient by ID
// @ID get-Patient
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [get]
func (ctl *PatientController) GetPatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListPatient handles request to get a list of Patient entities
// @Summary List Patient entities
// @Description list Patient entities
// @ID list-Patient
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [get]
func (ctl *PatientController) ListPatient(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	patients, err := ctl.client.Patient.
		Query().
		WithGender().
		WithDisease().
		WithMedicalcare().
		WithEmployee().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patients)
}

// NewPatientController creates and patients handles for the user controller
func NewPatientController(router gin.IRouter, client *ent.Client) *PatientController {
	pc := &PatientController{
		client: client,
		router: router,
	}
	pc.register()
	return pc
}

// InitUserController patient routes to the main engine
func (ctl *PatientController) register() {
	patients := ctl.router.Group("/patients")

	patients.POST("", ctl.PatientCreate)
	patients.GET(":id", ctl.GetPatient)
	patients.GET("", ctl.ListPatient)
}
