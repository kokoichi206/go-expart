.DEFAULT_GOAL = run
run: build
	@echo RUN
	@sudo ip netns exec router1 ./router --mode ch1

build:
	@echo BUILD
	@go build .
