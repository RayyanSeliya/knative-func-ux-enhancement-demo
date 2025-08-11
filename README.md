# Knative func CLI UX Enhancement Prototype

This prototype demonstrates key UX improvements for the Knative func CLI, showcasing professional error handling, colorized output, and improved command organization.

## 🎯 Key Improvements Demonstrated

### 1. **Enhanced Error Messages**
- Clear problem descriptions with actionable solutions
- Contextual suggestions based on the error type
- Professional formatting with color coding
- Multiple recovery paths provided

### 2. **Colorized Output & Visual Hierarchy**
- **Red**: Errors and critical issues
- **Yellow**: Warnings and next steps  
- **Green**: Success messages and confirmations
- **Blue**: Flags and technical parameters
- **Cyan**: Commands and primary actions
- **Magenta**: Examples and code snippets

### 3. **Workflow-Based Command Organization**
- Commands grouped by common developer workflows
- Clear progression paths (create → build → run → deploy)
- Contextual help showing next steps

### 4. **Smart Flag Handling**
- Automatic container mode enforcement for pack/s2i builders
- Clear error messages for conflicting flag combinations
- Consistent flag naming patterns across commands

## 🚀 Running the Demo

### Prerequisites
```bash
go version  # Requires Go 1.21+
```

### Installation
```bash
cd ux-prototype
go mod tidy
go build -o func-ux-demo
```

### Demo Commands

#### 1. **Main Help (Workflow Overview)**
```bash
./func-ux-demo help
```
Shows improved command organization with workflow-based grouping.

#### 1b. **Root Command Error**
```bash
./func-ux-demo
```
Shows enhanced error when no command is specified.

#### 2. **Enhanced Command Help**
```bash
./func-ux-demo create --help
./func-ux-demo run --help  
./func-ux-demo deploy --help
```
Demonstrates colorized help text with clear examples. Note: The `run --help` command shows the professional format from Demo 4.

#### 3. **Smart Error Handling**
```bash
# Enhanced error messages with actionable guidance
./func-ux-demo invoke                          # Demo 1: Better error for missing function
./func-ux-demo run --invalid-flag             # Demo 2: Better invalid flag handling  
./func-ux-demo create                         # Demo 3: Better missing language error
./func-ux-demo deploy                         # Demo 5: Better registry error handling
```

#### 4. **Improved Success Feedback**
```bash
./func-ux-demo create my-function --language python
./func-ux-demo run --container --builder pack
./func-ux-demo invoke --data '{"test": "data"}'
```

## 📊 UX Improvements Comparison

### BEFORE (Current func CLI)
```bash
$ func invoke
Error: '/path/to/dir' does not contain an initialized function

$ func run --invalid-flag
Error: unknown flag: --invalid-flag
Usage:
  func run [flags]
...

$ func create
Error: Required flag "language" not set.
Available language runtimes are:
  go
  node
  python
  quarkus
  rust
  springboot
  typescript

$ func deploy  
Error: Required flag "registry" not set.

$ func
# Shows basic help without workflow organization
```

### AFTER (Enhanced UX)
```bash
$ func-ux-demo invoke
Error: No function found in current directory

  This directory does not contain a function project.

To fix this:
  • Create a new function:    func create --language go --template http myfunction
  • Go to existing function:  cd /path/to/your/function
  • Specify function path:    func invoke --path /path/to/function

  Run 'func create --help' to learn about creating functions.

$ func-ux-demo run --invalid-flag
Error: Unknown flag '--invalid-flag'

  The flag '--invalid-flag' is not recognized for this command.

To fix this:
  • Use --help to see available flags
  • Did you mean one of these?
    --image          Set function image name
    --insecure       Allow insecure connections

  Common flags for 'func run':
    -t, --container  Run in container
    -r, --registry   Set registry
    -v, --verbose    Show detailed output

  Run 'func run --help' to see all available options.

$ func-ux-demo create
Error: Missing required language

  You need to specify which programming language to use.

To fix this:
  • Choose from these options:
    go          
    node        
    python      
    quarkus     
    rust        
    springboot  
    typescript  

  Example:
    func create --language python --template http myfunction

  Run 'func languages' to see detailed language information.

$ func-ux-demo deploy
Error: Container registry required

  You need to specify where to store your function's container image.

To fix this:
  • Set registry using:
    • Command line:    func deploy --registry docker.io/yourusername
    • Environment:     export FUNC_REGISTRY=docker.io/yourusername
    • Interactive mode: func deploy --confirm

  Popular registries:
    • Docker Hub:      docker.io/yourusername
    • GitHub:          ghcr.io/yourusername
    • Google Cloud:    gcr.io/your-project-id

  Example:
    func deploy --registry docker.io/myusername

  Run 'func config --help' to save registry settings permanently.

$ func-ux-demo
Error: Command required

  You need to specify a command to run.

To fix this:
  • Available commands:
    create   - Create a new function project
    run      - Run a function locally
    deploy   - Deploy function to Kubernetes cluster
    invoke   - Invoke function with test data
    help     - Show detailed help information

  Examples:
    ./func-ux-demo create --language python
    ./func-ux-demo run --help
    ./func-ux-demo help

$ func-ux-demo run --help
NAME
    func run - Run a function locally

USAGE
    func run [OPTIONS]

DESCRIPTION
    Run a function locally for development and testing.

BASIC OPTIONS
    -t, --container              Run function in container (default: true)
    -r, --registry REGISTRY      Container registry namespace
    -v, --verbose               Show detailed output

BUILD OPTIONS
        --build [auto|true|false]   Control function building (default: auto)
    -b, --builder BUILDER          Builder type: host, pack, s2i (default: pack)

ADVANCED OPTIONS
    -e, --env KEY=VALUE         Set environment variables
    -i, --image IMAGE           Override image name
    -p, --path PATH             Function directory (default: current)

EXAMPLES
    # Run function normally
    func run

    # Force rebuild and run
    func run --build=true

    # Run on host (Go only)
    func run --container=false

For more options, run: func run --help-advanced
```

## 🎨 Color Coding System

The prototype implements a professional color system:

- **🔴 Error Messages**: Red with clear problem identification (no emoji, just "Error:" prefix)
- **🟡 Warnings**: Yellow for important notices and next steps
- **🟢 Success**: Green for confirmations and completed actions
- **🔵 Technical Info**: Blue for flags, parameters, and technical details
- **🟣 Examples**: Magenta for code examples and commands
- **🟦 Commands**: Cyan for command names and primary actions

## 🏗️ Implementation for Real func CLI

This prototype demonstrates patterns that can be integrated into the actual func CLI:

1. **Error Enhancement**: Replace basic error messages with structured guidance
2. **Color System**: Add consistent color coding throughout the CLI
3. **Help Reorganization**: Group commands by workflow rather than alphabetically
4. **Smart Validation**: Add intelligent flag conflict detection and auto-correction
5. **Progress Feedback**: Provide clear status updates during operations

## 📈 Expected Impact

Based on this prototype, implementing these UX improvements would:

- **Reduce Support Questions**: Clear error messages with actionable guidance
- **Improve Onboarding**: Workflow-based help guides new users
- **Increase Productivity**: Less time debugging flag conflicts and errors
- **Enhance Consistency**: Standardized patterns across all commands
- **Better Accessibility**: Color coding with fallbacks for different terminals

## 🔧 Technical Implementation Notes

- Built with **cobra** CLI framework (same as func CLI)
- Uses **fatih/color** for cross-platform color support
- Modular command structure for easy integration
- Comprehensive error handling patterns
- Progressive disclosure of complexity

This prototype serves as a working example of the UX improvements proposed in the Knative func CLI enhancement research project.
