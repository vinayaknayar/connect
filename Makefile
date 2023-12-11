#DEV

build-dev:
	docker build -t connect -f containers/images/Dockerfile . && docker build -t turn -f containers/images/Dockerfile.turn .

run-dev:
	docker-compose -f containers/composes/dc.dev.yml up -d

clean-dev:
	docker-compose -f containers/composes/dc.dev.yml down