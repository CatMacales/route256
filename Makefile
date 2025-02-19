LOCAL_BIN:=$(CURDIR)/bin
CART_DIR:=$(CURDIR)/cart
LOMS_DIR:=$(CURDIR)/loms

.PHONY: run-all
run-all:
	cd cart && make run-cart

.PHONY: bin-deps
bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.4
	GOBIN=$(LOCAL_BIN) go install github.com/fzipp/gocyclo/cmd/gocyclo@v0.6.0
	GOBIN=$(LOCAL_BIN) go install github.com/uudashr/gocognit/cmd/gocognit@v1.2.0

.PHONY: test-coverage
test-coverage:
	@echo "Cart coverage"
	cd $(CART_DIR) && make test-coverage
	@echo "LOMS coverage"
	cd $(LOMS_DIR) && make test-coverage

.PHONY: cyclo-report
cyclo-report:
	@echo "Cart report"
	cd $(CART_DIR) && make cyclo-report
	@echo "LOMS report"
	cd $(LOMS_DIR) && make cyclo-report

.PHONY: cognit-report
cognit-report:
	@echo "Cart report"
	cd $(CART_DIR) && make cognit-report
	@echo "LOMS report"
	cd $(LOMS_DIR) && make cognit-report