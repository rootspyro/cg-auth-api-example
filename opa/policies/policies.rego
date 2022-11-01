package policies

default allow = false

allow = true {
	input.request.method == "GET"
}
