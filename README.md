# Rafkogo Framework

Rafkogo is a lightweight, modular web framework built with Go, designed to streamline the development of scalable web applications. It integrates **Gin** for routing and **GORM** for database operations, offering a structured and extensible architecture inspired by frameworks like Laravel.

## Features

- **Modular Architecture**: Organize code into modules (e.g., `user`, `chat`) with clear separation of concerns (handlers, services, repositories, models, routes, middleware).
- **Environment Configuration**: Load settings from a `.env` file using `godotenv` for flexible configuration management.
- **Database Integration**: Seamless MySQL support via GORM, with automatic migrations and CRUD operations for models like `User` and `Chat`.
- **Template Rendering**: Render HTML templates with Blade-like syntax (`.blade` files) using Go's `html/template` package.
- **Code Generation**: Generate boilerplate code for new modules using the `make.go` script to accelerate development.
- **RESTful API Support**: Create APIs with Gin, including grouped routes (e.g., `/api/users`) for modular endpoints.
- **Middleware**: Custom middleware (e.g., `LoggerMiddleware`) for request logging and processing.
- **Static Assets**: Serve CSS, images, and other assets from the `public` directory.
- **Chat Module**: Manage chat messages with database-backed CRUD operations.
- **User Management**: Basic user model with username and password support.
- **Responsive UI**: Modern, flexbox-based styling via `app.css` for a clean user interface.

## Prerequisites

- **Go**: Version 1.18 or higher
- **MySQL**: Version 5.7 or higher
- **Git**: For cloning the repository

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/mj-azimi/rafkogo.git
   cd rafkogo
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set up the** `.env` **file**: Copy `.env.example` to `.env` and configure your environment variables:

   ```env
   APP_NAME=rafkogo
   APP_ENV=local
   APP_KEY=base64:wFo4/nIywKykJGtD3G5PcvxshpgMC5sQ9cWwPvKHUUA=
   APP_DEBUG=true
   APP_URL=http://localhost:8000
   APP_PORT=8080
   
   DB_CONNECTION=mysql
   DB_HOST=localhost
   DB_PORT=3306
   DB_DATABASE=rafkogo
   DB_USERNAME=root
   DB_PASSWORD=""
   ```

4. **Create the database**: Create a MySQL database named `rafkogo` (or your preferred name) and update the `DB_DATABASE` variable in `.env`.

5. **Run the application**:

   ```bash
   go run main.go
   ```

   The server will start on the port specified in `.env` (default: 8080).

## Usage

### Accessing the Application

- Navigate to `http://localhost:8080` in your browser to view the welcome page.
- Access API endpoints under `/api` (e.g., `/api/users` for user-related routes).
- Static assets (CSS, images) are served from `/public`.

### Generating a New Module

Use the `make.go` script to generate boilerplate code for a new module:

```bash
go run cmd/scripts/make.go <module_name>
```

Example:

```bash
go run cmd/scripts/make.go product
```

This creates the following files in `internal/product`:

- `handler.go`
- `model.go`
- `repository.go`
- `service.go`
- `middleware.go`
- `routes.go`

### Directory Structure

```
rafkogo/
├── bootstrap/               # Application and database initialization
├── cmd/                     # CLI scripts (e.g., code generation)
├── config/                  # Environment configuration
├── internal/                # Application modules
│   └── user/                # User module
│       ├── database/        # Database models (e.g., user, chat)
│       ├── view/            # Blade-like templates
│       ├── handler.go       # Request handlers
│       ├── middleware.go    # Custom middleware
│       ├── routes.go        # Route definitions
│       └── ...
├── migration_register/      # Database migration registration
├── public/                  # Static assets (CSS, images)
├── routes/                  # API and web route definitions
├── utils/                   # Utility functions (e.g., template rendering)
├── storage/                 # Logs and other storage
├── .env                     # Environment variables
├── go.mod                   # Go module dependencies
└── main.go                  # Application entry point
```

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and commit (`git commit -m 'Add your feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

For questions or feedback, feel free to reach out at [my LinkedIn](https://www.linkedin.com/in/mohammadjavadazimi).
