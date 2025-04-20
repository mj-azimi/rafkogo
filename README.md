
# Rafkogo Framework

Rafkogo is a lightweight and modular web framework built with Go, designed to simplify the development of scalable web applications. It leverages popular libraries like **Gin** for routing and **GORM** for database operations, while providing a clean and organized structure inspired by frameworks like Laravel.

## Features
- **Modular Architecture**: Easily extend the framework by adding new modules (e.g., `user`, `product`) with a clear separation of concerns.
- **Environment Configuration**: Load settings from a `.env` file using the `godotenv` package.
- **Database Support**: Seamless integration with MySQL via GORM, with support for custom database configurations.
- **Template Rendering**: Render HTML templates using Go's `html/template` with Blade-like syntax (`.blade` files).
- **Code Generation**: Generate boilerplate code for new modules using the `make.go` script.
- **RESTful API Support**: Built-in support for creating APIs with Gin.

## Prerequisites
- **Go**: Version 1.18 or higher
- **MySQL**: Version 5.7 or higher
- **Git**: For cloning the repository

## Installation
1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/rafkogo.git
   cd rafkogo
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Set up the `.env` file**:
   Copy the `.env.example` file to `.env` and configure your environment variables:
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

4. **Create the database**:
   Create a MySQL database named `rafkogo` (or your preferred name) and update the `DB_DATABASE` variable in `.env`.

5. **Run the application**:
   ```bash
   go run main.go
   ```
   The server will start on the port specified in the `.env` file (default: `8080`).

## Usage
### Accessing the Application
- Open your browser and navigate to `http://localhost:8080` to see the welcome page.
- API endpoints are available under `/api` (e.g., `/api/users`).

### Generating a New Module
Use the `make.go` script to generate boilerplate code for a new module:
```bash
go run cmd/scripts/make.go <module_name>
```
Example:
```bash
go run cmd/scripts/make.go product
```
This will create the following files in `internal/product`:
- `handler.go`
- `model.go`
- `repository.go`
- `service.go`
- `middleware.go`
- `routes.go`

### Directory Structure
```
rafkogo/
├── bootstrap/        # Bootstrap application and database connection
├── cmd/              # CLI scripts (e.g., code generation)
├── config/           # Environment configuration
├── internal/         # Application modules (e.g., user)
│   └── user/
│       ├── view/     # Blade-like templates
│       ├── handler.go
│       ├── model.go
│       ├── routes.go
│       └── ...
├── routes/           # Route definitions
├── utils/            # Utility functions (e.g., template rendering)
├── .env              # Environment variables
├── go.mod            # Go module dependencies
└── main.go           # Application entry point
```

## Contributing
Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and commit (`git commit -m 'Add your feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a Pull Request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
For questions or feedback, feel free to reach out at [my LinkedIn](https://www.linkedin.com/in/mohammadjavadazimi).