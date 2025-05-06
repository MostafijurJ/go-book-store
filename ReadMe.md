# 📚 Book Store API

A RESTful API built using **Go** and **Gorilla Mux**, designed to manage books in a store. It demonstrates a simple CRUD
application with clean, modular architecture, supporting pagination, filtering, and error handling. Also consider
Swagger for API documentation.

---

## 🛠️ Setup Instructions

1. **Initialize Go Module**
   ```bash
   go mod init github.com/mostafijurj/go-book-store
   ```

2. **Install Gorilla Mux**
   ```bash
   go get -u github.com/gorilla/mux
   ```

3. **Install Swagger CLI**
   ```bash
   go get -u github.com/swaggo/swag/cmd/swag
   ```

4. **Generate Swagger Docs**
   ```bash
   swag init -g main.go
   ```

5. **Run the Application**
   ```bash
   go run main.go
   ```

---

## 📘 API Endpoints

### 1. ➕ Create a Book

- **Endpoint:** `POST /api/v1/books`
- **Description:** Creates a new book in the store.

#### 🔸 Request Body

```json
{
  "title": "Book Title",
  "author": "Author Name",
  "genre": "Fiction",
  "publisher": "Publisher Name",
  "edition": "First",
  "total_pages": 300
}
```

#### 🔹 Response

```json
{
  "id": 1,
  "title": "Book Title",
  "author": "Author Name",
  "genre": "Fiction",
  "publisher": "Publisher Name",
  "edition": "First",
  "publish_date": "2023-10-01T00:00:00Z",
  "total_pages": 300,
  "created_at": "2023-10-01T00:00:00Z"
}
```

---

### 2. 🔍 Get a Book by ID

- **Endpoint:** `GET /api/v1/books/{id}`
- **Description:** Retrieves a book by its ID.

#### 🔹 Response

```json
{
  "id": 1,
  "title": "Book Title",
  "author": "Author Name",
  "genre": "Fiction",
  "publisher": "Publisher Name",
  "edition": "First",
  "publish_date": "2023-10-01T00:00:00Z",
  "total_pages": 300,
  "created_at": "2023-10-01T00:00:00Z"
}
```

---

### 3. ✏️ Update a Book

- **Endpoint:** `PUT /api/v1/books/{id}`
- **Description:** Updates an existing book by its ID.

#### 🔸 Request Body

```json
{
  "title": "Updated Book Title",
  "author": "Updated Author Name",
  "genre": "Non-Fiction",
  "publisher": "Updated Publisher Name",
  "edition": "Second",
  "total_pages": 350
}
```

#### 🔹 Response

```json
{
  "id": 1,
  "title": "Updated Book Title",
  "author": "Updated Author Name",
  "genre": "Non-Fiction",
  "publisher": "Updated Publisher Name",
  "edition": "Second",
  "publish_date": "2023-10-02T00:00:00Z",
  "total_pages": 350,
  "created_at": "2023-10-01T00:00:00Z"
}
```

---

### 4. 📄 Get All Books

- **Endpoint:** `GET /api/v1/books`
- **Description:** Retrieves all books with optional filters and pagination.

#### 🔸 Query Parameters

| Parameter | Type   | Description                            |
|-----------|--------|----------------------------------------|
| `page`    | int    | Page number (default: 1)               |
| `limit`   | int    | Number of books per page (default: 10) |
| `author`  | string | Filter by author (optional)            |
| `genre`   | string | Filter by genre (optional)             |

#### 🔹 Response

```json
{
  "books": [
    {
      "id": 1,
      "title": "Updated Book Title",
      "author": "Updated Author Name",
      "genre": "Non-Fiction",
      "publisher": "Updated Publisher Name",
      "edition": "Second",
      "publish_date": "2023-10-02T00:00:00Z",
      "total_pages": 350,
      "created_at": "2023-10-01T00:00:00Z"
    },
    {
      "id": 2,
      "title": "Updated Book Title",
      "author": "Updated Author Name",
      "genre": "Non-Fiction",
      "publisher": "Updated Publisher Name",
      "edition": "Second",
      "publish_date": "2023-10-02T00:00:00Z",
      "total_pages": 350,
      "created_at": "2023-10-01T00:00:00Z"
    }
  ],
  "total": 2,
  "page": 1,
  "limit": 10
}
```

---
