package constants

const (
	// Error codes
	NotFound            = "RP-404"
	BadRequest          = "RP-400"
	Unauthorized        = "RP-401"
	Forbidden           = "RP-403"
	InternalServerError = "RP-500"

	// Messages
	EntityNotFound        = "%s with id %d does not exist."
	GetEntityByIdMessage  = "Found %s with id %d."
	SaveEntityError       = "Error while saving %s."
	SuccessMessage        = "You have successfully %s!"
	GetAllEntitiesError   = "Error while fetching all %s."
	GetAllEntitiesMessage = "Found %d %s."
	CreateEntityError     = "Error while creating %s."
	CreateEntityMessage   = "Created %s successfully."
	UpdateEntityError     = "Error while updating %s."
	UpdateEntityMessage   = "Updated %s successfully."

	// Queries
	FindByIdQuery = "id = ?"

	// Misc
	TimeFormat      = "2006-01-02 15:04:05"
	TraceIdHeader   = "x-trace-id"
	Locale          = "en"
	LocalEnv        = "local"
	DevelopmentEnv  = "development"
	ProductionEnv   = "prod"
	RequestIdCtxKey = "request_id"
)
