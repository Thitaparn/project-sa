package controllers
 
import (
   "context"
   "strconv"
   "github.com/TP/app/ent"
   "github.com/TP/app/ent/gender"
   "github.com/gin-gonic/gin"
)
 
// GenderController defines the struct for the Gender controller
type GenderController struct {
   client *ent.Client
   router gin.IRouter
}

// GenderCreate handles POST requests for adding Gender entities
// @Summary Create Gender
// @Description Create Gender
// @ID create-Gender
// @Accept   json
// @Produce  json
// @Param Gender body ent.Gender true "Gender entity"
// @Success 200 {object} ent.Gender
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
func (ctl *GenderController) GenderCreate(c *gin.Context) {
	obj := ent.Gender{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Gender binding failed",
		})
		return
	}
  
	g, err := ctl.client.Gender.
		Create().
		SetGenderName(obj.GenderName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, g)
 }
 

// GetGender handles GET requests to retrieve a Gender entity
// @Summary Get a Gender entity by ID
// @Description get Gender by ID
// @ID get-Gender
// @Produce  json
// @Param id path int true "Gender ID"
// @Success 200 {object} ent.Gender
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /genders/{id} [get]
func (ctl *GenderController) GetGender(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, g)
 }

// ListGender handles request to get a list of Gender entities
// @Summary List Gender entities
// @Description list Gender entities
// @ID list-Gender
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Gender
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /genders [get]
func (ctl *GenderController) ListGender(c *gin.Context) {
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
  
	genders, err := ctl.client.Gender.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, genders)
 }
 

// NewGenderController creates and registers handles for the Gender controller
func NewGenderController(router gin.IRouter, client *ent.Client) *GenderController {
	gc := &GenderController{
		client: client,
		router: router,
	}
	gc.register()
	return gc
 }
  
// InitGenderController registers routes to the main engine
 func (ctl *GenderController) register() {
	genders := ctl.router.Group("/genders")
	
	genders.POST("", ctl.GenderCreate)
	genders.GET(":id", ctl.GetGender)
	genders.GET("", ctl.ListGender)

 }
	  