package handlers

import (
	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandlers struct {
	svc services.AuthServiceInterface
}

func NewAuthHandlers() *AuthHandlers {
	return &AuthHandlers{
		svc: services.NewAuthService(),
	}
}

func (h AuthHandlers) Greet(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello!")
}

func (h AuthHandlers) CreateUser(c *gin.Context) {
	var body models.AuthRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("please provide valid input")))
		return
	}

	result, err := h.svc.CreateUser(&body)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}

func (h AuthHandlers) LogoutUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, err := h.svc.LogoutUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}

func (h AuthHandlers) DeleteUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, err := h.svc.DeleteUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}
func (h AuthHandlers) GetUserByID(c *gin.Context) {
	claims, _ := c.Get("claims")
	userID := int(claims.(jwt.MapClaims)["user_id"].(float64))

	result, err := h.svc.GetUserByID(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(result.Status, result)
}
func (h AuthHandlers) LoginUser(c *gin.Context) {
	var body *models.AuthRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("please provide valid input")))
		return
	}

	tokensResponse, err := h.svc.LoginUser(body)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("refresh_token", tokensResponse.RefreshToken, 0, "", config.Envs.WEB_URL, true, true)
	c.JSON(http.StatusOK, &models.Response{Success: true, Status: http.StatusOK, Data: tokensResponse})
}

func (h AuthHandlers) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusUnauthorized, fmt.Errorf("please login again")))
	}

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Envs.JWT_SECRET_KEY), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusUnauthorized, fmt.Errorf("invalid refresh token")))
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusUnauthorized, fmt.Errorf("invalid token claims")))
		return
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusUnauthorized, fmt.Errorf("invalid user_id in token")))
		return
	}
	userID := int(userIDFloat)

	tokensResponse, er := h.svc.GenerateTokens(userID, refreshToken)
	if er != nil {
		c.AbortWithStatusJSON(er.Status, er)
		return
	}

	c.SetCookie("refresh_token", tokensResponse.RefreshToken, 0, "", config.Envs.WEB_URL, true, true)
	c.JSON(http.StatusOK, &models.Response{Success: true, Status: http.StatusOK, Data: tokensResponse})
}
