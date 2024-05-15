package cmd

import (
	"fmt"
	"tactical/chat"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "ChatGPT via a CLI interface, uses GPT-4o",
	Long:  "This is a ChatGPT via a CLI interface, which uses the latest GPT-4o model.",
	Run:   processChat,
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
func processChat(cmd *cobra.Command, args []string) {
	fmt.Println("Awaiting Response...")
	response := chat.GetChatResponse(args[0])
	fmt.Println(response)
}
