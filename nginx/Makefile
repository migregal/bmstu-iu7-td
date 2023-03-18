.PHONY: generate-crypto
generate-crypto: OUT:=./out/crypto
generate-crypto:
	@mkdir -p $(OUT)
	openssl req -x509 -out $(OUT)/cert.pem -keyout $(OUT)/key.pem -newkey rsa:2048 -nodes -sha256 -subj '/CN=localhost' -extensions EXT -config ./cnf/cert.cnf
