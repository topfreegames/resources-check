TEST_PACKAGES=`find . -type f -name "*.go" ! \( -path "*vendor*" \) | sed -En "s/([^\.])\/.*/\1/p" | uniq`

start:
	@go run main.go start

mocks:
	@mockgen -source=model/interface.go -destination=model/mocks.go -package=model
	@echo 'created mocks in ./model'

test: unit

unit: unit-board clear-coverage-profiles unit-run gather-unit-profiles

clear-coverage-profiles:
	@find . -name '*.coverprofile' -delete

unit-board:
	@echo "\033[1;34m=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\033[0m"
	@echo "\033[1;34m=         Unit Tests         -\033[0m"
	@echo "\033[1;34m=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\033[0m"

unit-run:
	@ginkgo -tags unit -cover -r -randomizeAllSpecs -randomizeSuites -skipMeasurements ${TEST_PACKAGES}

gather-unit-profiles:
	@mkdir -p _build
	@echo "mode: count" > _build/coverage-unit.out
	@bash -c 'for f in $$(find . -name "*.coverprofile"); do tail -n +2 $$f >> _build/coverage-unit.out; done'
