package main

// User defines the UserModel.
// Use this to check whether a User is a Premium user or not.

type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users.
// Returns false if a process had to be killed
func HandleRequest(process func(), u *User) bool {
	process()
	return true
}

func main() {
	RunMockServer()
}
