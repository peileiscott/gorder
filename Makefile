.PHONY: gen
gen: proto

.PHONY: proto
proto:
	@./scripts/proto.sh order