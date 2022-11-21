package console

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/reizt/todo/src/core"
	"github.com/reizt/todo/src/utils"
)

type renderer struct{}

func Init() (*renderer, error) {
	return &renderer{}, nil
}

func (renderer) Version(version string) {
	fmt.Printf("todo cli %s\n", version)
}

func (renderer) Help() {
	fmt.Println(`Manage your TODO by CLI

USAGE
	todo <command> [flags]

COMMANDS
	list:	List todo
	add:	Add todo
	mod:	Modify todo
	del:	Delete todo

FLAGS
	list
	add
		-t, --title       Title
		-d, --description Description
		mod <id>
		-t, --title       Title
		-d, --description Description
		del <id>`)
}

func (renderer) AddHelp() {
	fmt.Println(`USAGE: todo add [flags]
	-t, --title       Title
	-d, --description Description`)
}

func (renderer) ModHelp() {
	fmt.Println(`USAGE: todo mod <id> [flags]
	-t, --title       Title
	-d, --description Description`)
}

func (renderer) DelHelp() {
	fmt.Println(`USAGE: todo del <id>`)
}

func (renderer) TitleIsRequired() {
	utils.PrintRed("Title is required.\n")
}

func (renderer) TitleIsTooLong(maxLength int) {
	utils.PrintRed(fmt.Sprintf("Title is too long. Title can be up to %d characters.\n", maxLength))
}

func (renderer) InvalidInput() {
	utils.PrintRed("Invalid input.\n")
}

func (renderer) UnexpectedError() {
	utils.PrintRed("ERROR: Unexpected error.\n")
}

func (renderer) List(todos []core.Todo) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(writer, "ID\tTitle\tDescription\t")
	for _, todo := range todos {
		fmt.Fprintf(writer, "%s\t%s\t%s\t\n", todo.ID, utils.NilSafe(todo.Title), utils.NilSafe(todo.Description))
	}
	writer.Flush()
}

func (renderer) ListFailedToFindMany(err error) {
	utils.PrintRed("ERROR: Failed to find todos.\n")
	utils.PrintRed(err.Error())
}

func (renderer) AddSucceeded(todo core.Todo) {
	utils.PrintGreen("Todo was added successfully.\n")
}

func (renderer) AddFailedToInsert(err error) {
	utils.PrintRed("ERROR: Failed to insert todo.\n")
	utils.PrintRed(err.Error())
}

func (renderer) NotFound(id string) {
	fmt.Printf("ERROR: Todo with the id \"%s\" doesn't exist.\n", id)
}

func (renderer) ModSucceeded(todo core.Todo) {
	utils.PrintGreen("Todo was updated successfully.\n")
}

func (renderer) ModFailedToUpdate(err error) {
	utils.PrintRed("ERROR: Failed to update todo.\n")
	utils.PrintRed(err.Error())
}

func (renderer) DelSucceeded(todo core.Todo) {
	utils.PrintGreen("Todo was deleted successfully.\n")
}

func (renderer) DelFailed(err error) {
	utils.PrintRed("ERROR: Failed to delete todo.\n")
	utils.PrintRed(err.Error())
}

func (renderer) ClearSucceeded() {
	utils.PrintGreen("All todos were deleted successfully.\n")
}

func (renderer) ClearFailed(err error) {
	utils.PrintRed("ERROR: Failed to clear todo.\n")
	utils.PrintRed(err.Error())
}

func (renderer) ClearCanceled() {
	fmt.Println("Canceled.")
}
