package app

import "net/http"

func star() {

	//define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//starting servers
	http.ListenAndServe("localhost:8000", nil)

}
