.PHONY: all install clean format lint test build run

all: clean build

install:
	# Installing Moonrepo proto https://moonrepo.dev/proto
	@if [ "$$(uname)" = "Linux" ] || [ "$$(uname)" = "Darwin" ]; then \
		bash -c "$$(curl -fsSL https://moonrepo.dev/install/proto.sh)"; \
	elif [ "$$(uname -r | grep -i microsoft)" != "" ]; then \
		bash -c "$$(curl -fsSL https://moonrepo.dev/install/proto.sh)"; \
	elif [ "$$(uname -s | grep -i mingw)" != "" ] || [ "$$(uname -s | grep -i cygwin)" != "" ] || [ "$$(uname -s | grep -i windows)" != "" ]; then \
		powershell -Command "irm https://moonrepo.dev/install/proto.ps1 | iex"; \
		powershell -Command "Set-ExecutionPolicy -Scope CurrentUser RemoteSigned"; \
	else \
		echo "Unsupported operating system. Please install manually:"; \
		exit 1; \
	fi;
	# Installing all the other tools with the proto
	proto use
	# Running the Moon setup task at the end
	moon :setup

clean:
	moon :clean

format:
	moon :format

lint:
	moon :lint

test:
	moon :test

build:
	moon :build

run:
	moon :run
