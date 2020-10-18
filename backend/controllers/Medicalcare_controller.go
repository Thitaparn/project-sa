package controllers
 
import (
   "context"
   "strconv"
   "github.com/gin-gonic/gin"
   "github.com/TP/app/ent"
   "github.com/TP/app/ent/medicalcare"
   
)
 
// MedicalCareController defines the struct for the MedicalCare controller
type MedicalCareController struct {
   client *ent.Client
   router gin.IRouter
}

// MedicalCareCreate handles POST requests for adding MedicalCare entities
// @Summary Create MedicalCare
// @Description Create MedicalCare
// @ID create-MedicalCare
// @Accept   json
// @Produce  json
// @Param MedicalCare body ent.MedicalCare true "MedicalCare entity"
// @Success 200 {object} ent.MedicalCare
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalcares [post]
func (ctl *MedicalCareController) MedicalCareCreate(c *gin.Context) {
	obj := ent.MedicalCare{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "MedicalCare binding failed",
		})
		return
	}
  
	m, err := ctl.client.MedicalCare.
		Create().	
		SetMedicalcareName(obj.MedicalcareName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, m)
 }
 
// GetMedicalCare handles GET requests to retrieve a MedicalCare entity
// @Summary Get a MedicalCare entity by ID
// @Description get MedicalCare by ID
// @ID get-MedicalCare
// @Produce  json
// @Param id path int true "MedicalCare ID"
// @Success 200 {object} ent.MedicalCare
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalcares/{id} [get]
func (ctl *MedicalCareController) GetMedicalCare(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	m, err := ctl.client.MedicalCare.
		Query().
		Where(medicalcare.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, m)
 }

// ListMedicalCare handles request to get a list of MedicalCare entities
// @Summary List MedicalCare entities
// @Description list MedicalCare entities
// @ID list-MedicalCare
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.MedicalCare
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalcares [get]
func (ctl *MedicalCareController) ListMedicalCare(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	medicalcares, err := ctl.client.MedicalCare.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, medicalcares)
 }
 
// NewMedicalCareController creates and registers handles for the MedicalCare controller
func NewMedicalCareController(router gin.IRouter, client *ent.Client) *MedicalCareController {
	mc := &MedicalCareController{
		client: client,
		router: router,
	}
	mc.register()
	return mc
 }
  
// InitMedicalCareController registers routes to the main engine
 func (ctl *MedicalCareController) register() {
	medicalcares := ctl.router.Group("/medicalcares")
	medicalcares.GET("", ctl.ListMedicalCare)
	// CRUD
	medicalcares.POST("", ctl.MedicalCareCreate)
	medicalcares.GET(":id", ctl.GetMedicalCare)
 }
	  