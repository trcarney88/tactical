# TACTICAL

Tactical (Todd's Awesome Cli for Triage, Incident Correction, and AnaLytics) is a CLI
with some helpful tools (i.e ChatGPT and Linear) to make life easier for the developer
and an easy way to run various scripts written in Go. I use this CLI to do my day to day
work and to run scripts to handle frequent asks from others at the company.

## How to run Tactical

First you will need to create a .env file that has all the secrets needed to power the
integrations.

### .env File

This file should be name `.env` and located in the root of the project.

#### Chat

To power the chat integration, you only need to add a ChatGPT API Key in the following
format. Ensure the key you create has access to the 4o-mini model.

`openAiKey=<secret key from Open AI>`

#### Issues

To power the issues integration, you will need the following keys from Linear, a linear
api key, your assignee id, a team id, and the id's for the todo, in progress, done,
canceled, backlog, and duplicate statuses you use. The format for these keys is shown below.

`linearKey=<linear api key>`
`assigneeId=<assignee id>`
`teamId=<team id>`
`todoId=<todo status id>`
`inProgressId=<in progress status id>`
`doneId=<done status id>`
`canceledId=<canceled status id>`
`backlogId=<backlog status id>`
`duplicateId=<duplicate status id>`git push --set-upstream origin implement-news-cmd

### How to install Tactical

#### Prerequisites

1. Install go on your computer
2. Add your local go/bin to your path

If you need help with any of these things, look [here](https://go.dev/doc/tutorial/compile-install)

#### Install

To install Tactical you need to perform the following steps

1. Clone/Fork this repo
2. `cd <path/to/tactical>`
3. `go build && go install`
