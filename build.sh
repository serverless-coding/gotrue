set -euxo pipefail

mkdir -p "$(pwd)/functions"
make build
GOBIN=$(pwd)/functions go install ./...
chmod +x "$(pwd)"/functions/*