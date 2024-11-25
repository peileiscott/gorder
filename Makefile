.PHONY: gen
gen: genproto genopenapi

.PHONY: genproto
genproto:
	@./scripts/genproto.sh order

.PHONE: genopenapi
genopenapi:
	@./scripts/genopenapi.sh order