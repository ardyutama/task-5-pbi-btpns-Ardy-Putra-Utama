package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"rakamin-golang/database"
	"rakamin-golang/helpers"
	"rakamin-golang/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type PhotoController struct{}

func NewPhotoController() *PhotoController {
	return &PhotoController{}
}

func (pc *PhotoController) UploadPhoto(c *gin.Context) {
	userID := helpers.ExtractUserID(c)

	var photoInput models.Photo
	if err := c.ShouldBind(&photoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		errorMessages := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to upload user photo", "message": errorMessages})
		return
	}

	splitedFileName := strings.Split(file.Filename, ".")
	fileFormat := splitedFileName[len(splitedFileName)-1]
	path := fmt.Sprint("images/user/", userID, "_", time.Now().Format("010206150405"), ".", fileFormat)

	fmt.Println("File path:", path)

	err = c.SaveUploadedFile(file, "public/"+path)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Error saving file:", err)
		data := gin.H{
			"is_uploaded": false,
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "upload ke direktori gagal", "message": data})
		return
	}

	_, err = govalidator.ValidateStruct(photoInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(photoInput.Title) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title length should not exceed 100 characters"})
		return
	}

	photoInput.UserID = userID
	photoInput.PhotoURL = path
	if err := database.DB.Create(&photoInput).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Photo uploaded successfully"})
}

func (pc *PhotoController) GetPhotos(c *gin.Context) {
	var photos []models.Photo
	if err := database.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch photos"})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func (pc *PhotoController) UpdatePhoto(c *gin.Context) {
	userID := helpers.ExtractUserID(c)
	photoID := c.Param("photoId")

	var updatedPhoto models.Photo
	if err := c.ShouldBind(&updatedPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if len(updatedPhoto.Title) > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title length should not exceed 100 characters"})
		return
	}

	photo := models.Photo{}
	if err := database.DB.Where("id = ? AND user_id = ?", photoID, userID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		errorMessages := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to upload user photo", "message": errorMessages})
		return
	}

	splitedFileName := strings.Split(file.Filename, ".")
	fileFormat := splitedFileName[len(splitedFileName)-1]
	path := fmt.Sprint("images/user/", userID, "_", time.Now().Format("010206150405"), ".", fileFormat)

	fmt.Println("File path:", path)

	err = c.SaveUploadedFile(file, "public/"+path)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Error saving file:", err)
		data := gin.H{
			"is_uploaded": false,
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "upload ke direktori gagal", "message": data})
		return
	}
	photo.UserID = userID
	photo.PhotoURL = path

	if err := database.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	userID := helpers.ExtractUserID(c)
	photoID := c.Param("photoId")

	photo := models.Photo{}
	if err := database.DB.Where("id = ? AND user_id = ?", photoID, userID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	if err := database.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
