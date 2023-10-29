current_dir := $(shell pwd)

.PHONY: gen
gen: clear
	docker run --rm -v $(current_dir):/project openapitools/openapi-generator-cli:v7.0.0 \
		generate -i /project/openapi.yaml -g go \
		--additional-properties=packageName=component -o /project/gen/component -p enumClassPrefix=true

.PHONY: clear
clear:
	ls gen/component | xargs -n1 | grep -v ".openapi-generator" | xargs -I{} rm -r "gen/component/{}"
