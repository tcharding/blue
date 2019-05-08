PROJ=github.com/tcharding/blue
EXE=serve

all: ${EXE}

serve:
	go build -o bin/${EXE} ${PROJ}/cmd/${EXE}

test:
	go test ./...

clean:
	rm -rf bin/${EXE}
