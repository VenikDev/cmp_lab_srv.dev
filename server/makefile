build:
	docker build -t cmp_lab_srv .

run: build
	docker run -it --name cmp_lab_srv -p 8080:8080 cmp_lab_srv

hub: build
	docker tag cmp_lab_srv venikshow/cmp-lab-srv
	docker push venikshow/cmp-lab-srv
