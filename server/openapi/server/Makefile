current_dir := $(shell pwd)

.PHONY: gen
gen: clear
	docker run --rm -v $(current_dir):/project openapitools/openapi-generator-cli:v7.0.0 \
		generate -i /project/components.yaml -g go \
		--additional-properties=packageName=component -o /project/gen/component -p enumClassPrefix=true

.PHONY: clear
clear:
	ls gen/component | xargs -n1 | grep -v ".openapi-generator" | xargs -I{} rm -r "gen/component/{}"

.PHONY: serve
serve:
	go run main.go

.PHONY: gen_method
gen_method: clear
	docker run --rm -v $(current_dir):/project openapitools/openapi-generator-cli:v7.0.0 \
		generate -i /project/components.yaml -g go \
		--additional-properties=packageName=component_method -o /project/gen/component_method -p enumClassPrefix=true

.PHONY: clear_method
clear_method:
	ls gen/component_method | xargs -n1 | grep -v ".openapi-generator" | xargs -I{} rm -r "gen/component_method/{}"

.PHONY: bench
bench:
	go test -bench .
