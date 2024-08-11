package handler

import (
	pb "api-gateway/genproto/health_analytics"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Add a new medical record
// @Description This endpoint allows you to add a new medical record to the system.
// @Tags MedicalRecords
// @Accept json
// @Produce json
// @Param medicalRecord body health_analytics.AddMedicalRecordRequest true "Medical Record Data"
// @Success 202 {object} health_analytics.AddMedicalRecordResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_records [post]
func (h *Handler) AddMedicalRecord(c *gin.Context) {
	req := pb.AddMedicalRecordRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.HealthService.AddMedicalRecord(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to add medical record error %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary Get a medical record by ID
// @Description This endpoint allows you to fetch a medical record by its ID.
// @Tags MedicalRecords
// @Accept json
// @Produce json
// @Param id path string true "Medical Record ID"
// @Success 202 {object} health_analytics.GetMedicalRecordResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_records/{id} [get]
func (h *Handler) GetMedicalRecord(c *gin.Context){
	req := pb.GetMedicalRecordRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.GetMedicalRecord(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("HealthService yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary Update a medical record
// @Description This endpoint allows you to update an existing medical record.
// @Tags MedicalRecords
// @Accept json
// @Produce json
// @Param medicalRecord body health_analytics.UpdateMedicalRecordRequest true "Updated Medical Record Data"
// @Success 202 {object} health_analytics.UpdateMedicalRecordResponse
// @Failure 400 {object} string "Bad Request"}
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_records [put]
func (h *Handler) UpdateMedicalRecord(c *gin.Context){
	req := pb.UpdateMedicalRecordRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.HealthService.UpdateMedicalRecord(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("UpdateMedicalRecord yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary Delete a medical record by ID
// @Description This endpoint allows you to delete a medical record by its ID.
// @Tags MedicalRecords
// @Accept json
// @Produce json
// @Param id path string true "Medical Record ID"
// @Success 202 {object} health_analytics.DeleteMedicalRecordResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_records/{id} [delete]
func (h *Handler) DeleteMedicalRecord(c *gin.Context){
	req := pb.DeleteMedicalRecordRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.DeleteMedicalRecord(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("DeleteMedicalRecord yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary List all medical records for a user
// @Description This endpoint allows you to list all medical records associated with a specific user ID.
// @Tags MedicalRecords
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 202 {object} health_analytics.ListMedicalRecordsResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_records/user/{userId} [get]
func (h *Handler) ListMedicalRecords(c *gin.Context){
	req := pb.ListMedicalRecordsRequest{
		UserId: c.Param("userId"),
	}

	resp, err := h.HealthService.ListMedicalRecords(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("ListMedicalRecords yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary Turmush tarzi ma'lumotlarini qo'shish
// @Description Yangi turmush tarzi ma'lumotlarini qo'shadi
// @Tags LifestyleData
// @Accept  json
// @Produce  json
// @Param data body health_analytics.AddLifestyleDataRequest true "Turmush tarzi ma'lumotlari"
// @Success 202 {object} health_analytics.AddLifestyleDataResponse
// @Failure 400 {object} string "Noto'g'ri so'rov"
// @Failure 500 {object} string "Ichki server xatoligi"
// @Router /health/lifestyle [post]
func (h *Handler) AddLifestyleData(c *gin.Context){
	req:=pb.AddLifestyleDataRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.HealthService.AddLifestyleData(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("AddLifestyleData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary Turmush tarzi ma'lumotlarini olish
// @Description Turmush tarzi ma'lumotlarini ID bo'yicha olish
// @Tags LifestyleData
// @Accept  json
// @Produce  json
// @Param id path string true "Turmush tarzi ma'lumotlari ID si"
// @Success 202 {object} health_analytics.GetLifestyleDataResponse
// @Failure 500 {object} string "Ichki server xatoligi"
// @Router /health/lifestyle/{id} [get]
func (h *Handler) GetLifestyleData(c *gin.Context){
	req:=pb.GetLifestyleDataRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.GetLifestyleData(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetLifestyleData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary Turmush tarzi ma'lumotlarini yangilash
// @Description Mavjud turmush tarzi ma'lumotlarini yangilash
// @Tags LifestyleData
// @Accept  json
// @Produce  json
// @Param data body health_analytics.UpdateLifestyleDataRequest true "Yangilanadigan turmush tarzi ma'lumotlari"
// @Success 202 {object} health_analytics.UpdateLifestyleDataResponse
// @Failure 400 {object} string "Noto'g'ri so'rov"
// @Failure 500 {object} string "Ichki server xatoligi"
// @Router /health/lifestyle [put]
func (h *Handler) UpdateLifestyleData(c *gin.Context){
	req:=pb.UpdateLifestyleDataRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.HealthService.UpdateLifestyleData(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("UpdateLifestyleData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary Turmush tarzi ma'lumotlarini o'chirish
// @Description Turmush tarzi ma'lumotlarini ID bo'yicha o'chirish
// @Tags LifestyleData
// @Accept  json
// @Produce  json
// @Param id path string true "Turmush tarzi ma'lumotlari ID si"
// @Success 202 {object} health_analytics.DeleteLifestyleDataResponse
// @Failure 500 {object} string "Ichki server xatoligi"
// @Router /health/lifestyle/{id} [delete]
func (h *Handler) DeleteLifestyleData(c *gin.Context){
	req:=pb.DeleteLifestyleDataRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.DeleteLifestyleData(c,&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("DeleteLifestyleData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

