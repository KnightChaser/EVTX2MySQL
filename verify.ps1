# Verifies the table structure of a MySQL database
# Usage: .\verify.ps1

# Read .env file and load MySQL connection details
$envFile = Get-Content -Path ".env" -Raw
$envVariables = ConvertFrom-StringData $envFile

$server = $envVariables["MYSQL_ACCESS_HOST"]
$database = $envVariables["MYSQL_DATABASE_NAME"]
$username = $envVariables["MYSQL_USERNAME"]
$password = $envVariables["MYSQL_PASSWORD"]

$table = Read-Host "Enter the name of the table(DB: ${database}) to check the latest 10 rows > "
# $query = "SELECT * FROM $table LIMIT 10"
$query = "SELECT * FROM $table LIMIT 10"

$command = "mysql.exe -h $server -u $username -p$password -D $database -e `"$query`""
Invoke-Expression -Command $command