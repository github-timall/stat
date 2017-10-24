package middleware

import "net/http"

type MwType func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, mws ...MwType) http.HandlerFunc {
	for _, m := range mws {
		f = m(f)
	}
	return f
}

//func Method(m string) Middleware {
//
//	// Create a new Middleware
//	return func(f http.HandlerFunc) http.HandlerFunc {
//
//		// Define the http.HandlerFunc
//		return func(w http.ResponseWriter, r *http.Request) {
//
//			// Do middleware things
//			if r.Method != m {
//				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
//				return
//			}
//
//			// Call the next middleware/handler in chain
//			f(w, r)
//		}
//	}
//}
//func Hello(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "hello world")
//}
//func main() {
//	http.HandleFunc("/", Chain(Hello, Method("GET")))
//	http.ListenAndServe(":8080", nil)
//}
