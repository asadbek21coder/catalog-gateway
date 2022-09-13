package v1

import (
	"net/http"
	"strconv"

	"github.com/asadbek21coder/catalog/gateway/genproto/book_service"
	"github.com/gin-gonic/gin"
)

// Book_service godoc
// @ID get-categories
// @Router /v1/categories [GET]
// @Summary get categories
// @Description get categories
// @Tags category
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=book_service.Category} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllCategories(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}
	// fmt.Println(c.Query("search"))
	resp, err := h.services.Service().GetAllCategories(
		c.Request.Context(),
		&book_service.GetAllRequest{
			Search: c.Query("search"),
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all categories", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Book_service godoc
// @ID update-category
// @Router /v1/categories [PUT]
// @Summary update category
// @Description update category
// @Tags category
// @Accept json
// @Produce json
// @Param book body book_service.Category true "post"
// @Success 200 {object} models.ResponseModel{data=book_service.Category} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var categories book_service.Category

	if err := c.BindJSON(&categories); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Service().UpdateCategory(c.Request.Context(), &categories)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while updating category", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Book_service godoc
// @ID delete-category
// @Router /v1/categories/{id} [DELETE]
// @Summary delete category
// @Description delete category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int32 true "id"
// @Success 200 {object} models.ResponseModel{data=book_service.Id} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while getting id", err)
		return
	}
	id1 := int32(id)

	resp, err := h.services.Service().DeleteCategory(
		c.Request.Context(),

		&book_service.Id{
			Id: id1,
		},
	)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while deleting category", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Book_service godoc
// @ID getbyid-category
// @Router /v1/categories/{id} [GET]
// @Summary getById category
// @Description GetById category
// @Tags category
// @Accept json
// @Produce json
// @Param id path int32 true "id"
// @Success 200 {object} models.ResponseModel{data=book_service.Category} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while getting id", err)
		return
	}
	id1 := int32(id)

	resp, err := h.services.Service().GetCategoryById(
		c.Request.Context(),

		&book_service.Id{
			Id: id1,
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting category by id", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// Book_service godoc
// @ID create-category
// @Router /v1/categories [POST]
// @Summary create category
// @Description Create category
// @Tags category
// @Accept json
// @Produce json
// @Param post body book_service.Category true "post"
// @Success 200 {object} models.ResponseModel{data=book_service.Category} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateCategory(c *gin.Context) {

	var category book_service.Category
	if err := c.BindJSON(&category); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Service().CreateCategory(
		c.Request.Context(),
		&category,
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error creating category", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}
