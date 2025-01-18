export GOEXPERIMENT=noboringcrypto
export CGO_ENABLED=0

# This controls the directory the built artifacts go into
export ARTIFACT_DIR=artifacts/
export GOOS=linux
make cloudflared

