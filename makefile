del:
	docker-compose down
# remove all cintainers
	docker compose rm
# remove all images
	docker-compose down --rmi all

deploy: del
	docker-compose create --scale srv=$(SCALE)
	docker-compose start