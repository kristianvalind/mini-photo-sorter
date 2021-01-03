.PHONY: build

build:
	cd cmd/mps/ && go build -o ../../mps

droplet:
	osacompile -o "MPS Droplet.app" droplet.applescript 