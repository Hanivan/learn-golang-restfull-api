package repository

import (
	"context"
	"database/sql"

	"Hanivan/learn-golang-restfull-api/model/domain"
)

// biasakan membuat contract dulu dalam interface
// baru implementasi nya dalam bentuk struct

// kalau pake db relational, usahakan pake fungsi yang bisa transactional

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindbyId(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
