if command -v go &> /dev/null
then 
rm test/output.txt
go run . test/input.txt test/output.txt assets/airport-lookup.csv
else 
echo "you need to install go, before running the app"
fi
