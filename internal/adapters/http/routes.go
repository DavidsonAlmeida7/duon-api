package http

import "net/http"

const (
	GetBarRouteConst = "/"

	GetUserRoutesConst = "/users"

	// === Route marker ===
)

const (
	StatusOk                 = http.StatusOK
	CreatedConst             = http.StatusCreated
	BadRequestConst          = http.StatusBadRequest
	Unauthorized             = http.StatusUnauthorized
	ForbiddenRequestConst    = http.StatusForbidden
	NotFoundConst            = http.StatusNotFound
	TimeoutConst             = http.StatusRequestTimeout
	InternalServerErrorConst = http.StatusInternalServerError
	NoContentConst           = http.StatusNoContent
)
