clear
echo "Running project tests..."
go test ./...-v
echo "All tests completed!"
sh clean.sh
