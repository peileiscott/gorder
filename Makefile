.PHONY: gen
gen: proto openapi

.PHONY: proto
proto:
	@./scripts/proto.sh order
	@./scripts/proto.sh stock

.PHONY: openapi
openapi:
	@./scripts/openapi.sh order