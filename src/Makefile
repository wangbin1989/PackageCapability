BINARY_DIR := bin
SRC_DIR := cmd

clean:
	@echo "Cleaning up..."
	@rm -rf $(BINARY_DIR)

build:
	@mkdir -p $(BINARY_DIR)
	@for dir in $(SRC_DIR)/*; do \
		if [ -d $$dir ]; then \
			go build -o $(BINARY_DIR)/$$(basename $$dir) ./$$dir; \
		fi; \
	done
