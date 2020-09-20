package constants

import "time"

const (

	// struct names

	ClassStructName                = "Class"
	VisitorStructName              = "Visitor"
	TimetableItemStructName        = "TimetableItem"
	VisitorTimetableItemStructName = "VisitorTimetableItem"

	// db table names

	ClassTableName                = "class"
	VisitorTableName              = "visitor"
	TimetableItemTableName        = "timetable_item"
	VisitorTimetableItemTableName = "visitor_timetable_item"
	MigrationsTableName           = "migrations"

	// migrations

	MigrationsFolder = "migrations"

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
	Uuid4ErrorMsg    = "%s resource: '%s' is not valid uuid4"
	EmailErrorMsg    = "%s resource: email is not valid"

	VisitorEmailExistsErrorMsg        = "%s resource: email is already registered"
	TimetableItemDateExistsErrorMsg   = "%s resource: date is already registered"
	BookingAlreadyExistsErrorMsg      = "%s resource: booking already exists"
	ClassDoesNotExistErrorMsg         = "%s resource: class does not exists"
	VisitorDoesNotExistErrorMsg       = "%s resource: visitor does not exists"
	TimetableItemDoesNotExistErrorMsg = "%s resource: visitor does not exists"

	// date

	DateFormat = "2006-01-02"

	// Server

	DefaultWriteTimout  = 60 * time.Second
	DefaultReadTimeout  = 60 * time.Second
	DefaultStoreTimeout = 60 * time.Second
)
