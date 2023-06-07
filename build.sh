set -euxo pipefail

PATH_PROJECT=$(pwd)

mkdir -p "$(pwd)/functions"

# 定义一个字符串数组
paths=("functions/authorize" "functions/signup" "functions/setting" "functions/settings" "functions/token" "functions/verify")

# 遍历数组并输出每个元素
for api in "${paths[@]}"
do
    mkdir -p "$(pwd)/functions"/"$api"
    pushd $api
    go build -ldflags "-X github.com/netlify/gotrue/cmd.Version=`git rev-parse HEAD`"
    GOBIN=$PATH_PROJECT/$api go install ./...
    chmod +x "$PATH_PROJECT"/"$api"/*
    popd
done


ps -aux | grep "go"