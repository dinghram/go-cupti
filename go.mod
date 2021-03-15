module github.com/c3sr/go-cupti

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1

replace github.com/uber/jaeger => github.com/jaegertracing/jaeger v1.22.0

replace github.com/jaegertracing/jaeger => github.com/uber/jaeger v1.22.0

require (
	github.com/GeertJohan/go-sourcepath v0.0.0-20150925135350-83e8b8723a9b
	github.com/benesch/cgosymbolizer v0.0.0-20190515212042-bec6fe6e597b
	github.com/c3sr/config v1.0.1
	github.com/c3sr/logger v1.0.1
	github.com/c3sr/nvidia-smi v1.0.0
	github.com/c3sr/tracer v1.0.0
	github.com/c3sr/utils v1.0.0
	github.com/c3sr/vipertags v1.0.0
	github.com/fatih/color v1.10.0
	github.com/iancoleman/strcase v0.1.3
	github.com/ianlancetaylor/cgosymbolizer v0.0.0-20210303021718-7cb7081f6b86 // indirect
	github.com/ianlancetaylor/demangle v0.0.0-20210309225245-94ab485c4a6d
	github.com/k0kubun/pp/v3 v3.0.7
	github.com/mailru/easyjson v0.7.7
	github.com/mitchellh/go-homedir v1.1.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/rainycape/dl v0.0.0-20151222075243-1b01514224a1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.3.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/unknwon/com v1.0.1
	gitlab.com/NebulousLabs/fastrand v0.0.0-20181126182046-603482d69e40
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/examples v0.0.0-20210312231957-21976fa3e38a // indirect
	gorgonia.org/cu v0.9.3
)
