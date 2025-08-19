@echo off
echo Testing Supabase connection with transaction pooling...
cd test
go mod tidy
go run connection_test.go
cd ..
echo Test completed.
pause
