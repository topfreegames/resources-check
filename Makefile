start:
	@go run main.go start

mocks:
	@mockgen -source=models/interface.go -destination=models/mocks.go -package=models
	@echo 'created mocks in ./models'
