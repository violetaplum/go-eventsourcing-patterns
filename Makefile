# Makefile

# 변수 설정
APP_NAME := account-service
DOCKER_COMPOSE_FILE := deployments/docker-compose.yml

.PHONY: help
help: ## 사용 가능한 명령어 표시
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## 도커 이미지 빌드
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

.PHONY: up
up: ## 컨테이너 실행
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

.PHONY: down
down: ## 컨테이너 중지, 볼륨삭제
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v




