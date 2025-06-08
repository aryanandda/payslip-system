package constants

const (
	// Auth errors
	ErrUnauthorized       = "Unauthorized"
	ErrInvalidToken       = "Invalid or expired token"
	ErrInvalidTokenClaims = "Invalid token claims"

	// Validation errors
	ErrInvalidRequest     = "Invalid request"
	ErrInvalidAmount      = "Amount must be greater than zero"
	ErrMissingDescription = "Description is required"

	// Attendance errors
	ErrWeekendAttendance         = "Cannot submit attendance on weekends"
	ErrAttendanceAlreadyExists   = "Attendance already submitted for today"

	// Overtime errors
	ErrOvertimeBefore5PM    = "Overtime can only be submitted after 5:00 PM"
	ErrOvertimeExceedsLimit = "Total overtime exceeds 3 hours for today"
	ErrOvertimeInvalidHours = "Overtime hours must be greater than zero"
	ErrOvertimeTooManyHours = "Overtime cannot exceed 3 hours per day"

	// Generic
	ErrInternalServer = "Internal server error"
)
