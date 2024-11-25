.PHONY: gen
gen: genproto genopenapi

.PHONY: genproto
genproto:
	@./scripts/genproto.sh order
	@./scripts/genproto.sh stock

.PHONE: genopenapi
genopenapi:
	@./scripts/genopenapi.sh order