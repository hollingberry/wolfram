BIN       := wolfram
SOURCES   := main.go $(wildcard **/*.go)

DESTDIR   := /usr/local/bin
MANDIR    := /usr/local/share/man/man1

VERSION   := 0.1.0

GO        := go

MAKEFLAGS := s

all: $(BIN) # man/$(BIN).1 man/$(BIN).1.html

build: $(BIN)

install: $(DESTDIR)/$(BIN) # $(MANDIR)/$(BIN).1

test:
	@$(GO) test ./...

bench:
	@$(GO) test -bench . ./...

clean:
	@$(GO) clean -i ./...
	@-rm -rf $(BIN)
	@-rm -rf $(DESTDIR)/$(BIN)
	@-rm -rf man/$(BIN).1
	@-rm -rf man/$(BIN).1.html
	@-rm -rf $(MANDIR)/$(BIN).1

$(BIN): $(SOURCES)
	@$(GO) build -o $@

$(DESTDIR)/$(BIN): $(BIN)
	@cp $< $@

# man/$(BIN).1: man/$(BIN).1.ronn
# 	@$(RONN) --roff $<
# 
# man/$(BIN).1.html: man/$(BIN).1.ronn
# 	@$(RONN) --html --style=toc $<
# 
# $(MANDIR)/$(BIN).1: man/$(BIN).1
# 	@cp $< $@

# TODO: Make this save the generated .pkg file to dist.
dist/$(BIN).pkg: $(BIN)
	@fpm -s dir \
	     -t osxpkg \
	     --osxpkg-identifier-prefix org.hollingberry \
	     --name $< \
	     --version $(VERSION) \
			 $<

.PHONY: all build install test clean
