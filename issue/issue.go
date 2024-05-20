package issue

import (
	"fmt"
	"log"
	"os"
)

var AllStatuses = []string{
	"Backlog",
	"To Do",
	"In Progress",
	"Done",
	"Canceled",
	"Duplicate",
}

func GetAssignedIssues() []Issue {
	issues := getAllAssignedIssues()

	return issues
}

func CreateIssue(title string, description string) {
	resp := createNewIssue(title, description)

	fmt.Println("Success:", resp.Success)

	if resp.Success == true {
		msg := fmt.Sprintf("%s: %s - %s", resp.Issue.Identifier, resp.Issue.Title, resp.Issue.State.Name)
		fmt.Println(msg)
	}
}

func UpdateIssue(issue Issue, status string) {
	var statusId string
	switch status {
	case "Backlog":
		statusId = os.Getenv("backlogId")
	case "To Do":
		statusId = os.Getenv("todoId")
	case "In Progress":
		statusId = os.Getenv("inProgressId")
	case "Done":
		statusId = os.Getenv("doneId")
	case "Canceled":
		statusId = os.Getenv("canceledId")
	case "Duplicate":
		statusId = os.Getenv("duplicateId")
	default:
		log.Fatal("Unsupported Status", "Status", status)
	}
	resp := updateStatus(issue.Id, statusId)

	fmt.Println("Success:", resp.Success)

	if resp.Success == true {
		msg := fmt.Sprintf("%s: %s - %s", resp.Issue.Identifier, resp.Issue.Title, resp.Issue.State.Name)
		fmt.Println("Updated Issue:", msg)
	}
}
