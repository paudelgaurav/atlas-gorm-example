package constants

// Upon updating or deleting any permission you might need to run a patch in the database
// however, there is not problem with adding new permissions
func GetAllPermissions() []string {
	return []string{
		"can_list_users",
		"can_create_users",
		"can_update_users",
		"can_ping",
	}
}
