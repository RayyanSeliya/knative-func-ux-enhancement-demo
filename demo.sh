#!/bin/bash

echo "============================================"
echo "Knative func CLI UX Enhancement Demo"
echo "============================================"
echo

echo "Building prototype..."
go mod tidy
go build -o func-ux-demo
echo

if [ ! -f func-ux-demo ]; then
    echo "Build failed! Please check for errors."
    exit 1
fi

echo "============================================"
echo "1. Enhanced Workflow Help"
echo "============================================"
./func-ux-demo
echo
read -p "Press Enter to continue..."

echo "============================================"
echo "2. Colorized Command Help"
echo "============================================"
echo "Running: ./func-ux-demo create --help"
./func-ux-demo create --help
echo
read -p "Press Enter to continue..."

echo "============================================"
echo "3. Enhanced Error Demonstrations"
echo "============================================"
echo
echo "3a. Missing Function Error:"
./func-ux-demo error-demo missing-function
echo
read -p "Press Enter to continue..."

echo "3b. Invalid Flag Error:"
./func-ux-demo error-demo invalid-flag
echo
read -p "Press Enter to continue..."

echo "3c. Missing Registry Error:"
./func-ux-demo error-demo missing-registry
echo
read -p "Press Enter to continue..."

echo "3d. Flag Conflict Error:"
./func-ux-demo error-demo flag-conflict
echo
read -p "Press Enter to continue..."

echo "============================================"
echo "4. Improved Success Feedback"
echo "============================================"
echo
echo "4a. Create Function:"
./func-ux-demo create my-function --language python
echo
read -p "Press Enter to continue..."

echo "4b. Run Function (with smart container enforcement):"
./func-ux-demo run --builder pack
echo
read -p "Press Enter to continue..."

echo "4c. Invoke Function:"
./func-ux-demo invoke --data '{"test": "data"}'
echo
read -p "Press Enter to continue..."

echo "============================================"
echo "Demo Complete!"
echo "============================================"
echo
echo "This prototype demonstrates:"
echo "- Enhanced error messages with actionable guidance"
echo "- Colorized output for better visual hierarchy"  
echo "- Workflow-based command organization"
echo "- Smart flag handling and validation"
echo
echo "For more details, see: ux-prototype/README.md"