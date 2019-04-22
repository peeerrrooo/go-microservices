package errors_handler

var Handler map[int]string = map[int]string{
	// Common errors.
	10: "Internal error",
	11: "Error in JSON parse",
	12: "Error in JSON ARRAY parse",
	13: "Error in JSON encode",
	14: "Error in JSON ARRAY encode",
	15: "Unknown URI",
	16: "Permissions denied",

	// Client service errors.
	100: "URI not found",
	/*
		REST POST ERRORS
	*/
	200: "Miss post id",
}
