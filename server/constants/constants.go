package constants

const (

	// db table names
	ClassTableName                = "class"
	VisitorTableName              = "visitor"
	TimetableItemTableName        = "timetable_item"
	VisitorTimetableItemTableName = "visitor_timetable_item"
	MigrationsTableName           = "migrations"
	MigrationsFolder              = "migrations"

	// validation
	RequiredErrorMsg = "%s resource: '%s' is required"
	RequiredTag      = "required"
	MinValueErrorMsg = "%s resource: '%s' min value is %s"
	MinTag           = "min"
	MaxValueErrorMsg = "%s resource: '%s' max value is %s"
	MaxTag           = "max"
	Uuid4Tag         = "uuid4"
	Uuid4ErrorMsg    = "%s resource: id is not valid uuid4"
)
