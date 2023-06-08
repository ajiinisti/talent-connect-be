package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/delivery/api/request"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type MentoringScheduleController struct {
	router *gin.Engine
	uc     usecase.MentoringScheduleUsecase
	api.BaseApi
}

func (u *MentoringScheduleController) listHandler(c *gin.Context) {
	mentoringSchedules, err := u.uc.FindAll()
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, mentoringSchedules, "OK")
}

func (r *MentoringScheduleController) createHandler(c *gin.Context) {
	var payload request.MentoringScheduleRequest
	if err := r.ParseRequestBody(c, &payload); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.uc.SaveData(&payload); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r.NewSuccessSingleResponse(c, payload, "OK")
}

func (r *MentoringScheduleController) updateHandler(c *gin.Context) {
	var payload request.MentoringScheduleRequest
	if err := r.ParseRequestBody(c, &payload); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.uc.SaveData(&payload); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r.NewSuccessSingleResponse(c, payload, "OK")
}

func (r *MentoringScheduleController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := r.uc.DeleteData(id); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func (u *MentoringScheduleController) listMentorHandler(c *gin.Context) {
	id := c.Param("id")
	mentoringSchedules, err := u.uc.FindScheduleByMentorId(id)
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, mentoringSchedules, "OK")
}

func (u *MentoringScheduleController) listMenteeHandler(c *gin.Context) {
	id := c.Param("id")
	mentoringSchedules, err := u.uc.FindScheduleByMenteeId(id)
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, mentoringSchedules, "OK")
}

func (u *MentoringScheduleController) getHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := u.uc.FindById(id)
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, payload, "OK")
}

func NewMentoringScheduleController(r *gin.Engine, uc usecase.MentoringScheduleUsecase) *MentoringScheduleController {
	controller := MentoringScheduleController{
		router: r,
		uc:     uc,
	}
	r.GET("/mentoring-schedules", controller.listHandler)
	r.GET("/mentoring-schedules/:id", controller.getHandler)
	r.GET("/mentoring-schedules/mentor/:id", controller.listMentorHandler)
	r.GET("/mentoring-schedules/mentee/:id", controller.listMenteeHandler)
	r.PUT("/mentoring-schedules", controller.updateHandler)
	r.POST("/mentoring-schedules", controller.createHandler)
	r.DELETE("/mentoring-schedules/:id", controller.deleteHandler)
	return &controller
}
