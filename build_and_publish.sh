#!/bin/bash
set -e

REGION="us-west-1"
ACCOUNT_ID=""
APP_NAME="bunny_app"
IMAGE_TAG="latest"
LOG_FILE="/tmp/bunny_app_build_and_publish.log"

usage() {
    echo "Usage: $0 -a <account_id> [-r <region>]"
    echo "  -a  AWS Account ID (required)"
    echo "  -r  AWS Region (defaults to us-west-1)"
    exit 1
}

while getopts "a:r:h" opt; do
  case ${opt} in
    a ) ACCOUNT_ID=$OPTARG ;;
    r ) REGION=$OPTARG ;;
    h ) usage ;;
    \? ) usage ;;
  esac
done

if [ -z "$ACCOUNT_ID" ]; then
    echo "Error: AWS Account ID is required."
    usage
fi

ECR_REGISTRY="${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com"
LOCAL_IMAGE="${APP_NAME}:${IMAGE_TAG}"
REMOTE_IMAGE="${ECR_REGISTRY}/${APP_NAME}:${IMAGE_TAG}"

rm -f "$LOG_FILE"

render_ui() {
    local current=$1
    local total=$2
    local message=$3
    
    local percent=$(( current * 100 / total ))
    local bar_width=20
    local filled=$(( current * bar_width / total ))
    local unfilled=$(( bar_width - filled ))
    
    local bar=$(printf "%${filled}s" | tr ' ' '█')
    local spaces=$(printf "%${unfilled}s")
    
    if [ "$current" -gt 0 ] || [ "$TOTAL_STEPS" -eq 5 ]; then
        printf "\033[1A\033[K" 
        printf "\033[1A\033[K" 
    fi
    
    printf "Progress: [%s%s] %d%%\n" "$bar" "$spaces" "$percent"
    printf "Current:  \033[36m%-50s\033[0m\n" "$message"
}

TOTAL_STEPS=4
START_STEP=1

echo "Initializing build environment..."
printf "\n\n"

if ! docker info >/dev/null 2>&1; then
    if command -v colima >/dev/null 2>&1; then
        TOTAL_STEPS=5
        START_STEP=2
        render_ui 1 5 "Booting Colima Virtual Machine..."
        colima start >> "$LOG_FILE" 2>&1
    else
        echo "Error: Docker daemon is unreachable and Colima is not installed."
        exit 1
    fi
fi

STEP=$START_STEP
render_ui $STEP $TOTAL_STEPS "Authenticating with AWS ECR Registry..."
aws ecr get-login-password --region "${REGION}" 2>> "$LOG_FILE" | docker login --username AWS --password-stdin "${ECR_REGISTRY}" >> "$LOG_FILE" 2>&1

STEP=$((STEP + 1))
render_ui $STEP $TOTAL_STEPS "Building Docker image (compiling Go binary)..."
(
    while kill -0 $$ 2>/dev/null; do
        sleep 1
        render_ui $((STEP - 1)) $TOTAL_STEPS "Building: running go build compiler tags..."
        sleep 1
        render_ui $((STEP - 1)) $TOTAL_STEPS "Building: assembling final scratch layers..."
        break
    done
) &
LOOP_PID=$!
docker build -t "${LOCAL_IMAGE}" . >> "$LOG_FILE" 2>&1
kill $LOOP_PID 2>/dev/null || true
render_ui $STEP $TOTAL_STEPS "Image successfully compiled locally."

STEP=$((STEP + 1))
render_ui $STEP $TOTAL_STEPS "Applying ECR registry tags..."
docker tag "${LOCAL_IMAGE}" "${REMOTE_IMAGE}" >> "$LOG_FILE" 2>&1

STEP=$((STEP + 1))
render_ui $STEP $TOTAL_STEPS "Pushing image layers to ECR (uploading to AWS)..."
docker push "${REMOTE_IMAGE}" >> "$LOG_FILE" 2>&1

render_ui $TOTAL_STEPS $TOTAL_STEPS "Deployment complete!"
echo "--------------------------------------------------"
echo "Success! Image pushed to: ${REMOTE_IMAGE}"
echo "Full command logs saved to: $LOG_FILE"