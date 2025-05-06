package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mostafijurj/go-book-store/pkg/models"
	"net/http"
	"strconv"
)

// CreateBook godoc
// @Summary      Create a new book
// @Description  Adds a new book to the store
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book body models.BookCreateRequest true "Book data"
// @Success      201  {object} models.BookResponse
// @Failure      400  {string} string "Invalid input"
// @Router       /book/create [post]
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var createRequest models.BookCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdBook := models.CreateBook(createRequest)

	setHeaderValues(w)

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(createdBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetBookById godoc
// @Summary Get book by ID
// @Description Retrieve a single book by its ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.BookResponse
// @Failure 404 {string} string "Book not found"
// @Router /book/{id} [get]
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book := models.GetBookById(id)

	if book == (models.BookResponse{}) {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	setHeaderValues(w)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetAllBook godoc
// @Summary Get all books
// @Description Retrieve a list of all books in the store with optional filtering and pagination
// @Tags books
// @Produce json
// @Param page query int false "Page number for pagination" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param author query string false "Filter by author"
// @Param genre query string false "Filter by genre"
// @Success 200 {array} models.PaginatedBooksResponse
// @Failure 404 {string} string "No books found"
// @Router /book [get]
func GetAllBook(writer http.ResponseWriter, r *http.Request) {
	// Parse query parameters for pagination and filtering
	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	limit, _ := strconv.Atoi(query.Get("limit"))
	author := query.Get("author")
	genre := query.Get("genre")

	// Set default values for pagination if not provided
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	// Fetch filtered and paginated books
	bookList := models.GetFilteredAndPaginatedBooks(author, genre, page, limit)

	// Marshal the response
	response, err := json.Marshal(bookList)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	setHeaderValues(writer)
	_, err = writer.Write(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update an existing book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.BookCreateRequest true "Updated book data"
// @Success 200 {object} models.Book
// @Failure 400 {string} string "Invalid input"
// @Failure 404 {string} string "Book not found"
// @Router /book/update/{id} [put]
// UpdateBook updates an existing book by its ID
func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(writer, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updateRequest models.BookCreateRequest
	if err := json.NewDecoder(request.Body).Decode(&updateRequest); err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedBook := models.UpdateBook(id, updateRequest)

	if updatedBook == (models.BookResponse{}) {
		http.Error(writer, "Book not found", http.StatusNotFound)
		return
	}

	setHeaderValues(writer)
	err = json.NewEncoder(writer).Encode(updatedBook)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func setHeaderValues(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
