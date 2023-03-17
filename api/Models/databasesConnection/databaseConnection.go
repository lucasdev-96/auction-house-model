package models

func GetDataBase(db string) any {
	var database interface{}
	switch db {
	case "mysql":
		database = "mysql"
	case "sqlite3":
		database = "sqlite3"
	default:
		database = ConnectPostegreSQL()
	}
	return database
}
