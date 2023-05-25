set -euxo pipefail

PATH_PROJECT=$(pwd)

mkdir -p "$(pwd)/functions"
mkdir -p "$(pwd)/functions/register"

# make build
# GOBIN=$(pwd)/functions go install ./...
# chmod +x "$(pwd)"/functions/*

pushd apis/register
go build -o gotrueRegister -ldflags "-X github.com/netlify/gotrue/cmd.Version=`git rev-parse HEAD`"
GOBIN=$PATH_PROJECT/functions/register go install ./...
chmod +x "$PATH_PROJECT"/functions/register/*
popd

ps -aux | grep "go"