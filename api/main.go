package main 

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Print("api is on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {

}