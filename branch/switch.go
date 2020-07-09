package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "HEAD"}

	// golang's switch does not auto-fallthrough, rather auto-break.
	// `fallthrough` keyword is needed when fallthrough is intended.
	switch r.Method {
	case "Get":
		println("get")
		fallthrough
	case "DELETE":
		println("delete")
	case "POST":
		println("post")
	case "PUT":
		println("put")
	default:
		println("unhandled method")
	}
}
