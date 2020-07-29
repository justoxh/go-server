package render

const (
	Failed = -1
	Ok     = 0
)

var statusMsg = map[int]string{
	Failed: "Failed",
	Ok:     "Success",
}

func getMessage(code int) string {
	return statusMsg[code]
}
