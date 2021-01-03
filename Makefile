.PHONY: clean build droplet

clean:
	rm -fr mps "MPS Droplet.app"

build: clean
	cd cmd/mps/ && go build -o ../../mps

droplet: build
	osacompile -o "MPS Droplet.app" droplet.applescript 
	cp mps "MPS Droplet.app/Contents/MacOS/mps"