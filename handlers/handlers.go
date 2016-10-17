package handlers

import "database/sql"

var db *sql.DB

func init() {
	// var w http.ResponseWriter

	// if openDatabaseConnection() != nil {
	// 	connection := hasDatabaseConnection()
	// 	if connection == false {
	// 		fmt.Println("Error - Database Server Down")
	// 		http.Error(w, http.StatusText(504), 504)
	// 	} else {
	// 		fmt.Println("Error - Can't connect to database")
	// 		http.Error(w, http.StatusText(503), 503)
	// 	}
	// }
}

//openDatabaseConnection opens a database connection to the persistent storage
// func openDatabaseConnection() error {
// 	var err error
//
// 	db, err = sql.Open("mysql", "root:root@tcp("+os.Getenv("DOCKER_HOST_IP")+":3307)/ecofox-insights-user-service")
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	fmt.Println("Database connection successfully created")
// 	return nil
// }
//
// //hasDatabaseConnection checks if the database is up and runnig
// func hasDatabaseConnection() bool {
// 	err := db.Ping()
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }
