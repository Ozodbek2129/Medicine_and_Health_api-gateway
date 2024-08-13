package handler

import (
	pb "api-gateway/genproto/health_analytics"
	pbu "api-gateway/genproto/user"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

// @Summary Add a new medical record
// @Description This endpoint allows you to add a new medical record to the system.
// @Tags MedicalRecords
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param medicalRecord body health_analytics.AddMedicalRecordRequest true "Medical Record Data"
// @Success 202 {object} health_analytics.AddMedicalRecordResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_recordsAdd [post]
func (h *Handler) AddMedicalRecord(c *gin.Context) {
	req := pb.AddMedicalRecordRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req1 := pbu.UserId{
		Userid: req.UserId,
	}

	resp1, err := h.UserService.IdCheck(c, &req1)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Id ni tekshirishda xatolik %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !resp1.B {
		h.Log.Warn(fmt.Sprintf("Foydalanuvchi ID topilmadi: %v", req.UserId))
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Foydalanuvchi topilmadi"})
		return
	}

	resp, err := h.HealthService.AddMedicalRecord(c, &req)
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
// @Security ApiKeyAuth
// @Param id path string true "Medical Record ID"
// @Success 202 {object} health_analytics.GetMedicalRecordResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_recordsGet/{id} [get]
func (h *Handler) GetMedicalRecord(c *gin.Context) {
	req := pb.GetMedicalRecordRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.GetMedicalRecord(c, &req)
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
// @Security ApiKeyAuth
// @Param medicalRecord body health_analytics.UpdateMedicalRecordRequest true "Updated Medical Record Data"
// @Success 202 {object} health_analytics.UpdateMedicalRecordResponse
// @Failure 400 {object} string "Bad Request"}
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_recordsUp [put]
func (h *Handler) UpdateMedicalRecord(c *gin.Context) {
	req := pb.UpdateMedicalRecordRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.HealthService.UpdateMedicalRecord(c, &req)
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
// @Security ApiKeyAuth
// @Param id path string true "Medical Record ID"
// @Success 202 {object} health_analytics.DeleteMedicalRecordResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_records/{id} [delete]
func (h *Handler) DeleteMedicalRecord(c *gin.Context) {
	req := pb.DeleteMedicalRecordRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.DeleteMedicalRecord(c, &req)
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
// @Security ApiKeyAuth
// @Param userId path string true "User ID"
// @Success 202 {object} health_analytics.ListMedicalRecordsResponse
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/medical_records/user/{userId} [get]
func (h *Handler) ListMedicalRecords(c *gin.Context) {
	req := pb.ListMedicalRecordsRequest{
		UserId: c.Param("userId"),
	}

	resp, err := h.HealthService.ListMedicalRecords(c, &req)
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
// @Security ApiKeyAuth
// @Param data body health_analytics.AddLifestyleDataRequest true "Turmush tarzi ma'lumotlari"
// @Success 202 {object} health_analytics.AddLifestyleDataResponse
// @Failure 400 {object} string "Noto'g'ri so'rov"
// @Failure 500 {object} string "Ichki server xatoligi"
// @Router /health/lifestyleAdd [post]
func (h *Handler) AddLifestyleData(c *gin.Context) {
	req := pb.AddLifestyleDataRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req1 := pbu.UserId{
		Userid: req.UserId,
	}

	resp1, err := h.UserService.IdCheck(c, &req1)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Id ni tekshirishda xatolik %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !resp1.B {
		h.Log.Warn(fmt.Sprintf("Foydalanuvchi ID topilmadi: %v", req.UserId))
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Foydalanuvchi topilmadi"})
		return
	}

	resp, err := h.HealthService.AddLifestyleData(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("AddLifestyleData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary Get All Lifestyle Data
// @Description Foydalanuvchilarning barcha lifestyle (turmush tarzi) ma'lumotlarini olish uchun ishlatiladi.
// @Tags LifestyleData
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int true "Cheklov (sahifadagi yozuvlar soni)"
// @Param page query int true "Sahifa (qaysi sahifadagi yozuvlar ko'rinadi)"
// @Success 202 {object} health_analytics.GetAllLifestyleDataResponse "Ma'lumotlar muvaffaqiyatli qaytarildi"
// @Failure 400 {object} string "Invalid limit or page parameter" "Noto'g'ri limit yoki sahifa parametri"
// @Failure 500 {object} string "Internal Server Error" "Serverda xatolik yuz berdi"
// @Router /health/getalllifestyledata/{limit}/{page} [get]
func (h *Handler) GetAllLifestyleData(c *gin.Context) {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	req := pb.GetAllLifestyleDataRequest{
		Limit: limit,
		Page:  page,
	}

	resp, err := h.HealthService.GetAllLifestyleData(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetAllLifestyleData yuborishda xatolik: %v", err))
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
// @Security ApiKeyAuth
// @Param id path string true "Turmush tarzi ma'lumotlari ID si"
// @Success 202 {object} health_analytics.GetLifestyleDataResponse
// @Failure 500 {object} string "Ichki server xatoligi"
// @Router /health/lifestyleGet/{id} [get]
func (h *Handler) GetLifestyleData(c *gin.Context) {
	req := pb.GetLifestyleDataRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.GetLifestyleData(c, &req)
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
// @Security ApiKeyAuth
// @Param data body health_analytics.UpdateLifestyleDataRequest true "Yangilanadigan turmush tarzi ma'lumotlari"
// @Success 202 {object} health_analytics.UpdateLifestyleDataResponse
// @Failure 400 {object} string "Noto'g'ri so'rov"
// @Failure 500 {object} string "Ichki server xatoligi"
// @Router /health/lifestyleUp [put]
func (h *Handler) UpdateLifestyleData(c *gin.Context) {
	req := pb.UpdateLifestyleDataRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.HealthService.UpdateLifestyleData(c, &req)
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
// @Security ApiKeyAuth
// @Param id path string true "Turmush tarzi ma'lumotlari ID si"
// @Success 202 {object} health_analytics.DeleteLifestyleDataResponse
// @Failure 500 {object} string "Ichki server xatoligi"
// @Router /health/lifestyleDel/{id} [delete]
func (h *Handler) DeleteLifestyleData(c *gin.Context) {
	req := pb.DeleteLifestyleDataRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.DeleteLifestyleData(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("DeleteLifestyleData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary      Add Wearable Data
// @Description  Ushbu endpoint yangi kiyiladigan qurilma ma'lumotlarini qo'shadi.
// @Tags         wearable-data
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        data  body  health_analytics.AddWearableDataRequest  true  "Kiyiladigan qurilma ma'lumotlari"
// @Success      202   {object}  health_analytics.AddWearableDataResponse
// @Failure      400   {object}  string "bodydan malumotlarni olishda xatolik: <error_message>"
// @Failure      500   {object}  string "AddWearableData yuborishda xatolik: <error_message>"
// @Router       /health/wearable-dataAdd [post]
func (h *Handler) AddWearableData(c *gin.Context) {
	req := pb.AddWearableDataRequest{}

	// Request body'dan ma'lumotlarni olish
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// User ID tekshirish
	req1 := pbu.UserId{
		Userid: req.UserId,
	}

	resp1, err := h.UserService.IdCheck(c, &req1)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Id ni tekshirishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !resp1.B {
		h.Log.Warn(fmt.Sprintf("Foydalanuvchi ID topilmadi: %v", req.UserId))
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Foydalanuvchi topilmadi"})
		return
	}

	// UUID va vaqt yaratish
	id := uuid.NewString()
	vaqt := time.Now().Format(time.RFC3339)

	// Xabarni yaratish
	wearableData := pb.AddWearableDataRequest{
		UserId:            req.UserId,
		DeviceType:        req.DeviceType,
		DataType:          req.DataType,
		DataValue:         req.DataValue,
		RecordedTimestamp: req.RecordedTimestamp,
	}

	messageBody, err := json.Marshal(struct {
		Id string `json:"id"`
		*pb.AddWearableDataRequest
	}{
		Id:                     id,
		AddWearableDataRequest: &wearableData,
	})
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to marshal request to JSON: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	q, err := h.RabbitMQChannel.QueueDeclare(
		"wearable_data_queue", // name
		false,                 // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)

	// RabbitMQ ga xabar yuborish
	err = h.RabbitMQChannel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageBody),
		})
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to publish a message to RabbitMQ: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response yuborish
	c.JSON(http.StatusAccepted, &pb.AddWearableDataResponse{
		WearableData: &pb.WearableData{
			Id:                id,
			UserId:            req.UserId,
			DeviceType:        req.DeviceType,
			DataType:          req.DataType,
			DataValue:         req.DataValue,
			RecordedTimestamp: req.RecordedTimestamp,
			CreatedAt:         vaqt,
			UpdatedAt:         vaqt,
		},
	})
}

// ------------------------------------------------------------------------------------------------

// func (h *Handler) AddWearableData(c *gin.Context) {
// 	req := pb.AddWearableDataRequest{}

// 	err := json.NewDecoder(c.Request.Body).Decode(&req)
// 	if err != nil {
// 		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	req1 := pbu.UserId{
// 		Userid: req.UserId,
// 	}

// 	resp1, err := h.UserService.IdCheck(c, &req1)
// 	if err != nil {
// 		h.Log.Error(fmt.Sprintf("Id ni tekshirishda xatolik %v", err))
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if !resp1.B {
// 		h.Log.Warn(fmt.Sprintf("Foydalanuvchi ID topilmadi: %v", req.UserId))
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Foydalanuvchi topilmadi"})
// 		return
// 	}

// 	resp, err := h.HealthService.AddWearableData(c, &req)
// 	if err != nil {
// 		h.Log.Error(fmt.Sprintf("AddWearableData yuborishda xatolik: %v", err))
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusAccepted, resp)
// }

// ------------------------------------------------------------------------------------------------

// @Summary Get All Wearable Data
// @Description Ushbu endpoint foydalanuvchilarning barcha wearable (kiyiladigan) ma'lumotlarini olish uchun ishlatiladi.
// @Tags wearable-data
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int true "Cheklov (sahifadagi yozuvlar soni)"
// @Param page query int true "Sahifa (qaysi sahifadagi yozuvlar ko'rinadi)"
// @Success 202 {object} health_analytics.GetAllWearableDataResponse "Ma'lumotlar muvaffaqiyatli qaytarildi"
// @Failure 400 {object} string "Invalid limit or page parameter" "Noto'g'ri limit yoki sahifa parametri"
// @Failure 500 {object} string "Internal Server Error" "Serverda xatolik yuz berdi"
// @Router /health/wearabledata/{limit}/{page} [get]
func (h *Handler) GetAllWearableData(c *gin.Context) {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	req := pb.GetAllWearableDataRequest{
		Limit: limit,
		Page:  page,
	}

	resp, err := h.HealthService.GetAllWearableData(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetAllWearableData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary      Get Wearable Data
// @Description  Ushbu endpoint ID bo'yicha kiyiladigan qurilma ma'lumotlarini qaytaradi.
// @Tags         wearable-data
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id    path      string  true  "Kiyiladigan qurilma ma'lumotlari ID'si"
// @Success      202   {object}  health_analytics.GetWearableDataResponse
// @Failure      500   {object}  string "GetWearableData yuborishda xatolik: <error_message>"
// @Router       /health/wearable-dataGet/{id} [get]
func (h *Handler) GetWearableData(c *gin.Context) {
	req := pb.GetWearableDataRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.GetWearableData(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetWearableData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary      Update Wearable Data
// @Description  Ushbu endpoint mavjud kiyiladigan qurilma ma'lumotlarini yangilaydi.
// @Tags         wearable-data
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        data  body  health_analytics.UpdateWearableDataRequest  true  "Yangilangan kiyiladigan qurilma ma'lumotlari"
// @Success      202   {object}  health_analytics.UpdateWearableDataResponse
// @Failure      400   {object}  string "bodydan malumotlarni olishda xatolik: <error_message>"
// @Failure      500   {object}  string "UpdateWearableData yuborishda xatolik: <error_message>"
// @Router       /health/wearable-dataUp [put]
func (h *Handler) UpdateWearableData(c *gin.Context) {
	req := pb.UpdateWearableDataRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.HealthService.UpdateWearableData(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("UpdateWearableData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary      Delete Wearable Data
// @Description  Ushbu endpoint berilgan ID bo'yicha kiyiladigan qurilma ma'lumotlarini o'chiradi.
// @Tags         wearable-data
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id    path      string  true  "Kiyiladigan qurilma ma'lumotlari ID'si"
// @Success      202   {object}  health_analytics.DeleteWearableDataResponse
// @Failure      500   {object}  string "DeleteWearableData yuborishda xatolik: <error_message>"
// @Router       /health/wearable-dataDel/{id} [delete]
func (h *Handler) DeleteWearableData(c *gin.Context) {
	req := pb.DeleteWearableDataRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.DeleteWearableData(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("DeleteWearableData yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary      Sog'liq bo'yicha tavsiyalar yaratish
// @Description  Foydalanuvchi uchun sog'liq bo'yicha tavsiyalarni yaratadi.
// @Tags         Health
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        body  body  health_analytics.GenerateHealthRecommendationsRequest  true  "Tavsiyalar yaratish uchun ma'lumotlar"
// @Success      202  {object}  health_analytics.GenerateHealthRecommendationsResponse
// @Failure      400  {object}  string  "Xatolik: Notog'ri so'rov"
// @Failure      500  {object}  string  "Xatolik: Ichki server xatosi"
// @Router       /health/recommendationsAdd [post]
func (h *Handler) GenerateHealthRecommendations(c *gin.Context) {
	req := pb.GenerateHealthRecommendationsRequest{}

	// Request body'dan ma'lumotlarni olish
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// User ID tekshirish
	req1 := pbu.UserId{
		Userid: req.UserId,
	}

	resp1, err := h.UserService.IdCheck(c, &req1)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Id ni tekshirishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !resp1.B {
		h.Log.Warn(fmt.Sprintf("Foydalanuvchi ID topilmadi: %v", req.UserId))
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Foydalanuvchi topilmadi"})
		return
	}

	// UUID va vaqt yaratish
	id := uuid.NewString()
	vaqt := time.Now().Format(time.RFC3339)

	// Xabarni yaratish
	healthData := pb.GenerateHealthRecommendationsRequest{
		UserId:    req.UserId,
		RecommendationType: req.RecommendationType,
		Description: req.Description,
		Priority: req.Priority,
	}

	messageBody, err := json.Marshal(struct {
		Id string `json:"id"`
		*pb.GenerateHealthRecommendationsRequest
	}{
		Id:                                id,
		GenerateHealthRecommendationsRequest: &healthData,
	})
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to marshal request to JSON: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// RabbitMQ ga xabar yuborish uchun navbat yaratish
	q, err := h.RabbitMQChannel.QueueDeclare(
		"health_recommendations_queue", // name
		false,                          // durable
		false,                          // delete when unused
		false,                          // exclusive
		false,                          // no-wait
		nil,                            // arguments
	)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to declare a queue: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Queue yaratishda xatolik"})
		return
	}

	// RabbitMQ ga xabar yuborish
	err = h.RabbitMQChannel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(messageBody),
		})
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to publish a message to RabbitMQ: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "RabbitMQ ga xabar yuborishda xatolik"})
		return
	}

	// Response yuborish
	c.JSON(http.StatusAccepted, gin.H{"message": "Health recommendations request yuborildi", "id": id, "created_at": vaqt})
}

// func (h *Handler) GenerateHealthRecommendations(c *gin.Context) {
// 	req := pb.GenerateHealthRecommendationsRequest{}

// 	err := json.NewDecoder(c.Request.Body).Decode(&req)
// 	if err != nil {
// 		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	req1 := pbu.UserId{
// 		Userid: req.UserId,
// 	}

// 	resp1, err := h.UserService.IdCheck(c, &req1)
// 	if err != nil {
// 		h.Log.Error(fmt.Sprintf("Id ni tekshirishda xatolik %v", err))
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if !resp1.B {
// 		h.Log.Warn(fmt.Sprintf("Foydalanuvchi ID topilmadi: %v", req.UserId))
// 		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Foydalanuvchi topilmadi"})
// 		return
// 	}

// 	resp, err := h.HealthService.GenerateHealthRecommendations(c, &req)
// 	if err != nil {
// 		h.Log.Error(fmt.Sprintf("GenerateHealthRecommendations yuborishda xatolik: %v", err))
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusAccepted, resp)
// }

// @Summary Generate Health Recommendations by ID
// @Description Ushbu endpoint berilgan ID bo'yicha o'chirilmagan (deleted_at "0") sog'liq tavsiyalarini oladi.
// @Tags Health
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Health Recommendation ID"
// @Success 202 {object} health_analytics.GenerateHealthRecommendationsIdResponse
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/recommendations/{id} [get]
func (h *Handler) GenerateHealthRecommendationsId(c *gin.Context) {
	req := pb.GenerateHealthRecommendationsIdRequest{
		Id: c.Param("id"),
	}

	resp, err := h.HealthService.GenerateHealthRecommendationsId(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GenerateHealthRecommendationsId yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary      Real vaqt sog'liq monitoringi
// @Description  Foydalanuvchining real vaqt sog'liq monitoringini qaytaradi.
// @Tags         Health
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        user_id  path  string  true  "Foydalanuvchi ID"
// @Success      202  {object}  health_analytics.GetRealtimeHealthMonitoringResponse
// @Failure      500  {object}  string  "Xatolik: Ichki server xatosi"
// @Router       /health/monitoring/{user_id}/realtime [get]
func (h *Handler) GetRealtimeHealthMonitoring(c *gin.Context) {
	req := pb.GetRealtimeHealthMonitoringRequest{
		UserId: c.Param("user_id"),
	}

	requ := pbu.UserId{
		Userid: req.UserId,
	}

	respUser, err := h.UserService.GetByUserId(c, &requ)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetByUserId yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.HealthService.GetRealtimeHealthMonitoring(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetRealtimeHealthMonitoring yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &pb.GetRealtimeHealthMonitoringResponse{
		FirstName:          respUser.FirstName,
		LastName:           respUser.LastName,
		RecommendationType: resp.RecommendationType,
		Description:        resp.Description,
		Priority:           resp.Priority,
	}
	c.JSON(http.StatusAccepted, res)
}

// @Summary      Kunlik sog'liq hisobotini olish
// @Description  Foydalanuvchining belgilangan kun uchun sog'liq hisobotini qaytaradi.
// @Tags         Health
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        user_id  path  string  true  "Foydalanuvchi ID"
// @Param        date     query  string  true  "Kun (format: yyyy/mm/dd)"
// @Success      202  {object}  health_analytics.GetDailyHealthSummaryResponse
// @Failure      500  {object}  string  "Xatolik: Ichki server xatosi"
// @Router       /health/summary/{user_id}/daily [get]
func (h *Handler) GetDailyHealthSummary(c *gin.Context) {
	fmt.Println("sdfghjkl;")
	req := pb.GetDailyHealthSummaryRequest{
		UserId: c.Param("user_id"),
		Date:   c.Query("date"),
	}

	requ := pbu.UserId{
		Userid: req.UserId,
	}

	respUser, err := h.UserService.GetByUserId(c, &requ)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetByUserId yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.HealthService.GetDailyHealthSummary(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetDailyHealthSummary yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := &pb.GetDailyHealthSummaryResponse{
		FirstName:          respUser.FirstName,
		LastName:           respUser.LastName,
		RecommendationType: resp.RecommendationType,
		Description:        resp.Description,
		Priority:           resp.Priority,
	}
	c.JSON(http.StatusAccepted, res)
}

// @Summary      Haftalik sog'liq hisobotini olish
// @Description  Foydalanuvchining belgilangan haftalik sog'liq hisobotini qaytaradi.
// @Tags         Health
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        user_id    path  string  true  "Foydalanuvchi ID"
// @Param        start_date query  string  true  "Boshlang'ich sana (format: yyyy/mm/dd)"
// @Success      202  {object}  health_analytics.GetWeeklyHealthSummaryResponse
// @Failure      500  {object}  string  "Xatolik: Ichki server xatosi"
// @Router       /health/summary/{user_id}/weekly [get]
func (h *Handler) GetWeeklyHealthSummary(c *gin.Context) {
	req := pb.GetWeeklyHealthSummaryRequest{
		UserId:    c.Param("user_id"),
		StartDate: c.Query("start_date"),
	}

	resp, err := h.HealthService.GetWeeklyHealthSummary(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("GetWeeklyHealthSummary yuborishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}
