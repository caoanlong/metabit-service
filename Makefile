GOOS = linux # linux/windows/darwin
GOARCH = amd64 # amd64/arm64

Host="root@45.32.250.46"
HostTest="root@64.176.38.50"
Cert="~/.ssh/id_rsa"
Directory=/home/backend/metabit

build: compile upload
build-test: compile upload-test

compile:
	go mod tidy
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./bin/metabit_app app/main.go
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./bin/metabit_admin admin/main.go

upload: upload-app upload-admin
upload-test: upload-app-test upload-admin-test

upload-app:
	scp -i $(Cert) ./app/conf.yaml $(Host):$(Directory)/metabit_app/conf.yaml
	scp -i $(Cert) ./app/conf.prod.yaml $(Host):$(Directory)/metabit_app/conf.prod.yaml
	scp -i $(Cert) ./bin/metabit_app $(Host):$(Directory)/metabit_app/metabit_app_new
	ssh -i $(Cert) $(Host) "sudo supervisorctl stop metabit-app"
	ssh -i $(Cert) $(Host) "cd /home/backend/metabit/metabit_app; rm -f ./metabit_app; mv metabit_app_new metabit_app; chmod +x metabit_app"
	ssh -i $(Cert) $(Host) "sudo supervisorctl start metabit-app"

upload-admin:
	scp -i $(Cert) ./admin/conf.yaml $(Host):$(Directory)/metabit_admin/conf.yaml
	scp -i $(Cert) ./admin/conf.prod.yaml $(Host):$(Directory)/metabit_admin/conf.prod.yaml
	scp -i $(Cert) ./bin/metabit_admin $(Host):$(Directory)/metabit_admin/metabit_admin_new
	ssh -i $(Cert) $(Host) "sudo supervisorctl stop metabit-admin"
	ssh -i $(Cert) $(Host) "cd /home/backend/metabit/metabit_admin; rm -f ./metabit_admin; mv metabit_admin_new metabit_admin; chmod +x metabit_admin"
	ssh -i $(Cert) $(Host) "sudo supervisorctl start metabit-admin"

upload-app-test:
	scp -i $(Cert) ./app/conf.yaml $(HostTest):$(Directory)/metabit_app/conf.yaml
	scp -i $(Cert) ./app/conf.test.yaml $(HostTest):$(Directory)/metabit_app/conf.test.yaml
	scp -i $(Cert) ./bin/metabit_app $(HostTest):$(Directory)/metabit_app/metabit_app_new
	ssh -i $(Cert) $(HostTest) "sudo supervisorctl stop metabit-app"
	ssh -i $(Cert) $(HostTest) "cd /home/backend/metabit/metabit_app; rm -f ./metabit_app; mv metabit_app_new metabit_app; chmod +x metabit_app"
	ssh -i $(Cert) $(HostTest) "sudo supervisorctl start metabit-app"

upload-admin-test:
	scp -i $(Cert) ./admin/conf.yaml $(HostTest):$(Directory)/metabit_admin/conf.yaml
	scp -i $(Cert) ./admin/conf.test.yaml $(HostTest):$(Directory)/metabit_admin/conf.test.yaml
	scp -i $(Cert) ./bin/metabit_admin $(HostTest):$(Directory)/metabit_admin/metabit_admin_new
	ssh -i $(Cert) $(HostTest) "sudo supervisorctl stop metabit-admin"
	ssh -i $(Cert) $(HostTest) "cd /home/backend/metabit/metabit_admin; rm -f ./metabit_admin; mv metabit_admin_new metabit_admin; chmod +x metabit_admin"
	ssh -i $(Cert) $(HostTest) "sudo supervisorctl start metabit-admin"
