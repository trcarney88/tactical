package cmd

import (
	"fmt"

	"tactical/issue"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Create and View Linear issues",
	Long:  "Create and View Linear Issues",
	Run:   processIssue,
}

func init() {
	rootCmd.AddCommand(issueCmd)

	issueCmd.Flags().BoolP("get-all", "g", false, "All all assigned issues from linear")
	issueCmd.Flags().BoolP("create", "c", false, "Create a new issue in Linear")
	issueCmd.Flags().BoolP("update", "u", false, "Update the status of an issue in Linear")
}

func processIssue(cmd *cobra.Command, args []string) {
	isGetAll, _ := cmd.Flags().GetBool("get-all")
	isCreate, _ := cmd.Flags().GetBool("create")
	isUpdate, _ := cmd.Flags().GetBool("update")

	if isGetAll {
		getAllIssues()
	} else if isCreate {
		createIssue()
	} else if isUpdate {
		updateIssue()
	} else {
		cmd.Help()
	}
}

func getAllIssues() {
	issues := issue.GetAssignedIssues()

	for _, issue := range issues {
		msg := fmt.Sprintf("%s: %s - %s", issue.Identifier, issue.Title, issue.State.Name)
		fmt.Println(msg)
	}
}

func createIssue() {
	var title string
	var description string

	// Change KeyMap for huh.NewText()
	defaultKeyMap := huh.NewDefaultKeyMap()
	var keyMap = huh.KeyMap{
		Quit:        defaultKeyMap.Quit,
		Confirm:     defaultKeyMap.Confirm,
		Input:       defaultKeyMap.Input,
		Select:      defaultKeyMap.Select,
		MultiSelect: defaultKeyMap.MultiSelect,
		Note:        defaultKeyMap.Note,
		Text: huh.TextKeyMap{
			Next:    defaultKeyMap.Text.Next,
			Prev:    defaultKeyMap.Text.Prev,
			NewLine: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "new line")),
			Editor:  defaultKeyMap.Text.Editor,
			Submit:  key.NewBinding(key.WithKeys("ctrl+s"), key.WithHelp("ctrl+s", "submit")),
		},
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("New Linear Issue"),
			huh.NewInput().
				Value(&title).
				Title("Title:").
				Placeholder("Issue Title").
				Description("This will appear on all the boards in linear, the issue id is automatically adding by linear"),
			huh.NewText().
				Value(&description).
				Title("Description").
				Placeholder("Description of the issue").
				Description("A Description of the problem and the work needed to fix the problem").
				Lines(5),
		),
	).WithKeyMap(&keyMap)

	err := form.Run()
	if err != nil {
		log.Fatal("Error with form", "Error", err)
	}

	issue.CreateIssue(title, description)
}

func updateIssue() {
	var selectedIssue string
	var selectedStatus string

	issues := issue.GetAssignedIssues()

	issuesMap := make(map[string]issue.Issue)
	var keys []string
	for _, issue := range issues {
		key := fmt.Sprintf("%s: %s - %s", issue.Identifier, issue.Title, issue.State.Name)

		keys = append(keys, key)
		issuesMap[key] = issue
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Update the status of an issue"),
			huh.NewSelect[string]().
				Options(huh.NewOptions(keys...)...).
				Title("Issue you want to update").
				Value(&selectedIssue),
			huh.NewSelect[string]().
				Options(huh.NewOptions(issue.AllStatuses...)...).
				Title("New Status for Issue").
				Value(&selectedStatus),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal("Error with form", "Error", err)
	}

	log.Info("Update Issue", "Selected Issue", issuesMap[selectedIssue])
	issue.UpdateIssue(issuesMap[selectedIssue], selectedStatus)
}
