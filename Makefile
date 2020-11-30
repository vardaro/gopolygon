default:
	go test `go list ./... | grep -v examples` -count=1
