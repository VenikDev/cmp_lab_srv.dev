del:
	docker-compose down
# remove all cintainers
	docker compose rm
# remove all images
	docker-compose down --rmi all

deploy: del
	docker-compose up --scale srv=3