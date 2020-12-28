env GOOS=darwin GOARCH=amd64 go build
mv pally-split ../mac/
env GOOS=linux GOARCH=amd64 go build
mv pally-split ../linux/
env GOOS=windows GOARCH=amd64 go build
mv pally-split.exe ../win/
