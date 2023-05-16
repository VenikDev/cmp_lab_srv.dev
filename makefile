del_all_static:
	cd server/static && del /f /q * && rmdir /s assets

build_client: del_all_static
	cd client_v2 && yarn build
	cd client_v2/dist && move * ../../server/static && move assets ../../server/static

run_srv: build_client
	docker-compose up