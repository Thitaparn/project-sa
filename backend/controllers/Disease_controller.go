package controllers

import (
	"context"
	"strconv"

	"github.com/TP/app/ent"
	"github.com/TP/app/ent/disease"
	"github.com/gin-gonic/gin"
)

// DiseaseController defines the struct for the Disease controller
type DiseaseController struct {
	client *ent.Client
	router gin.IRouter
}

// DiseaseCreate handles POST requests for adding Disease entities
// @Summary Create Disease
// @Description Create Disease
// @ID create-Disease
// @Accept   json
// @Produce  json
// @Param Disease body ent.Disease true "Disease entity"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [post]
func (ctl *DiseaseController) DiseaseCreate(c *gin.Context) {
	obj := ent.Disease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Disease binding failed",
		})
		return
	}

	d, err := ctl.client.Disease.
		Create().
		SetDiseaseName(obj.DiseaseName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetDisease handles GET requests to retrieve a Disease entity
// @Summary Get a Disease entity by ID
// @Description get Disease by ID
// @ID get-Disease
// @Produce  json
// @Param id path int true "Disease ID"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [get]
func (ctl *DiseaseController) GetDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	d, err := ctl.client.Disease.
		Query().
		Where(disease.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, d)
}

// ListDisease handles request to get a list of Disease entities
// @Summary List Disease entities
// @Description list Disease entities
// @ID list-Disease
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [get]
func (ctl *DiseaseController) ListDisease(c *gin.Context) {
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

	diseases, err := ctl.client.Disease.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, diseases)
}

// NewDiseaseController creates and registers handles for the Disease controller
func NewDiseaseController(router gin.IRouter, client *ent.Client) *DiseaseController {
	dc := &DiseaseController{
		client: client,
		router: router,
	}
	dc.register()
	return dc
}

// InitDiseaseController registers routes to the main engine
func (ctl *DiseaseController) register() {
	diseases := ctl.router.Group("/diseases")

	diseases.POST("", ctl.DiseaseCreate)
	diseases.GET(":id", ctl.GetDisease)
	diseases.GET("", ctl.ListDisease)
}
