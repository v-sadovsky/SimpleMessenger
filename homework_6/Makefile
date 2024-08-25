
# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN := $(CURDIR)/bin

# Назначаем GOBIN локальную директорию bin и устанавливаем в нее необходимые плагины
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	"$(CURDIR)/scripts/install_deps.sh"

# Генерация .pb файлов с помощью buf
.buf-generate:
	$(info run buf generate...)
	PATH="$(LOCAL_BIN):$(PATH)" $(LOCAL_BIN)/buf generate

.proto-format:
	$(info run buf format...)
	$(LOCAL_BIN)/buf format -w

# go mod tidy
.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy

.vendor-reset:
	"$(CURDIR)/scripts/vendor_reset.sh"

# Устанавливаем proto описания google/api
.vendor-googleapis:
	"$(CURDIR)/scripts/vendor_obtain_googleapis.sh"

# Устанавливаем proto описания google/protobuf
.vendor-google-protobuf:
	"$(CURDIR)/scripts/vendor_obtain_protobuf.sh"

# Устанавливаем proto описания protoc-gen-openapiv2/options
.vendor-protoc-gen-openapiv2:
	"$(CURDIR)/scripts/vendor_obtain_openapi.sh"

# Устанавливаем proto описания validate
.vendor-protovalidate:
	"$(CURDIR)/scripts/vendor_validate.sh"

.vendor-tidy:
	"$(CURDIR)/scripts/vendor_tidy.sh"

# vendor
vendor:	.vendor-reset .vendor-googleapis .vendor-google-protobuf .vendor-protoc-gen-openapiv2 .vendor-protovalidate .vendor-tidy

# Генерация кода из protobuf
generate: .bin-deps .buf-generate .proto-format .proto-lint .tidy

# Линтер proto файлов
.proto-lint:
	$(LOCAL_BIN)/protolint -config_path ./.protolint.yaml ./proto/

# Линтер
lint: .proto-lint

# Форматирование protobuf файлов
format: .proto-format

## Билд приложения
#build:
#	go build -o $(LOCAL_BIN) ./cmd/profiles/client
#	go build -o $(LOCAL_BIN) ./cmd/profiles/server

swagger-ui:
	"$(CURDIR)/scripts/run_swagger_ui.sh"


# Объявляем, что текущие команды не являются файлами и
# интсрументируем Makefile не искать изменения в файловой системе
.PHONY: \
	.bin-deps \
	.buf-generate \
	.tidy \
	.vendor-protovalidate \
	.proto-lint \
	proto-format \
	vendor \
	generate \
	build \
	lint
