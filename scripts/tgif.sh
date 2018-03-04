#!/bin/zsh
setopt nullglob

ARTICLES_REPO="${HOME}/source/gocode/src/github.com/Magicking/articles"
RUSH_DIR="${HOME}/source/gocode/src/github.com/Magicking/typescript_rushs"
RUSH_PREFIX="ttystudio_frames_rush_"
IDX=(${RUSH_DIR}/${RUSH_PREFIX}*)
if [ x"${1}" != "x" ]; then
  RUSH_NAME="${1}-${IDX}"
  RUSH_BRANCH="${1}"
else
  # https://stackoverflow.com/questions/2793812/generate-a-random-filename-in-unix-shell
  RUSH_NAME=`cat /dev/urandom | tr -cd 'a-f0-9' | head -c 32`
  RUSH_BRANCH="${RUSH_NAME}"
fi
RUSH_FILE="${RUSH_DIR}/${RUSH_PREFIX}${RUSH_NAME}".json

"${ARTICLES_REPO}"/scripts/ttystudio/bin/ttystudio --record "${RUSH_FILE}" --interval=100
PWD=`pwd`
cd "${RUSH_DIR}"
git checkout -b "${RUSH_PREFIX}${RUSH_BRANCH}"
git add ${RUSH_FILE}
git commit ${RUSH_FILE} -m "Add ttystudio rush json frames for « ${RUSH_NAME} » attempt ${IDX}"
git push --set-upstream origin "${RUSH_PREFIX}${RUSH_BRANCH}"&
git checkout master
cd "${PWD}"
