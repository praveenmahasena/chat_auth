all: test
all: vet
all: package
all: package_race


test:
	go test ./... -v
vet:
	go vet ./...
package: server

package_race: server_race

server:
	go build -o ./bin/server .

server_race:
	go build --race -o ./bin/server_race .
