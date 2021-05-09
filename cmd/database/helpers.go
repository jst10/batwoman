package database
func extendConditions(conditions string, extension string) string {
	if len(conditions) > 0 {
		conditions = conditions + " and "
	}
	conditions = conditions + extension
	return conditions
}