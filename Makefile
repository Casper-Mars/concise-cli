
.PHONY: test
test:
	make clean
	mkdir out
	go build -o out/cli main.go
	cd out \
	&& \
	./cli -m test-cli -p 1.0.1

.PHONY: clean
clean:
	rm -rf out