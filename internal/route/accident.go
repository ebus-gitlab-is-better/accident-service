package route

import (
	"accident-service/internal/biz"
	"context"
	"encoding/json"
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AccidentRoute struct {
	uc *biz.AccidentUseCase
}

func NewAccidentRoute(uc *biz.AccidentUseCase) *AccidentRoute {
	return &AccidentRoute{uc: uc}
}

func (r *AccidentRoute) Register(router *gin.RouterGroup) {
	router.POST("/", r.create)
	router.PUT("/:id", r.update)
	router.GET("/", r.list)
	router.DELETE("/:id", r.delete)
}

type AccidentDTO struct {
	Name      string
	Lat       float64
	Lon       float64
	StartDate time.Time
	EndDate   *time.Time
}

// @Summary	Create accident
// @Accept		json
// @Produce	json
// @Tags		accident
// @Param		dto	body	route.AccidentDTO	true	"dto"
// @Success	200
// @Failure	401
// @Failure	403
// @Failure	500
// @Failure	400
// @Failure	404
// @Router		/accident/ [post]
func (r *AccidentRoute) create(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(400, &gin.H{
			"error": err.Error(),
		})
		return
	}
	dto := AccidentDTO{}

	err = json.Unmarshal(body, &dto)
	if err != nil {
		c.AbortWithStatusJSON(400, &gin.H{
			"error": err.Error(),
		})
		return
	}

	err = r.uc.Create(context.TODO(), &biz.Accident{
		Name:      dto.Name,
		Lat:       dto.Lat,
		Lon:       dto.Lon,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
	})
	if err != nil {
		c.AbortWithStatusJSON(400, &gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(200)
}

// @Summary	Update accident
// @Accept		json
// @Produce	json
// @Tags		accident
// @Param		id	path	int	true	"Accident ID"	Format(uint64)
// @Param		dto	body	route.AccidentDTO	true	"dto"
// @Success	200
// @Failure	401
// @Failure	403
// @Failure	500
// @Failure	400
// @Failure	404
// @Router		/accident/{id} [put]
func (r *AccidentRoute) update(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "parse id error",
		})
		return
	}
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(400, &gin.H{
			"error": err.Error(),
		})
		return
	}
	dto := AccidentDTO{}

	err = json.Unmarshal(body, &dto)
	if err != nil {
		c.AbortWithStatusJSON(400, &gin.H{
			"error": err.Error(),
		})
		return
	}

	err = r.uc.Update(context.TODO(), &biz.Accident{
		Id:        uint64(idUint),
		Name:      dto.Name,
		Lat:       dto.Lat,
		Lon:       dto.Lon,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
	})
	if err != nil {
		c.AbortWithStatusJSON(400, &gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(200)
}

// @Summary	Delete accident
// @Accept		json
// @Produce	json
// @Tags		accident
// @Param		id	path	int	true	"Accident ID"	Format(uint64)
// @Success	200
// @Failure	401
// @Failure	403
// @Failure	500
// @Failure	400
// @Failure	404
// @Router		/accident/{id} [delete]
func (r *AccidentRoute) delete(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "parse id error",
		})
		return
	}

	err = r.uc.Delete(context.TODO(), uint64(idUint))
	if err != nil {
		c.AbortWithStatusJSON(400, &gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(200)
}

type ListDTO struct {
	Accidents []*biz.Accident
	Count     int64
}

// @Summary	List accident
// @Accept		json
// @Produce	json
// @Tags		accident
// @Success	200 {object} route.ListDTO
// @Failure	401
// @Failure	403
// @Failure	500
// @Failure	400
// @Failure	404
// @Router		/accident/ [get]
func (r *AccidentRoute) list(c *gin.Context) {
	list, total, err := r.uc.List(context.TODO())
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "parse id error",
		})
		return
	}
	c.JSON(200, &ListDTO{
		Accidents: list,
		Count:     total,
	})
}
