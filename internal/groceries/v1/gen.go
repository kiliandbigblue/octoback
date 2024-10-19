package v1

//go:generate rm -rf mocks
//go:generate mkdir mocks
//go:generate mockery --disable-version-string --dir ../../store/v1 --name Querier --structname Querier --output mocks

