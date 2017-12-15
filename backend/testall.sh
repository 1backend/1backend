touch coverage.txt
for pkg in `go list ./...`
do
    go test -v -race -coverprofile=coverage_tmp.txt -covermode=atomic $pkg || ERROR="Error testing $pkg"
    tail -n +2 coverage_tmp.txt >> coverage.txt || die "Unable to append coverage for $pkg"
done