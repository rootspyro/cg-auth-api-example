package authz

default allow = false
allow {
    input.method == "POST"
    input.path = ["tasks"]
    input.user == input.owner
}

allow {
    input.method == "GET"
    input.path = ["tasks"]
    input.user == input.owner
}
