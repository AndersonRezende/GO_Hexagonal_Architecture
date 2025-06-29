package util

import (
	"bufio"
	"fmt"
	"gohexarc/internal/domain"
	"io"
	"strings"
)

const List string = "list"
const Create string = "create"
const Update string = "update"
const Delete string = "delete"
const Get string = "get"

func PrintUsage(w io.Writer) {
	fmt.Fprintln(w, "\nAvailable cli commands:")
	fmt.Fprintln(w, "list   - List users")
	fmt.Fprintln(w, "create - Create user")
	fmt.Fprintln(w, "update - Update user")
	fmt.Fprintln(w, "delete - Delete user")
	fmt.Fprintln(w, "get - Get user")
	fmt.Fprintln(w, "exit - Exit the CLI")
}

func PrintUser(w io.Writer, user domain.User) {
	fmt.Fprintln(w, "\n-----User Details-----")
	fmt.Fprintf(w, "User ID: %s\n", user.ID)
	fmt.Fprintf(w, "Name: %s\n", user.Name)
	fmt.Fprintf(w, "Email: %s\n", user.Email)
	fmt.Fprintln(w, "----------------------")
}

func ReadInput(inputText string, r *bufio.Reader, w io.Writer) string {
	fmt.Fprint(w, "\n"+inputText)
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}
