package constants

const (

	// struct names

	ClassStructName   = "Class"
	VisitorStructName = "Visitor"

	// db table names

	ClassTableName                = "class"
	VisitorTableName              = "visitor"
	TimetableItemTableName        = "timetable_item"
	VisitorTimetableItemTableName = "visitor_timetable_item"
	MigrationsTableName           = "migrations"
	MigrationsFolder              = "migrations"

	// validation tags

	RequiredTag = "required"
	MinTag      = "min"
	MaxTag      = "max"
	Uuid4Tag    = "uuid4"
	EmailTag    = "email"

	// validation error messages

	RequiredErrorMsg = "%s resource: '%s' is required"
	MinValueErrorMsg = "%s resource: '%s' min value is %s"
	MaxValueErrorMsg = "%s resource: '%s' max value is %s"
	Uuid4ErrorMsg    = "%s resource: id is not valid uuid4"
	EmailErrorMsg    = "%s resource: email is not valid"
)
