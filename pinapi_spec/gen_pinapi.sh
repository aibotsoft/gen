#!/bin/sh

SCRIPT="$0"
echo "# START SCRIPT: $SCRIPT"

executable="./../openapi-generator-cli.jar"

SPEC="./spec/merged/pinapi.yaml"
GENERATOR="go-experimental"
#GENERATOR="go"
STUB_DIR="./../pinapi"
USER_AGENT="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36"

echo "Removing files and folders under $STUB_DIR"
rm -rf $STUB_DIR/{*.go,*.sh,*.md,*.mod,*.sum}

java -jar $executable generate -i $SPEC -g $GENERATOR -o $STUB_DIR --type-mappings integer=int  --http-user-agent "$USER_AGENT"

# --additional-properties --generate-alias-as-model true
echo "Removing mod files"
rm -rf $STUB_DIR/api
#rm -rf $STUB_DIR/go.mod
#rm -rf $STUB_DIR/go.sum
#rm -rf $STUB_DIR/git_push.sh
#rm -rf $STUB_DIR/readme.md
#rm -rf $STUB_DIR/.travis.yml
#rm -rf $STUB_DIR/.gitignore
#rm -rf $STUB_DIR/docs
#rm -rf $STUB_DIR/.openapi-generator
