mockgen -destination=application/mocks/application.go -source=application/product.go application


cobra add cli
go run ./main.go cli
