package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func createDemoCommand() *cobra.Command {
	var language string
	var template string
	var registry string

	cmd := &cobra.Command{
		Use:   "create [NAME]",
		Short: "Create a new function project",
		Long: formatHelpText(`
Create a new function project from a template.

This command scaffolds a new function project with the specified runtime
and template, setting up the basic structure needed for development.
		`),
		Example: exampleColor.Sprint(`  # Create a Python HTTP function
  func-ux-demo create my-function --language python

  # Create a Go CloudEvent function  
  func-ux-demo create event-handler --language go --template cloudevents

  # Create with custom registry
  func-ux-demo create api-func --language node --registry ghcr.io/myorg`),
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateDemo(args, language, template, registry)
		},
	}

	// Organized flag groups
	cmd.Flags().StringVarP(&language, "language", "l", "", "Function runtime language (required)")
	cmd.Flags().StringVarP(&template, "template", "t", "http", "Function template type")
	cmd.Flags().StringVarP(&registry, "registry", "r", "", "Container registry for function images")
	
	cmd.MarkFlagRequired("language")

	return cmd
}

func runCreateDemo(args []string, language, template, registry string) error {
	var name string
	if len(args) > 0 {
		name = args[0]
	} else {
		name = "my-function"
	}

	// Validate language
	validLanguages := []string{"go", "python", "node", "java", "rust"}
	if !contains(validLanguages, language) {
		return fmt.Errorf("invalid language '%s'", language)
	}

	// Enhanced success feedback
	fmt.Println()
	successColor.Printf("✓ Created function project: %s\n", name)
	fmt.Printf("  Language: %s\n", flagColor.Sprint(language))
	fmt.Printf("  Template: %s\n", flagColor.Sprint(template))
	if registry != "" {
		fmt.Printf("  Registry: %s\n", flagColor.Sprint(registry))
	}
	fmt.Println()
	
	warningColor.Println("🔄 Next Steps:")
	fmt.Printf("  1. %s\n", exampleColor.Sprintf("cd %s", name))
	fmt.Printf("  2. %s\n", exampleColor.Sprint("func-ux-demo run --help"))
	fmt.Printf("  3. %s\n", exampleColor.Sprint("func-ux-demo run"))
	fmt.Println()

	return nil
}

func runDemoCommand() *cobra.Command {
	var container bool
	var build string
	var registry string
	var address string
	var builder string

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a function locally",
		Long: formatHelpText(`
Run a function locally for development and testing. The function can be executed
either in a container or directly on the host system.
		`),
		Example: exampleColor.Sprint(`  # Run function with automatic building
  func run

  # Force rebuild and run in container
  func run --build=true

  # Run on host system (Go functions only)
  func run --container=false

  # Run with custom environment
  func run --env DEBUG=true --env PORT=9000`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRunDemo(container, build, registry, address, builder)
		},
	}

	// Organized flags with clear descriptions
	cmd.Flags().BoolVarP(&container, "container", "c", false, "Run in container mode")
	cmd.Flags().StringVar(&build, "build", "auto", "Build behavior: auto|true|false")
	cmd.Flags().StringVarP(&registry, "registry", "r", "", "Container registry URL")
	cmd.Flags().StringVarP(&address, "address", "a", "localhost:8080", "Address to serve function")
	cmd.Flags().StringVarP(&builder, "builder", "b", "pack", "Builder type: pack|s2i|host")

	return cmd
}

func runRunDemo(container bool, build, registry, address, builder string) error {
	// Check if in function directory
	if !isInFunctionDirectory() {
		showEnhancedError(
			"Not in function directory",
			"This command must be run from within a function project directory.",
			[]string{
				"Run 'func-ux-demo create' to create a new function",
				"Navigate to an existing function directory",
				"Use 'func-ux-demo info' to verify function project",
			})
		return fmt.Errorf("not in function directory")
	}

	// Smart container enforcement for pack builder
	if builder == "pack" && !container {
		warningColor.Println("⚠️  Pack builder requires container mode")
		fmt.Printf("   Automatically enabling %s\n", flagColor.Sprint("--container=true"))
		container = true
	}

	fmt.Println()
	successColor.Println("🚀 Starting function...")
	fmt.Printf("  Mode: %s\n", flagColor.Sprint(getRunMode(container)))
	fmt.Printf("  Builder: %s\n", flagColor.Sprint(builder))
	fmt.Printf("  Address: %s\n", flagColor.Sprint(address))
	fmt.Println()
	
	successColor.Printf("✓ Function running on %s\n", address)
	fmt.Printf("  %s\n", exampleColor.Sprint("curl http://"+address))
	fmt.Printf("  %s\n", exampleColor.Sprint("func-ux-demo invoke"))
	fmt.Println()

	return nil
}

func deployDemoCommand() *cobra.Command {
	var registry string
	var namespace string
	var build bool

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy function to Kubernetes cluster",
		Long: formatHelpText(`
Deploy the function as a Knative Service to a Kubernetes cluster.

This command builds the function container (if needed) and deploys it
to the current Kubernetes context and namespace.
		`),
		Example: exampleColor.Sprint(`  # Deploy with auto-build
  func-ux-demo deploy --registry ghcr.io/myorg

  # Deploy to specific namespace
  func-ux-demo deploy --registry ghcr.io/myorg --namespace production

  # Deploy without building
  func-ux-demo deploy --build=false`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDeployDemo(registry, namespace, build)
		},
	}

	cmd.Flags().StringVarP(&registry, "registry", "r", "", "Container registry URL (required)")
	cmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Kubernetes namespace")
	cmd.Flags().BoolVar(&build, "build", true, "Build container before deploying")
	
	cmd.MarkFlagRequired("registry")

	return cmd
}

func runDeployDemo(registry, namespace string, build bool) error {
	if !isInFunctionDirectory() {
		showEnhancedError(
			"Not in function directory", 
			"This command must be run from within a function project directory.",
			[]string{
				"Run 'func-ux-demo create' to create a new function",
				"Navigate to an existing function directory",
			})
		return fmt.Errorf("not in function directory")
	}

	fmt.Println()
	successColor.Println("🚀 Deploying function...")
	
	if build {
		fmt.Printf("  %s Building container image...\n", warningColor.Sprint("⏳"))
		fmt.Printf("  %s Pushing to %s...\n", warningColor.Sprint("⏳"), flagColor.Sprint(registry))
	}
	
	fmt.Printf("  %s Deploying to Kubernetes...\n", warningColor.Sprint("⏳"))
	
	if namespace != "" {
		fmt.Printf("  Namespace: %s\n", flagColor.Sprint(namespace))
	}
	
	fmt.Println()
	successColor.Println("✓ Function deployed successfully!")
	fmt.Printf("  URL: %s\n", exampleColor.Sprint("https://my-function.example.com"))
	fmt.Printf("  Test: %s\n", exampleColor.Sprint("func-ux-demo invoke --target https://my-function.example.com"))
	fmt.Println()

	return nil
}

func invokeDemoCommand() *cobra.Command {
	var target string
	var data string
	var contentType string

	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "Invoke function with test data",
		Long: formatHelpText(`
Invoke a function with test data for development and testing.

Can invoke local functions or deployed functions by specifying a target URL.
		`),
		Example: exampleColor.Sprint(`  # Invoke local function
  func-ux-demo invoke

  # Invoke with custom data
  func-ux-demo invoke --data '{"name": "World"}'

  # Invoke deployed function
  func-ux-demo invoke --target https://my-function.example.com`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInvokeDemo(target, data, contentType)
		},
	}

	cmd.Flags().StringVarP(&target, "target", "t", "", "Function URL (default: local)")
	cmd.Flags().StringVarP(&data, "data", "d", "{}", "Request data")
	cmd.Flags().StringVar(&contentType, "content-type", "application/json", "Content type")

	return cmd
}

func runInvokeDemo(target, data, contentType string) error {
	if target == "" && !isInFunctionDirectory() {
		showEnhancedError(
			"No function to invoke",
			"No target specified and not in a function directory.",
			[]string{
				"Use --target to specify a deployed function URL",
				"Run from a function directory to invoke locally",
				"Start local function with 'func-ux-demo run'",
			})
		return fmt.Errorf("no function to invoke")
	}

	invokeTarget := target
	if invokeTarget == "" {
		invokeTarget = "http://localhost:8080"
	}

	fmt.Println()
	successColor.Printf("🔥 Invoking function: %s\n", invokeTarget)
	fmt.Printf("  Data: %s\n", flagColor.Sprint(data))
	fmt.Printf("  Content-Type: %s\n", flagColor.Sprint(contentType))
	fmt.Println()
	
	successColor.Println("✓ Response received:")
	fmt.Printf("  Status: %s\n", successColor.Sprint("200 OK"))
	fmt.Printf("  Body: %s\n", exampleColor.Sprint(`{"message": "Hello, World!"}`))
	fmt.Println()

	return nil
}

func errorDemoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "error-demo [TYPE]",
		Short: "Demonstrate enhanced error handling",
		Long: formatHelpText(`
Demonstrate various enhanced error scenarios and their improved messaging.

Available error types:
  missing-function  - Show error when not in function directory
  invalid-flag      - Show error for invalid flag usage  
  missing-registry  - Show error for missing required registry
  flag-conflict     - Show error for conflicting flags
  deploy-no-function - Show error for deploy without function project
		`),
		ValidArgs: []string{"missing-function", "invalid-flag", "missing-registry", "flag-conflict", "deploy-no-function"},
		RunE: func(cmd *cobra.Command, args []string) error {
			errorType := "missing-function"
			if len(args) > 0 {
				errorType = args[0]
			}
			return runErrorDemo(errorType)
		},
	}

	return cmd
}

func runErrorDemo(errorType string) error {
	fmt.Printf("\n%s Demonstrating: %s\n\n", 
		warningColor.Sprint("🎭"), 
		commandColor.Sprint(errorType))

	switch errorType {
	case "missing-function":
		showEnhancedError(
			"No function found in current directory",
			"This directory does not contain an initialized function project.",
			[]string{
				"Create a new function:    func create --language go --template http <name>",
				"Navigate to existing function: cd /path/to/function", 
				"Specify function path:     func invoke --path /path/to/function",
				"",
				"Run 'func create --help' for more information about creating functions.",
			})

	case "invalid-flag":
		showEnhancedError(
			"Unknown flag '--invalid-flag'",
			"The flag '--invalid-flag' is not recognized for this command.",
			[]string{
				"Use --help to see available flags",
				"Did you mean: --build, --builder, or --registry?",
				"Check flag spelling and format",
			})

	case "missing-registry":
		showEnhancedError(
			"Container registry required",
			"A container registry is required to store function images during deployment.",
			[]string{
				"Command flag:    func deploy --registry docker.io/username",
				"Environment:     export FUNC_REGISTRY=docker.io/username", 
				"Interactive:     func deploy --confirm",
				"",
				"Popular registry options:",
				"• Docker Hub:      docker.io/username",
				"• GitHub:          ghcr.io/username",
				"• Google Cloud:    gcr.io/project-id",
				"• Azure:           username.azurecr.io",
				"",
				"Run 'func config --help' for persistent configuration options.",
			})

	case "flag-conflict":
		showEnhancedError(
			"Invalid flag combination",
			"The --builder=pack option requires container execution, but --container=false was specified.",
			[]string{
				"Remove --container=false:     func run --builder=pack",
				"Use host builder:             func run --container=false --builder=host",
				"Use default configuration:    func run",
				"",
				"Note: Pack and S2I builders require container execution. Use 'host' builder for non-containerized runs.",
			})

	case "deploy-no-function":
		showEnhancedError(
			"No function project found",
			"This directory doesn't contain a function project that can be deployed.",
			[]string{
				"Create a new function here:     func create --language python --template http myfunction",
				"Go to existing function:       cd /path/to/your/function",
				"Deploy from specific path:     func deploy --path /path/to/function",
				"",
				"What is a function project?",
				"A function project contains your code, configuration files, and metadata that",
				"func needs to build and deploy your function.",
				"",
				"Examples of creating functions:",
				"  func create --language go --template http my-go-func",
				"  func create --language node --template cloudevent my-node-func", 
				"  func create --language python --template http my-python-func",
				"",
				"Run 'func create --help' to see all available languages and templates.",
			})

	default:
		showEnhancedError(
			"Unknown error type",
			fmt.Sprintf("Error type '%s' is not recognized.", errorType),
			[]string{
				"Available types: missing-function, invalid-flag, missing-registry, flag-conflict, deploy-no-function",
				"Use: func-ux-demo error-demo --help",
			})
	}

	return nil
}

// Helper functions
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func isInFunctionDirectory() bool {
	// Check for func.yaml or other function indicators
	_, err := os.Stat("func.yaml")
	if err == nil {
		return true
	}
	
	// Check for common function files
	files := []string{"package.json", "go.mod", "requirements.txt", "pom.xml"}
	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			return true
		}
	}
	
	return false
}

func getRunMode(container bool) string {
	if container {
		return "Container"
	}
	return "Source"
}