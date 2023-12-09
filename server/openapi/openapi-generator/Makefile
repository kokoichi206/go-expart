current_dir := $(shell pwd)

.PHONY: gen
gen: clear
	docker run --rm -v $(current_dir):/project openapitools/openapi-generator-cli:v7.0.0 \
		generate -i /project/openapi.yaml -g go \
		--additional-properties packageName=component -o /project/gen/component -p enumClassPrefix=true

.PHONY: clear
clear:
	ls gen/component | xargs -n1 | grep -v ".openapi-generator" | xargs -I{} rm -r "gen/component/{}"

.PHONY: genm
genm: clearm
	docker run --rm -v $(current_dir):/project openapitools/openapi-generator-cli:v7.0.0 \
		generate -i /project/openapi.yaml -g go \
		--additional-properties packageName=component_method -o /project/gen/component_method -p enumClassPrefix=true

.PHONY: clearm
clearm:
	ls gen/component_method | xargs -n1 | grep -v ".openapi-generator" | xargs -I{} rm -r "gen/component_method/{}"

.PHONY: gen710
gen710: clear710
	docker run --rm -v $(current_dir):/project openapitools/openapi-generator-cli:v7.1.0 \
		generate -i /project/openapi.yaml -g go \
		--additional-properties generateMarshalJSON=false \
		--additional-properties disallowAdditionalPropertiesIfNotPresent=false \
		--additional-properties packageName=component_710 -o /project/gen/component_710 -p enumClassPrefix=true

.PHONY: clear710
clear710:
	ls gen/component_710 | xargs -n1 | grep -v ".openapi-generator" | xargs -I{} rm -r "gen/component_710/{}"

.PHONY: bench
bench:
	go test -bench .
