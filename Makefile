build: 
	@go build -o bin/BankingApi
run: build
	@./bin/BankingApi