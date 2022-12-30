TEST := $(wildcard test/*.sh)

build:
	go build

test: build
	$(MAKE) $(TEST)
	@printf '\e[32mpassed all tests!\e[0m\n' 

$(TEST):
	@echo 'testing' $@
	@./$@
	@printf '\e[32mOk!\e[0m\n'

clean:
	go clean
	rm -rf out/

.PHONY: build clean test $(TEST)