package main

//go:generate gourd gen main $GOFILE

// port to use
var port string

func init() {
	port = "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
}

// main
//
//gourd:db upperio
//gourd:router gorillamux
//gourd:middleware negroni
//gourd:auth /oauth2 osin
//gourd:rest /api/posts Post
//gourd:rest /api/comments Comment
//
func main() {
	// use the gourd generated main (gourdMain)
	// alternatively, user may copy the content of gourdMain
	// here then modify
	http.HandleFunc("/", MainHandler())
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
