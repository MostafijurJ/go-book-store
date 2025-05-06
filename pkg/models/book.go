package models

import (
	"time"
)

// Static slice to store library in memory
var library []Book
var nextId = 1

type BookCreateRequest struct {
	Title      string `json:"title" validate:"required" errorMessage:"Title is required"`
	Author     string `json:"author" validate:"required" errorMessage:"Author is required"`
	Genre      string `json:"genre" validate:"required" errorMessage:"Genre is required"`
	Publisher  string `json:"publisher" validate:"required" errorMessage:"Publisher is required"`
	Edition    string `json:"edition" validate:"required" errorMessage:"Edition is required"`
	TotalPages int    `json:"total_pages" validate:"required" errorMessage:"Total pages is required"`
}

type BookResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Genre       string    `json:"genre"`
	Publisher   string    `json:"publisher"`
	Edition     string    `json:"edition"`
	PublishDate time.Time `json:"publish_date"`
	TotalPages  int       `json:"total_pages"`
	CreatedAt   time.Time `json:"created_at"`
}

// Book struct represents a book in the library
type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Genre       string    `json:"genre"`
	Publisher   string    `json:"publisher"`
	Edition     string    `json:"edition"`
	PublishDate time.Time `json:"publish_date"`
	TotalPages  int       `json:"total_pages"`
	CreatedAt   time.Time `json:"created_at"`
}

type PaginatedBooksResponse struct {
	Books []BookResponse `json:"books"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

func GetFilteredAndPaginatedBooks(author, genre string, page, limit int) PaginatedBooksResponse {
	var filteredBooks []Book

	// Filter books by author and genre
	for _, book := range library {
		if (author == "" || book.Author == author) && (genre == "" || book.Genre == genre) {
			filteredBooks = append(filteredBooks, book)
		}
	}

	// Apply pagination
	start := (page - 1) * limit
	end := start + limit
	if start > len(filteredBooks) {
		return PaginatedBooksResponse{
			Books: []BookResponse{},
			Total: 0,
			Page:  page,
			Limit: limit,
		}
	}
	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	// Map to BookResponse
	var paginatedBooks []BookResponse
	for _, book := range filteredBooks[start:end] {
		paginatedBooks = append(paginatedBooks, mapToBookResponse(book))
	}

	return PaginatedBooksResponse{
		Books: paginatedBooks,
		Total: len(filteredBooks),
		Page:  page,
		Limit: limit,
	}
}

func GetBookById(bookId int) BookResponse {
	for _, book := range library {
		if book.ID == bookId {
			return mapToBookResponse(book)
		}
	}
	return BookResponse{}
}

func CreateBook(book BookCreateRequest) BookResponse {
	bookToStore := Book{
		ID:          nextId,
		Title:       book.Title,
		Author:      book.Author,
		Genre:       book.Genre,
		Publisher:   book.Publisher,
		Edition:     book.Edition,
		PublishDate: time.Now(),
		TotalPages:  book.TotalPages,
		CreatedAt:   time.Now(),
	}

	library = append(library, bookToStore)
	nextId++

	return mapToBookResponse(bookToStore)
}

func UpdateBook(bookId int, book BookCreateRequest) BookResponse {
	for i, b := range library {
		if b.ID == bookId {
			library[i].Title = book.Title
			library[i].Author = book.Author
			library[i].Genre = book.Genre
			library[i].Publisher = book.Publisher
			library[i].Edition = book.Edition
			library[i].TotalPages = book.TotalPages
			library[i].PublishDate = time.Now()
			return mapToBookResponse(library[i])
		}
	}
	return BookResponse{}
}

func mapToBookResponse(bookToStore Book) BookResponse {
	return BookResponse{
		ID:          bookToStore.ID,
		Title:       bookToStore.Title,
		Author:      bookToStore.Author,
		Genre:       bookToStore.Genre,
		Publisher:   bookToStore.Publisher,
		Edition:     bookToStore.Edition,
		PublishDate: bookToStore.PublishDate,
		TotalPages:  bookToStore.TotalPages,
		CreatedAt:   bookToStore.CreatedAt,
	}
}
