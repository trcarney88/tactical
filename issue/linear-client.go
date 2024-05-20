package issue

import (
	"context"
	"os"

	"github.com/charmbracelet/log"
	"github.com/machinebox/graphql"
)

var linerApiUrl = "https://api.linear.app/graphql"

func getAllAssignedIssues() []Issue {
	assigneeId := os.Getenv("assigneeId")

	client := graphql.NewClient(linerApiUrl)

	query := `query Issues($filter: IssueFilter) {
  	issues(filter: $filter) {
    	nodes {
     		identifier
      	title
       	id
	      assignee {
	        name
	        email
	      }
			  description
	      state {
	        name
	      }
			}
		}
	}
	`

	request := graphql.NewRequest(query)

	var assigneeFilter = AssigneeFilter{
		Assignee: Assignee{
			Id: Id{
				Eq: assigneeId,
			},
		},
	}
	request.Var("filter", assigneeFilter)
	request.Header.Add("Authorization", os.Getenv("linearKey"))

	var resp IssuesResponse
	err := client.Run(context.Background(), request, &resp)
	if err != nil {
		log.Fatal("Error with Linear API call", "Error", err)
	}

	return resp.Issues.Nodes
}

type CreateNewIssueResponse struct {
	Issue   Issue
	Success bool
}

func createNewIssue(title string, description string) CreateNewIssueResponse {
	teamId := os.Getenv("teamId")
	assigneeId := os.Getenv("assigneeId")
	backlogId := os.Getenv("backlogId")

	client := graphql.NewClient(linerApiUrl)

	mutation := `mutation IssueCreate($input: IssueCreateInput!) {
  	issueCreate(input: $input) {
	    issue {
	      identifier
	      title
	      id
	      state {
	        name
	      }
	      assignee {
	        email
	        name
	      }
	      description
	    }
			success
		}
	}`

	request := graphql.NewRequest(mutation)

	var createInput = CreateIssueVars{
		Title:       title,
		Description: description,
		TeamId:      teamId,
		StateId:     backlogId,
		AssigneeId:  assigneeId,
	}

	request.Var("input", createInput)
	request.Header.Add("Authorization", os.Getenv("linearKey"))

	var createResponse IssueCreate
	err := client.Run(context.Background(), request, &createResponse)
	if err != nil {
		log.Fatal("Error creating issue", "Error", err)
	}

	return CreateNewIssueResponse{
		Issue:   createResponse.IssueCreate.Issue,
		Success: createResponse.IssueCreate.Success,
	}
}

type UpdateIssueResponse struct {
	Issue   Issue
	Success bool
}

func updateStatus(issueId string, stateId string) UpdateIssueResponse {
	log.Info("Updating Status...", "IssueId", issueId, "Status Id", stateId)

	client := graphql.NewClient(linerApiUrl)

	mutation := `mutation IssueUpdate($input: IssueUpdateInput!, $issueUpdateId: String!) {
  	issueUpdate(input: $input, id: $issueUpdateId) {
    	issue {
	      identifier
	      title
	      id
	      assignee {
	        name
	        email
				}
      	description
       	state {
        	name
        }
      }
      success
    }
  }`

	request := graphql.NewRequest(mutation)

	statusIdMap := map[string]string{"stateId": stateId}
	request.Var("input", statusIdMap)
	request.Var("issueUpdateId", issueId)
	request.Header.Add("Authorization", os.Getenv("linearKey"))

	var updateResponse IssueUpdate
	err := client.Run(context.Background(), request, &updateResponse)
	if err != nil {
		log.Fatal("Error updating issue", "Error", err)
	}

	return UpdateIssueResponse{
		Issue:   updateResponse.IssueUpdate.Issue,
		Success: updateResponse.IssueUpdate.Success,
	}
}
