.PHONY: fmt
fmt:
	find . -name "*.go" -type f -not -path "./vendor/*" | xargs gofmt -s -w

.PHONY: fmt-check
fmt-check:
	@if git diff --quiet --exit-code; then \
		$(MAKE) fmt && git diff --exit-code || { \
			git checkout .; \
			echo; \
			echo "Please run 'make fmt' and commit the result"; \
			echo; \
			false; } >&2; \
	else { \
		echo; \
		echo "'make fmt-check' cannot be run with unstaged changes"; \
		echo; \
		false; } >&2; \
	fi
