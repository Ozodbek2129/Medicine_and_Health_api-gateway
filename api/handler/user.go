package handler

import (
	pb "api-gateway/genproto/user"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Foydalanuvchiga yangi bildirishnoma qo'shish
// @Description  Foydalanuvchiga yangi bildirishnoma qo'shish uchun ishlatiladi.
// @Tags         Notifications
// @Accept       json
// @Produce      json
// @Param        notifications body user.NotificationsAddRequest true "Notifications Add Request Body"
// @Success      202 {object} user.NotificationsAddResponse
// @Failure      400 {object} string "Error message"
// @Failure      500 {object} string "Error message"
// @Router       /user/notifications [post]
func (h *Handler) NotificationsAdd(c *gin.Context) {
	req := pb.NotificationsAddRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.UserService.NotificationsAdd(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to add medical record error %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary      Foydalanuvchi bildirishnomalarini olish
// @Description  Foydalanuvchi uchun saqlangan barcha bildirishnomalarni olish uchun ishlatiladi.
// @Tags         Notifications
// @Accept       json
// @Produce      json
// @Param        user_id   path  string  true  "Foydalanuvchi IDsi"
// @Success      202 {object} user.NotificationsGetResponse
// @Failure      500 {object} string "Error message"
// @Router       /user/notifications/{user_id} [get]
func (h *Handler) NotificationsGet(c *gin.Context){
	req:=pb.NotificationsGetRequest{
		UserId: c.Param("user_id"),
	}

	resp, err := h.UserService.NotificationsGet(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to add medical record error %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// @Summary      Foydalanuvchi bildirishnomasini yangilash
// @Description  Foydalanuvchi uchun mavjud bildirishnomani yangilash uchun ishlatiladi.
// @Tags         Notifications
// @Accept       json
// @Produce      json
// @Param        notifications body user.NotificationsPutRequest true "Notifications Put Request Body"
// @Success      202 {object} user.NotificationsPutResponse
// @Failure      400 {object} string "Error message"
// @Failure      500 {object} string "Error message"
// @Router       /user/notificationsPut [put]
func (h *Handler) NotificationsPut(c *gin.Context){
	req:=pb.NotificationsPutRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("bodydan malumotlarni olishda xatolik: %v", err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.UserService.NotificationsPut(c, &req)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to add medical record error %v", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, resp)
}