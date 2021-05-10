clear
echo "Compile Application...."
echo "Generate MyCloud (Windows APP)..."

GOOS=windows
GOARCH=amd64
go build -ldflags "-s -w" -o dist/result.exe main.go

echo "Generate MyCloud (Linux APP)..."

GOOS=linux
GOARCH=amd64
go build -ldflags "-s -w" -o dist/result main.go

echo "Compilation Complete!"

