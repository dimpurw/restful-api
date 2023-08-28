package repository

import (
	"context"
	"restful-api/model"
)

type BookRepository interface {
	Save(ctx context.Context, book model.Book)
	Update(ctx context.Context, book model.Book)
	Delete(ctx context.Context, bookId int)
	findById(ctx context.Context, bookId int)(model.Book, error)
	findAll(ctx context.Context) []model.Book
}