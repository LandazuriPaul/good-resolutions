#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
app="good-resolutions-api"
src="$srcPath/$app/$pkgFile"

echo "Live reload server launched (with realize)..."

if [ "$1" == "" ]; then
  mode="development"
  printf "\nNo mode given"
else
  mode="$1"
fi
printf "\n> Server running in $mode mode"
printf "\n> Loading $mode.env configuration file...\n"

if [ "$mode" != "development" ] && [ "$mode" != "production" ]; then
  printf "\n/!\ The running mode has to be either \"development\" or \"production\"\n"
  exit 1
fi

printf "\nStart running: $app\n"
# Set all ENV vars for the server to run
export $(grep -v '^#' $mode.env | xargs) && time $GOPATH/bin/realize start --run $src
# This should unset all the ENV vars, just in case.
unset $(grep -v '^#' $mode.env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped running: $app\n\n"
