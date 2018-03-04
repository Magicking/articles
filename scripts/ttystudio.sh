#!/bin/zsh
setopt nullglob

ARTICLES_GIT_DIR="${HOME}/source/gocode/src/github.com/Magicking/articles"
RUSH_DIR="${ARTICLES_GIT_DIR}/typescript"
RUSH_PREFIX="ttystudio_frames_rush_"
IDX=(${RUSH_DIR}/${RUSH_PREFIX}*)
if [ x"${1}" != "x" ]; then
  RUSH_NAME="${1}-${IDX}"
else
  # https://stackoverflow.com/questions/2793812/generate-a-random-filename-in-unix-shell
  RUSH_NAME=`cat /dev/urandom | tr -cd 'a-f0-9' | head -c 32`
fi
RUSH_FILE="${RUSH_DIR}/${RUSH_PREFIX}${RUSH_NAME}".json

"${ARTICLES_GIT_DIR}"/scripts/ttystudio/bin/ttystudio --record "${RUSH_FILE}" --interval=100
echo git --git-dir="${ARTICLES_GIT_DIR}/.git/" checkout -b "${RUSH_NAME}"
echo git --git-dir="${ARTICLES_GIT_DIR}/.git/" add ${RUSH_NAME}
echo git --git-dir="${ARTICLES_GIT_DIR}/.git/" commit ${RUSH_NAME} -m "Add ttystudio rush json frames for « ${RUSH_NAME} » attempt ${IDX}"
echo git --git-dir="${ARTICLES_GIT_DIR}/.git/" checkout master
