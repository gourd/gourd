package main

//go:generate gourd gen main $GOFILE

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
	gourdMain()
}
