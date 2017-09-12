PHONY: update-deps install-deps fmt lint golint test test-with-coverage build
# TODO: When Go 1.9 is released vendor folder should be ignored automatically
HOST=ec2-user@test.wanliu.biz
PACKAGES=`go list ./... | grep -v vendor | grep -v mocks`
TIMESTAMP=$(shell date +%s)
APPNAME=BlueGreenDemoApplication
GROUPNAME=BlueGreenDemoFleet-i3ggrub
PUSH_SCRIPT=aws deploy push --application-name $(APPNAME) \
	--s3-location s3://jiejie-deploy/oauth2-server-$(TIMESTAMP).zip \
	--source ./build
AWSPUSH=$(shell $(PUSH_SCRIPT))

# sed 's/<deployment-group-name>.*//'
# aws deploy create-deployment --application-name Jiejie --s3-location bucket=jiejie-deploy,key=oauth2-server-1504866286.zip,bundleType=zip,eTag=0521449d68f8b6b0109b060e0ece04b9 --deployment-group-name <deployment-group-name> --deployment-config-name <deployment-config-name> --description <description>

update-deps:
	rm -rf Godeps
	rm -rf vendor
	go get github.com/tools/godep
	godep save ./...

install-deps:
	go get github.com/tools/godep
	godep restore

fmt:
	for pkg in ${PACKAGES}; do \
		go fmt $$pkg; \
	done;

lint:
	gometalinter --disable-all -E vet -E gofmt -E misspell -E ineffassign -E goimports -E deadcode --tests --vendor ./...

golint:
	for pkg in ${PACKAGES}; do \
		golint $$pkg; \
	done;

test:
	for pkg in ${PACKAGES}; do \
		go test $$pkg; \
	done;

test-with-coverage:
	echo "" > coverage.out
	echo "mode: set" > coverage-all.out
	for pkg in ${PACKAGES}; do \
		go test -coverprofile=coverage.out -covermode=set $$pkg; \
		tail -n +2 coverage.out >> coverage-all.out; \
	done;
	#go tool cover -html=coverage-all.out

build-linux:
	@go generate ./web
	@docker run -it --rm -v $(shell pwd):/go/src/github.com/wanliu/go-oauth2-server --entrypoint "go" go-oauth2-server:latest build -o build/oauth2-server-linux .

build-docker:
	@docker build --rm -t go-oauth2-server:latest .

clean-deploy: oauth2-server.zip
	@rm oauth2-server.zip

deploy-cd: build/oauth2-server-linux
	@cp appspec.yml build/
	@cp -r scripts build/
	$(eval RETURN := $(AWSPUSH))
	@echo "$(RETURN)" | sed 's/To deploy with this revision, run: //' | sed 's/<deployment-group-name>.*/$(GROUPNAME)/' | sh

deploy: build/oauth2-server-linux
	@-rm build/oauth2-server.zip
	@cp -r scripts build/
	@mkdir -p build/public/; cp -r public/css build/public
	@cp -r public/img build/public
	@cp public/favicon.ico build/public
	@cd build; zip -r oauth2-server.zip .
	@ansible-playbook -i ansible/hosts ansible/site.yml
