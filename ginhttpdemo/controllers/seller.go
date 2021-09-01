package controllers

import (
	"errors"
	"ginhttpdemo/database"
	"ginhttpdemo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type SellerRepository struct {
	Db *gorm.DB
}

func New() *SellerRepository {
	db := database.InitDb()
	db.AutoMigrate(&models.Seller{})
	return &SellerRepository{Db: db}
}

func (repository *SellerRepository) CreateSeller(c *gin.Context) {
	var seller models.Seller
	c.BindJSON(&seller)
	err := models.CreateSeller(repository.Db, &seller)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, seller)
}

func (repository *SellerRepository) GetSellers(c *gin.Context) {
	var seller []models.Seller
	err := models.GetSellers(repository.Db, &seller)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, seller)
}

func (repository *SellerRepository) GetSeller(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var seller models.Seller
	err := models.GetSeller(repository.Db, &seller, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, seller)
}

func (repository *SellerRepository) UpdateSeller(c *gin.Context) {
	var seller models.Seller
	id, _ := c.Params.Get("id")
	err := models.GetSeller(repository.Db, &seller, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&seller)
	err = models.UpdateSeller(repository.Db, &seller)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, seller)
}

func (repository *SellerRepository) DeleteSeller(c *gin.Context) {
	var seller models.Seller
	id, _ := c.Params.Get("id")
	err := models.DeleteSeller(repository.Db, &seller, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Seller deleted successfully"})
}