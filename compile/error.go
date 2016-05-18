package compile

import (
	"fmt"
)

// compileError represents error in the generation process
type gourdError struct {
	code int
	msg  string
}

// ExitCode implements cli.ExitCoder interface
func (err gourdError) ExitCode() int {
	return err.code
}

// Code return the OS return code for the error
func (err gourdError) Code() int {
	return err.code
}

// Error implements error interface
func (err gourdError) Error() string {
	return err.msg
}

// Error returns a GourdError with "fail" return code
func Error(msg string, v ...interface{}) GourdError {
	return gourdError{
		code: 1,
		msg:  fmt.Sprintf(msg, v...),
	}
}

// Errorf is an alias of Error
func Errorf(msg string, v ...interface{}) GourdError {
	return Error(msg, v...)
}

// Exit returns a GourdError with "success" return code
func Exit(msg string, v ...interface{}) GourdError {
	return gourdError{
		code: 0,
		msg:  fmt.Sprintf(msg, v...),
	}
}
