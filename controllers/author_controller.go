package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-restful-api/config"
	"go-restful-api/helpers"
	"go-restful-api/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateAuthor(ctx *fiber.Ctx) error {
	var author models.Author
	var authorResponse models.AuthorResponse

	if err := ctx.BodyParser(&author); err != nil {
		return helpers.ResponseWithError(ctx, http.StatusUnprocessableEntity, err.Error())
	}

	err := config.DB.Where("name = ? OR email = ?", author.Name, author.Email).First(&author).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ResponseWithError(ctx, http.StatusInternalServerError, err.Error())
		}

	} else {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, "Author with the same name or email already exists")
	}

	err = config.DB.Create(&author).Error
	if err != nil {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, err.Error())
	}

	if err = config.DB.First(&author).First(&authorResponse).Error; err != nil {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, err.Error())
	}

	return helpers.ResponseWithSuccess(ctx, http.StatusCreated, "Success create author", authorResponse)

}

func GetAuthors(ctx *fiber.Ctx) error {
	authors := &[]models.Author{}

	if err := config.DB.Find(&authors).Error; err != nil {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, err.Error())
	}

	if len(*authors) == 0 {
		return helpers.ResponseWithError(ctx, http.StatusNotFound, "Author empty")
	}

	return helpers.ResponseWithSuccess(ctx, http.StatusOK, "Success get authors", authors)
}

func GetAuthorByID(ctx *fiber.Ctx) error {
	var author models.Author
	id := ctx.Params("id")
	if id == "" {
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, "Author ID empty")
	}

	err := config.DB.Where("id = ?", id).First(&author).Error
	if err != nil {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, err.Error())
	}

	return helpers.ResponseWithSuccess(ctx, http.StatusOK, "Success get author", author)

}

func UpdateAuthor(ctx *fiber.Ctx) error {
	var author models.Author
	var authorResponse models.Author

	idParams := ctx.Params("id")
	if idParams == "" {
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, "Book ID empty")
	}

	authorID, err := strconv.ParseUint(idParams, 10, 0)
	if err != nil {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, "Invalid Author ID")
	}
	uintAuthorID := uint(authorID)

	if err = ctx.BodyParser(&author); err != nil {
		return helpers.ResponseWithError(ctx, http.StatusUnprocessableEntity, "Request failed")
	}

	tx := config.DB.Begin()
	err = tx.Where("name = ? OR email = ?", author.Name, author.Email).First(&authorResponse).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, "Database error")
	} else if err == nil && (authorResponse.Name == author.Name || authorResponse.Email == author.Email) {
		tx.Rollback()
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, "Author with the same name or email already exists")
	}

	err = tx.Model(&models.Author{}).Where("id = ?", uintAuthorID).Updates(&author).Error
	if err != nil {
		tx.Rollback()
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, err.Error())
	}

	tx.Commit()

	err = config.DB.First(&authorResponse, uintAuthorID).Error
	if err != nil {
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	return helpers.ResponseWithSuccess(ctx, http.StatusOK, "Success update author", authorResponse)

}

func DeleteAuthor(ctx *fiber.Ctx) error {
	var author models.Author
	idParams := ctx.Params("id")
	if idParams == "" {
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, "ID does no exists")
	}

	err := config.DB.Where("id = ?", idParams).First(&author).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return helpers.ResponseWithError(ctx, http.StatusNotFound, "Author not found")
		}

		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	err = config.DB.Delete(&author, idParams).Error
	if err != nil {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, err.Error())
	}

	return helpers.ResponseWithSuccess(ctx, http.StatusOK, "Success delete author", nil)
}
