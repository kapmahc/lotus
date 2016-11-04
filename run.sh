#!/bin/sh
while true; do
  go run main.go s &	
  inotifywait -e modify -e move -e create -e delete -e attrib -r `pwd`
  fuser -k 8080/tcp
done
