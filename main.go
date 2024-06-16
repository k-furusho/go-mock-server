func main()  {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}