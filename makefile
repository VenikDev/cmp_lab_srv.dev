del_all_static:
	cd server/static && tree /F && del /f /q * && rmdir /s assets

build_client: del_all_static
	cd client_v2 && yarn build
	cd client_v2/dist && tree /F && move * ../../server/static && move assets ../../server/static

del:
	docker-compose down
# remove all cintainers
	docker compose rm
# remove all images
	docker-compose down --rmi all

deploy: del
	docker-compose up