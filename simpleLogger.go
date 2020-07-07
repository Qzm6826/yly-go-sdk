package ylyOpenApi

// Info will be invoked to print the response body, request header and body
// Error will be invoked to print http error if needed

type SimpleLogger struct {
}

func (logger *SimpleLogger) Info(message string) {
	println("info  ------------>\n\n", message, "\n")
}

func (logger *SimpleLogger) Error(message string) {
	println("error ------------>\n\n", message, "\n")
}