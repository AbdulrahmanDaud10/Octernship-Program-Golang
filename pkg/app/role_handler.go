package app

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AbdulrahmanDaud10/diveInputCalories/pkg/api"
	"github.com/AbdulrahmanDaud10/diveInputCalories/pkg/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// create Role
func CreateRole(c *gin.Context) {
	var Role api.Role
	c.BindJSON(&Role)
	err := repository.CreateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// get Roles
func GetRoles(c *gin.Context) {
	var Role []api.Role
	err := repository.GetRoles(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// get Role by id
func GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Role api.Role
	err := repository.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// update Role
func UpdateRole(c *gin.Context) {
	var Role api.Role
	id, _ := strconv.Atoi(c.Param("id"))
	err := repository.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&Role)
	err = repository.UpdateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}
