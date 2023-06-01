set -euxo pipefail

PATH_PROJECT=$(pwd)

mkdir -p "$(pwd)/functions"
mkdir -p "$(pwd)/functions/register"
mkdir -p "$(pwd)/functions/settings"

pushd functions/setting
go build -ldflags "-X github.com/netlify/gotrue/cmd.Version=`git rev-parse HEAD`"
GOBIN=$PATH_PROJECT/functions/setting go install ./...
chmod +x "$PATH_PROJECT"/functions/setting/*
popd

ps -aux | grep "go"