package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-restful-api/config"
	"go-restful-api/helpers"
	"go-restful-api/models"
	"gorm.io/gorm"
	"net/http"
)

func CreateBook(ctx *fiber.Ctx) error {
	var book models.Book
	var bookResponse models.BookResponse

	if err := ctx.BodyParser(&book); err != nil {
		return helpers.ResponseWithError(ctx, http.StatusUnprocessableEntity, "Request failed")
	}

	var author models.Author
	if err := config.DB.Find(&author, book.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ResponseWithError(ctx, http.StatusNotFound, "Author not found")
		}
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Where("title = ?", book.Title).Find(&book).Error; err != nil {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, err.Error())
	}

	err := config.DB.Create(&book).Error
	if err != nil {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, err.Error())
	}

	err = config.DB.Joins("Author").First(&book).First(&bookResponse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ResponseWithError(ctx, http.StatusNotFound, "Book not found")
		}
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	return helpers.ResponseWithSuccess(ctx, http.StatusCreated, "Success create book", bookResponse)

}

func GetBooks(ctx *fiber.Ctx) error {
	var books []models.Book
	var bookResponse []models.BookResponse

	if err := config.DB.Joins("Author").Find(&books).Find(&bookResponse).Error; err != nil {
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	return helpers.ResponseWithSuccess(ctx, http.StatusOK, "List's Book", bookResponse)
}

func GetBookByID(ctx *fiber.Ctx) error {
	var book models.Book
	var bookResponse models.BookResponse

	id := ctx.Params("id")
	if id == "" {
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, "Book ID empty")
	}

	if err := config.DB.Joins("Author").First(&book, id).First(&bookResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ResponseWithError(ctx, http.StatusNotFound, "Book not found")
		}
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, err.Error())

	}

	return helpers.ResponseWithSuccess(ctx, http.StatusOK, "Detail Book", bookResponse)
}

func UpdateBook(ctx *fiber.Ctx) error {
	var book models.Book
	var bookResponse models.BookResponse

	idParams := ctx.Params("id")
	if idParams == "" {
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, "Book ID empty")
	}

	if err := ctx.BodyParser(&book); err != nil {
		return helpers.ResponseWithError(ctx, http.StatusUnprocessableEntity, "Request failed")
	}

	err := config.DB.Model(&models.Book{}).Where("id = ?", idParams).Updates(book).Error
	if err != nil {
		return helpers.ResponseWithError(ctx, http.StatusBadRequest, err.Error())
	}

	err = config.DB.Joins("Author").First(&book, idParams).First(&bookResponse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ResponseWithError(ctx, http.StatusNotFound, "Book not found")
		}
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	return helpers.ResponseWithSuccess(ctx, http.StatusOK, "Success update author", bookResponse)

}

func DeleteBook(ctx *fiber.Ctx) error {
	book := &models.Book{}
	idParams := ctx.Params("id")
	if idParams == "" {
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, "Book ID empty")
	}

	if err := config.DB.Delete(&book, idParams).Error; err != nil {
		return helpers.ResponseWithError(ctx, http.StatusInternalServerError, err.Error())
	}

	return helpers.ResponseWithSuccess(ctx, http.StatusOK, "Success delete book", nil)
}
