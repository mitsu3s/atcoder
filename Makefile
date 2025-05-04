.PHONY: create

create:
	@mkdir -p $(ID)
	@for dir in a b c d; do \
		mkdir -p $(ID)/$$dir; \
		touch $(ID)/$$dir/$$dir.go; \
		touch $(ID)/$$dir/input.txt; \
	done
	@echo "Created directory structure for $(ID)"
