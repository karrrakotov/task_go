build:
	docker build -t happydayaway/task_img .
run:
	docker run -d -p 80:8080 --rm --name task_cont happydayaway/task_img
stop-img:
	docker stop happydayaway/task_img
stop-cont:
	docker stop task_cont
rmi:
	docker rmi happydayaway/task_img