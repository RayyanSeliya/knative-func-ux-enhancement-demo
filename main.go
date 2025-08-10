package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

// Color definitions for consistent UX
var (
	// Error colors
	errorColor   = color.New(color.FgRed, color.Bold)
	warningColor = color.New(color.FgYellow, color.Bold)
	successColor = color.New(color.FgGreen, color.Bold)
	
	// Help text colors
	commandColor     = color.New(color.FgCyan, color.Bold)
	flagColor        = color.New(color.FgBlue)
	exampleColor     = color.New(color.FgMagenta)
	descriptionColor = color.New(color.FgWhite)
	
	// Section headers
	headerColor = color.New(color.FgYellow, color.Bold, color.Underline)
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "func-ux-demo",
		Short: "Enhanced func CLI UX Demonstration",
		Long: formatHelpText(`
Enhanced Knative func CLI with improved UX patterns

This demo showcases professional CLI improvements including:
• Clear command organization with workflow-based grouping
• Enhanced error messages with actionable guidance  
• Colorized output for better visual hierarchy
• Consistent flag patterns across commands
• Progressive disclosure of complexity

Common Workflows:
  Development:  create → build → run → invoke
  Deployment:   create → build → deploy
  Events:       create → deploy → subscribe
		`),
		Run: func(cmd *cobra.Command, args []string) {
			showWorkflowHelp()
		},
	}

	// Add demo commands
	rootCmd.AddCommand(createDemoCommand())
	rootCmd.AddCommand(runDemoCommand())
	rootCmd.AddCommand(deployDemoCommand())
	rootCmd.AddCommand(invokeDemoCommand())
	rootCmd.AddCommand(errorDemoCommand())

	if err := rootCmd.Execute(); err != nil {
		showEnhancedError("Command execution failed", err.Error(), []string{
			"Check command syntax with --help",
			"Verify all required flags are provided",
			"Ensure you're in a function directory",
		})
		os.Exit(1)
	}
}

func formatHelpText(text string) string {
	return descriptionColor.Sprint(text)
}

func showWorkflowHelp() {
	fmt.Println()
	headerColor.Println("KNATIVE FUNCTIONS CLI")
	fmt.Println()
	
	fmt.Println(headerColor.Sprint("USAGE"))
	fmt.Printf("    %s\n", flagColor.Sprint("func <command> [options]"))
	fmt.Println()
	
	printCommandCategory("GETTING STARTED", []CommandInfo{
		{"create", "Create a new function from template"},
		{"run", "Run function locally for development"},
		{"invoke", "Test function with sample data"},
	})
	
	printCommandCategory("DEPLOYMENT", []CommandInfo{
		{"build", "Build function container image"},
		{"deploy", "Deploy function to Kubernetes cluster"},
		{"delete", "Remove function from cluster"},
	})
	
	printCommandCategory("MANAGEMENT", []CommandInfo{
		{"list", "List deployed functions"},
		{"describe", "Show detailed function information"},
		{"logs", "View function execution logs"},
	})
	
	printCommandCategory("CONFIGURATION", []CommandInfo{
		{"config", "Manage function configuration"},
		{"subscribe", "Configure event subscriptions"},
	})
	
	printCommandCategory("UTILITY COMMANDS", []CommandInfo{
		{"languages", "List supported runtimes"},
		{"templates", "List available templates"},
		{"version", "Show version information"},
	})
	
	fmt.Printf("Use %s for detailed information about each command.\n", exampleColor.Sprint("func <command> --help"))
	fmt.Println()
}

type CommandInfo struct {
	Name        string
	Description string
}

func printCommandCategory(category string, commands []CommandInfo) {
	commandColor.Printf("  %s:\n", category)
	for _, cmd := range commands {
		fmt.Printf("    %s  %s\n", 
			flagColor.Sprintf("%-10s", cmd.Name), 
			descriptionColor.Sprint(cmd.Description))
	}
	fmt.Println()
}

func showEnhancedError(title, details string, suggestions []string) {
	fmt.Println()
	errorColor.Printf("✗ Error: %s\n", title)
	fmt.Println()
	
	if details != "" {
		fmt.Printf("  %s\n", details)
		fmt.Println()
	}
	
	if len(suggestions) > 0 {
		warningColor.Println("💡 Suggestions:")
		for _, suggestion := range suggestions {
			fmt.Printf("  • %s\n", suggestion)
		}
		fmt.Println()
	}
}

func showSuccess(message string) {
	successColor.Printf("✓ %s\n", message)
}