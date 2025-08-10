#!/bin/bash

echo "Building func CLI UX prototype..."
echo

go mod tidy
if [ $? -ne 0 ]; then
    echo "Failed to download dependencies"
    exit 1
fi

go build -o func-ux-demo
if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

echo "Build successful!"
echo
echo "Testing the prototype..."
echo

echo "============================================"
echo "1. Main Help (Workflow Overview)"
echo "============================================"
./func-ux-demo
echo

echo "============================================"
echo "2. Create Command Help"
echo "============================================"
./func-ux-demo create --help
echo

echo "============================================"
echo "3. Enhanced Error Demo - Missing Function"
echo "============================================"
./func-ux-demo error-demo missing-function
echo

echo "============================================"
echo "4. Enhanced Error Demo - Invalid Flag"
echo "============================================"
./func-ux-demo error-demo invalid-flag
echo

echo "============================================"
echo "5. Create Function Demo"
echo "============================================"
./func-ux-demo create my-function --language python
echo

echo "============================================"
echo "6. Run Function Demo (Smart Container Logic)"
echo "============================================"
./func-ux-demo run --builder pack
echo

echo "============================================"
echo "Test Complete!"
echo "============================================"
echo
echo "The prototype demonstrates:"
echo "- Enhanced error messages with actionable guidance"
echo "- Colorized output for better visual hierarchy"
echo "- Workflow-based command organization"
echo "- Smart flag handling and validation"
echo