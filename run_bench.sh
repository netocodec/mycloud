clear
echo "Running project benchmarks..."
go test ./... -bench=. -v
echo "All benchmarks completed!"
sh clean.sh