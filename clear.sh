set -euxo pipefail

PATH_PROJECT=$(pwd)


# 定义一个字符串数组
paths=("functions/authorize" "functions/signup" "functions/setting" "functions/token" "functions/verify")

# 遍历数组并输出每个元素
for api in "${paths[@]}"
do
    pushd $api
    api=$(echo "$api" | cut -d'/' -f2)
    rm -f $api
    popd
done