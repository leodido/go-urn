.PHONY = clean

urn: urn.peg.go main.go
	go build -o urn

urn.peg.go: urn.peg
	@peg -switch -inline $?

clean:
	rm -rf urn *.peg.go