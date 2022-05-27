package utils

import "errors"

const (
	AppName = "orders"
	Port    = "port"
)

// API  error codes
const (
	CodeInvalidRequestBody = "INVALID_REQUEST_BODY"
	CodeFailedSaveOrder    = "FAILED_SAVE_ORDER"
)

var (
	ErrInsertFailed = errors.New("failed to insert record to database")
)

// database config names
const (
	DbUser           = "DB_USER"
	DbPwd            = "DB_PWD"
	DbName           = "DB_NAME"
	DbPort           = "DB_PORT"
	DbHost           = "DB_HOST"
	DbHostedOnCloud  = "DB_CLOUD"
	DbConnectionName = "DB_INSTANCE_CONNECTION_NAME"
	DbTimeZone       = "DB_TIMEZONE"
)
