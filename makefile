SHELL := /bin/bash

# ==============================================================================
# Local support

up:
	go run app/services/department-api/main.go -race

docker-up:
	docker compose -f zarf/docker/docker-compose.yaml up -d

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -v ./...
	go mod tidy
	go mod vendor


# ==============================================================================
# Metrics and Tracing

metrics-view:
	expvarmon -ports="localhost:4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"

# test-endpoint:
# 	curl -il localhost:4000/debug/vars

test-endpoint:
	curl -il localhost:3000/test

test-endpoint-auth:
	curl -il -H "Authorization: Bearer ${TOKEN}" localhost:3000/test/auth

liveness:
	curl -il http://localhost:4000/debug/liveness

readiness:
	curl -il http://localhost:4000/debug/readiness

pgcli:
	pgcli postgresql://postgres:postgres@localhost

migrate:
	go run app/tooling/admin/main.go

query:
	@curl -s "http://localhost:3000/users?page=1&rows=2&orderBy=name,ASC"

# ==============================================================================
# Running tests within the local computer
# go install honnef.co/go/tools/cmd/staticcheck@latest
# go install golang.org/x/vuln/cmd/govulncheck@latest

test:
	go test -count=1 ./...
	staticcheck -checks=all ./...
	govulncheck ./...

