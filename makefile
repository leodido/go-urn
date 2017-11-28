SHELL = /bin/bash

.PHONY = clean

# Get the path of this makefile.
# Place this before include directives, if any.
this := $(lastword $(MAKEFILE_LIST))

directory := grammar

$(directory)/urn_{parser,lexer,base_listener,listener,base_visitor,visitor}.go: $(directory)/Urn.g4 $(directory)/antlr
	$(SHELL) -c "./$(directory)/antlr -Dlanguage=Go -o $(directory) -package $(directory) -listener -no-visitor $<"

test: $(directory)/urn_{parser,lexer,base_listener,listener,base_visitor,visitor}.go *_test.go
	go test -v

$(directory)/antlr:
	@docker pull leodido/antlr
	@docker create --name antlr leodido/antlr
	@docker cp antlr:antlr $(directory)
	@docker rm antlr

clean:
	rm -f $(directory)/urn_{parser,lexer,base_listener,listener,base_visitor,visitor}.go $(directory)/Urn{,Lexer}.tokens $(directory)/antlr