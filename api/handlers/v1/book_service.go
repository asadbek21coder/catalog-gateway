package v1

import (
	"net/http"
	"strconv"

	"github.com/asadbek21coder/catalog/gateway/genproto/book_service"
	"github.com/gin-gonic/gin"
)

// Book_service godoc
// @ID get-books
// @Router /v1/books [GET]
// @Summary get books
// @Description get books
// @Tags book service
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=book_service.Books} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAll(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}
	// fmt.Println(c.Query("search"))
	resp, err := h.services.Service().GetAll(
		c.Request.Context(),
		&book_service.GetAllRequest{
			Search: c.Query("search"),
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all books", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Book_service godoc
// @ID update-book
// @Router /v1/books [PUT]
// @Summary update book
// @Description update book
// @Tags book service
// @Accept json
// @Produce json
// @Param post body book_service.Book true "post"
// @Success 200 {object} models.ResponseModel{data=book_service.Book} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Update(c *gin.Context) {
	var books book_service.Book

	if err := c.BindJSON(&books); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Service().Update(c.Request.Context(), &books)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating book", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Book_service godoc
// @ID delete-book
// @Router /v1/books/{id} [DELETE]
// @Summary delete book
// @Description delete book
// @Tags book service
// @Accept json
// @Produce json
// @Param id path int32 true "id"
// @Success 200 {object} models.ResponseModel{data=book_service.Id} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while getting id", err)
		return
	}
	id1 := int32(id)

	resp, err := h.services.Service().Delete(
		c.Request.Context(),

		&book_service.Id{
			Id: id1,
		},
	)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while deleting book", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Book_service godoc
// @ID getbyid-book
// @Router /v1/books/{id} [GET]
// @Summary getById book
// @Description GetById book
// @Tags book service
// @Accept json
// @Produce json
// @Param id path int32 true "id"
// @Success 200 {object} models.ResponseModel{data=book_service.Book} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while getting id", err)
		return
	}
	id1 := int32(id)

	resp, err := h.services.Service().GetById(
		c.Request.Context(),

		&book_service.Id{
			Id: id1,
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting book by id", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// Book_service godoc
// @ID create-book
// @Router /v1/books [POST]
// @Summary create book
// @Description Create book
// @Tags book service
// @Accept json
// @Produce json
// @Param post body book_service.Book true "post"
// @Success 200 {object} models.ResponseModel{data=book_service.Book} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) Create(c *gin.Context) {

	var book book_service.Book
	if err := c.BindJSON(&book); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Service().Create(
		c.Request.Context(),
		&book,
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error creating book", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}
