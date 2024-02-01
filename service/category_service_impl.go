package service

import (
	"Hanivan/learn-golang-restfull-api/exception"
	"Hanivan/learn-golang-restfull-api/helper"
	"Hanivan/learn-golang-restfull-api/model/domain"
	"Hanivan/learn-golang-restfull-api/model/web"
	"Hanivan/learn-golang-restfull-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

// karena pake mysql(db transactional) maka pake gaya transactional

type CategoryServiceImpl struct {
	CategoryRepopsitory repository.CategoryRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepopsitory: repository,
		DB:                  DB,
		Validate:            validate}
}

func (service CategoryServiceImpl) Create(ctx context.Context, r web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: r.Name,
	}

	category = service.CategoryRepopsitory.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Update(ctx context.Context, r web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepopsitory.FindbyId(ctx, tx, r.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = r.Name

	category = service.CategoryRepopsitory.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepopsitory.FindbyId(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepopsitory.Delete(ctx, tx, category)

}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepopsitory.FindbyId(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepopsitory.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
