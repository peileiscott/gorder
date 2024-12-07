.PHONY: gen
gen: proto openapi

.PHONY: proto
proto:
	@./scripts/proto.sh order

.PHONY: openapi
openapi:
	@./scripts/openapi.sh order