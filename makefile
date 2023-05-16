del_all_static:
	cd server/static && del /f /q * && rmdir /s assets

build_client: del_all_static
	cd client_v2 && yarn build
	cd client_v2/dist && move * ../../server/static && move assets ../../server/static

del:
	docker-compose down
# remove all cintainers
	docker compose rm
# remove all images
	docker rmi comparison_lab-service1
	docker rmi comparison_lab-redis

deploy: build_client del
	docker-compose up