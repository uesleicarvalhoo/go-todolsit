package utils

var (
	// Database
	ErrDatabaseConnection = "Error on connect to database"
	ErrRunMigrations      = "Error on run database migrations"

	// Auth
	ErrInvalidCredentials = "Invalid Credentials"
	ErrGenerateToken      = "Error on generate Access Token"
	ErrAuthHeaderNotFound = "Header 'X-Authorization' not found"

	// User
	ErrEmailNotFound   = "Email not found"
	ErrPasswordInvalid = "Password not match"
	ErrEmailDuplicated = "Current email is already in use"
)
