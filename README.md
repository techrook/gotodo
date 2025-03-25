# Go Todo App

A simple CRUD (Create, Read, Update, Delete) To-Do application built with Golang.
## Features

- ✅ Add new tasks
- ✔️ Mark tasks as completed(WIP)
- 🗑️ Delete tasks(WIP)
- 📋 View all tasks

## Prerequisites

- Go (version 1.16+)
- MySQL
- Git

## Tech Stack

- Backend: Go (Golang)
- Router: Gorilla Mux
- Database: MySQL
- Frontend: HTML, Bootstrap

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/gotodo.git
   cd gotodo
   ```

2. Set up environment variables:
   Create a `.env` file in the project root with:
   ```
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_USER=your_username
   DB_PASS=your_password
   ```

3. Install dependencies:
   ```sh
   go mod tidy
   ```

4. Set up MySQL database:
   - Ensure MySQL is installed and running
   - Create a database named `gotodo`

5. Run the application:
   ```sh
   go run main.go
   ```

## Project Structure

```
gotodo/
│
├── main.go           # Entry point
├── config/           # Database configuration
├── controllers/      # Request handlers
├── models/           # Data models and database interactions
├── routes/           # Application routes
└── templates/        # HTML templates
```

## Routes

- `GET /`: View all tasks
- `POST /add`: Add a new task
- `PATCH /complete/{id}`: Mark a task as completed
- `DELETE /delete/{id}`: Delete a task

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Credits

Inspired by [ichtrojan/go-todo](https://github.com/ichtrojan/go-todo)


## Contact

Itohowo Monday Umoh - [https://x.com/Itohowo23]

Project Link: [https://github.com/techrook/gotodo](https://github.com/techrook/gotodo)