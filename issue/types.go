package issue

// GetAllIssues Response
type IssuesResponse struct {
	Issues Nodes `json:"issues"`
}

type Nodes struct {
	Nodes []Issue `json:"nodes"`
}

type Issue struct {
	Identifier  string        `json:"identifier"`
	Title       string        `json:"title"`
	Id          string        `json:"id"`
	Assignee    IssueAssignee `json:"assignee"`
	Description string        `json:"description"`
	State       State         `json:"state"`
}

type IssueAssignee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type State struct {
	Name string `json:"name"`
}

// Input Variable type for GetAllIssues Query
type AssigneeFilter struct {
	Assignee Assignee `json:"assignee"`
}
type Id struct {
	Eq string `json:"eq"`
}
type Assignee struct {
	Id Id `json:"id"`
}

// Create Issue Response
type IssueCreated struct {
	Issue   Issue `json:"issue"`
	Success bool  `json:"success"`
}

type IssueCreate struct {
	IssueCreate IssueCreated `json:"issueCreate"`
}

// Input Variable for Create Issue
type CreateIssueVars struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	TeamId      string `json:"teamId"`
	StateId     string `json:"stateId"`
	AssigneeId  string `json:"assigneeId"`
}

// Update Issue Response
type IssueUpdate struct {
	IssueUpdate IssueUpdated `json:"issueUpdate"`
}

type IssueUpdated struct {
	Issue   Issue `json:"issue"`
	Success bool  `json:"success"`
}
