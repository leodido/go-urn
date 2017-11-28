SHELL=/bin/bash

.PHONY = clean

directory := grammar

# Get the path of this makefile.
# Place this before include directives, if any.
this := $(lastword $(MAKEFILE_LIST))

$(directory)/urn_{parser,lexer,base_listener,listener,base_visitor,visitor}.go: $(directory)/Urn.g4 $(directory)/antlr
	$(SHELL) -c "./$(directory)/antlr -Dlanguage=Go -o $(directory) -package $(directory) -listener -no-visitor $<"

$(directory)/antlr:
	@docker pull leodido/antlr
	@docker create --name antlr leodido/antlr
	@docker cp antlr:antlr $(directory)
	@docker rm antlr

test: tcmd = "go test -v"
ifdef COVERAGE
test: tcmd = "go test -v -coverprofile=cov.out"
endif
test: $(directory)/urn_{parser,lexer,base_listener,listener,base_visitor,visitor}.go *_test.go 
	$(SHELL) -c $(tcmd)

cov.out: 
	$(MAKE) -f $(this) test COVERAGE=yes

coverage: cov.out
	@go tool cover -func=$<

clean:
	rm -f $(directory)/urn_{parser,lexer,base_listener,listener,base_visitor,visitor}.go $(directory)/Urn{,Lexer}.tokens $(directory)/antlr *.out