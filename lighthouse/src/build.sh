env GOOS=darwin GOARCH=amd64 go build
mv lighthouse-batch-check ../mac/
env GOOS=linux GOARCH=amd64 go build
mv lighthouse-batch-check ../linux/
env GOOS=windows GOARCH=amd64 go build
mv lighthouse-batch-check.exe ../win/
