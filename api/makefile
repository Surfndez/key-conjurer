.PHONY: build
SERVERLESS_BUILDS = $(shell ls -d serverless_functions/**/*)

ifndef SERVERLESS_BUILDS
$(error SERVERLESS_BUILDS not set)
else
$(info ${SERVERLESS_BUILDS} will be built)
endif

ifndef TF_WORKSPACE
$(error TF_WORKSPACE is not set)
endif

ifndef AUTHENTICATOR
$(error AUTHENTICATOR is not set)
endif

ifndef MFA
$(error MFA is not set)
endif

ifndef LOGSTASH_ENDPOINT
$(error LOGSTASH_ENDPOINT is not set)
endif

ifndef S3_TF_BUCKET_NAME
$(error S3_TF_BUCKET_NAME is not set)
endif

ifndef SECRETS_RETRIEVER
$(error SECRETS_RETRIEVER is not set)
endif

version = $(shell git rev-parse --short HEAD)
ldflags = "-X github.com/riotgames/key-conjurer/api/consts.Version=$(version)-${TF_WORKSPACE} -X github.com/riotgames/key-conjurer/api/consts.AuthenticatorSelect=$(AUTHENTICATOR) -X github.com/riotgames/key-conjurer/api/consts.MFASelect=$(MFA) -X github.com/riotgames/key-conjurer/api/consts.LogstashEndpoint=$(LOGSTASH_ENDPOINT) -X github.com/riotgames/key-conjurer/api/consts.SettingsRetrieverSelect=$(SECRETS_RETRIEVER)"

deploy:
	make build \
	&& make zip \
	&& make upload

clean:
	rm -rf ../builds

build: $(SERVERLESS_BUILDS)

$(SERVERLESS_BUILDS): build_dir
	$(info building ${@})
	GOOS=linux go build -ldflags $(ldflags) -o ../builds/$(TF_WORKSPACE)/$(subst serverless_functions/,,$@) $@/*.go;

build_dir:
	mkdir -p builds

test:
	go test ./...
