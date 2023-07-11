build: clean
	@go build -o output/proxyserver

clean:
	@rm -f output/proxyserver

run: build
	@output/proxyserver