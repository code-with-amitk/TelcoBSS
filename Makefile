.PHONY: all build-cpp clean

all: build-cpp

build-cpp:
	cmake -S cmd/rating-engine -B cmd/rating-engine/build
	cmake --build cmd/rating-engine/build

clean:
	rm -rf cmd/rating-engine/build
