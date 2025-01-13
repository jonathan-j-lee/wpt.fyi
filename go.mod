module github.com/web-platform-tests/wpt.fyi

go 1.22.7

toolchain go1.22.9

require (
	cloud.google.com/go/cloudtasks v1.13.3
	cloud.google.com/go/datastore v1.20.0
	cloud.google.com/go/logging v1.13.0
	cloud.google.com/go/secretmanager v1.14.3
	cloud.google.com/go/storage v1.50.0
	github.com/deckarep/golang-set v1.8.0
	github.com/dgryski/go-farm v0.0.0-20240924180020-3414d57e47da
	github.com/gobwas/glob v0.2.3
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gomodule/redigo v1.9.2
	github.com/google/go-github/v65 v65.0.0
	github.com/google/uuid v1.6.0
	github.com/gorilla/handlers v1.5.2
	github.com/gorilla/mux v1.8.1
	github.com/gorilla/securecookie v1.1.2
	github.com/phayes/freeport v0.0.0-20220201140144-74d24b5ae9f5
	github.com/samthor/nicehttp v1.0.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.10.0
	github.com/taskcluster/taskcluster-lib-urls v13.0.1+incompatible
	github.com/taskcluster/taskcluster/v67 v67.1.0
	github.com/tebeka/selenium v0.9.9
	go.uber.org/mock v0.5.0
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/oauth2 v0.25.0
	google.golang.org/api v0.216.0
	google.golang.org/genproto v0.0.0-20241118233622-e639e219e697
	google.golang.org/genproto/googleapis/api v0.0.0-20241209162323-e6fa225c2576
	google.golang.org/grpc v1.69.4
	gopkg.in/yaml.v3 v3.0.1
)

require (
	cel.dev/expr v0.16.2 // indirect
	cloud.google.com/go v0.117.0 // indirect
	cloud.google.com/go/auth v0.13.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.6 // indirect
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	cloud.google.com/go/iam v1.2.2 // indirect
	cloud.google.com/go/longrunning v0.6.2 // indirect
	cloud.google.com/go/monitoring v1.21.2 // indirect
	github.com/BurntSushi/xgb v0.0.0-20200324125942-20f126ea2843 // indirect
	github.com/BurntSushi/xgbutil v0.0.0-20190907113008-ad855c713046 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.25.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.48.1 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.48.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cncf/xds/go v0.0.0-20240905190251-b4127c9b8d78 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/envoyproxy/go-control-plane v0.13.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.14.1 // indirect
	github.com/johncgriffin/overflow v0.0.0-20211019200055-46fa312c352c // indirect
	github.com/orcaman/writerseeker v0.0.0-20200621085525-1d3f536ff85e // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/taskcluster/httpbackoff/v3 v3.1.0 // indirect
	github.com/taskcluster/slugid-go v1.1.0 // indirect
	github.com/tent/hawk-go v0.0.0-20161026210932-d341ea318957 // indirect
	go.opentelemetry.io/contrib/detectors/gcp v1.31.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.55.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0 // indirect
	go.opentelemetry.io/otel v1.31.0 // indirect
	go.opentelemetry.io/otel/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/sdk v1.31.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.31.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.9.0 // indirect
	golang.org/x/tools v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250102185135-69823020774d // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// The project has been moved to GitHub and we don't want to depend on bzr (used by launchpad).
replace launchpad.net/gocheck v0.0.0-20140225173054-000000000087 => gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
