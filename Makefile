run-kafka: ## run kafka with zookeeper
	docker-compose -f docker-compose.yml up kafka zookeeper -d


stop-kafka: ## stop kafka with zookeeper
	docker-compose -f docker-compose.yml down

logs-kafka: ## logs kafka with zookeeper
	docker-compose -f docker-compose.yml logs -f kafka zookeeper

help: ## print this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
