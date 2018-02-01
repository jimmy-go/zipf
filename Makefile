
help:
	@echo "Available commands:"
	@echo "\tmake install			Install dependencies."
	@echo "\tmake test			Run tests."
	@echo "\tmake coverage			Show coverage in html."
	@echo "\tmake clean			Clean build files."

install:
	@echo "Make: Install"
	./scripts/install.sh

.PHONY: test
test:
	@echo "Make: Test"
	./scripts/test.sh

coverage:
	@echo "Make: Coverage"
	./scripts/coverage.sh

clean:
	@echo "Make: Clean"
	./scripts/clean.sh
