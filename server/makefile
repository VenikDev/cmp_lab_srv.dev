del_all_static:
	cd static && tree /F && del /f /q * && rmdir /s assets

build_client: del_all_static
	cd ../client_v2 && yarn build
	cd ../client_v2/dist \
		&& tree /F \
		&& move * ../../server/static \
		&& move assets ../../server/static

build:
	docker build -t cmp_lab_srv .

run: build
	docker run -it \
		--name cmp_lab_srv \
		-p 80:8080 \
		--link redis:redis cmp_lab_srv

hub: build_client build
	docker tag cmp_lab_srv venikshow/cmp-lab-srv
	docker push venikshow/cmp-lab-srv