package cmd

import (
	"fmt"
	"tactical/chat"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "ChatGPT via a CLI interface, uses GPT-4o-mini",
	Long:  "This is a ChatGPT via a CLI interface, which uses the latest GPT-4o-mini model.",
	Run:   processChat,
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
func processChat(cmd *cobra.Command, args []string) {
	fmt.Println("Awaiting Response...")
	response := chat.GetChatResponse(args[0])
	out, err := glamour.Render(response, "dark")
	if err != nil {
		log.Fatal("Error rendering markdown output", "Error", err)
	}
	fmt.Println(out)
}
