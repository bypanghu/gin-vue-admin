
VERSION_FILE=version
CURRENT_VERSION=$(shell cat $(VERSION_FILE))
IMAGE="crpi-1hrxw3c4eln7c2uw.cn-shanghai.personal.cr.aliyuncs.com/bypanghu/gva-frontend"


build:
	$(eval NEW_VERSION=$(shell echo $(CURRENT_VERSION) | awk -F. '{print $$1 "." $$2 "." $$3+1}'))
	@echo "更新版本号为：$(NEW_VERSION)"
	@docker build --platform linux/amd64 -t "$(IMAGE):$(NEW_VERSION)" .
	@echo "已完成 build"
	@docker push "$(IMAGE):$(NEW_VERSION)"
	@echo $(NEW_VERSION) > $(VERSION_FILE)
	@echo "已完成 push 到镜像仓库"
