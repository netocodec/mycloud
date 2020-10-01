echo "Cleaning up the mess..."
find . -type f -name '*.db' -delete
find . -type f -name '*.config' -delete
find . -type d -name 'fcloud' -exec rm -rf {} \;
echo "All project cleaned! Ready to startup testing again!"