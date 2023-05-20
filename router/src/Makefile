.DEFAULT_GOAL = run
run: build
	@echo RUN
	@sudo ip netns exec router1 ./router --mode ch4

build:
	@echo BUILD
	@go build .
