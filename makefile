SHELL = /bin/bash

.PHONY = clean

directory := grammar

$(directory)/urn_{parser,lexer,base_listener,listener,base_visitor,visitor}.go: $(directory)/Urn.g4
	$(SHELL) -c "./$(directory)/antlr -Dlanguage=Go -o $(directory) -package $(directory) -listener -no-visitor $?"

test: $(directory)/urn_{parser,lexer,base_listener,listener,base_visitor,visitor}.go *_test.go
	go test

clean:
	rm -f $(directory)/urn_{parser,lexer,base_listener,listener,base_visitor,visitor}.go $(directory)/Urn{,Lexer}.tokens