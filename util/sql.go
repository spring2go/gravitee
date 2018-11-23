package util

import (
	"database/sql"
)

// IntOrNull returns properly configured sql.NullInt64
func IntOrNull(n int64) sql.NullInt64 {
	return sql.NullInt64{Int64: n, Valid: true}
}

// PositiveIntOrNull returns proprely configured sql.NullInt64 for a positive number
func PositiveIntOrNull(n int64) sql.NullInt64 {
	if n <= 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: n, Valid: true}
}

// FloatOrNull returns properly configured sql.NullFloat64
func FloatOrNull(n float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: n, Valid: true}
}

// PositiveFloatOrNull returns properly configured sql.NullFloat64 for a positive number
func PositiveFloatOrNull(n float64) sql.NullFloat64 {
	if n <= 0 {
		return sql.NullFloat64{Float64: 0.0, Valid: false}
	}
	return sql.NullFloat64{Float64: n, Valid: true}
}

// StringOrNull returns properly configured sql.NullString
func StringOrNull(str string) sql.NullString {
	if str == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: str, Valid: true}
}
