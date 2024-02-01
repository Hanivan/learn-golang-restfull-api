package controller

import (
	"Hanivan/learn-golang-restfull-api/helper"
	"Hanivan/learn-golang-restfull-api/model/web"
	"Hanivan/learn-golang-restfull-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}

func (c CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromequestBody(r, &categoryCreateRequest)

	categoryRespone := c.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRespone,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromequestBody(r, &categoryUpdateRequest)

	categoryId, err := strconv.Atoi(p.ByName("categoryId"))
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = categoryId

	categoryRespone := c.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRespone,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId, err := strconv.Atoi(p.ByName("categoryId"))
	helper.PanicIfError(err)

	c.CategoryService.Delete(r.Context(), categoryId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId, err := strconv.Atoi(p.ByName("categoryId"))
	helper.PanicIfError(err)

	categoryResponse := c.CategoryService.FindById(r.Context(), categoryId)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryResponses := c.CategoryService.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
