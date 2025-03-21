package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "cloud.google.com/go"
	_ "github.com/aws/aws-sdk-go/aws"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "go.mongodb.org/mongo-driver/mongo"
	_ "k8s.io/client-go/kubernetes"

	_ "github.com/apache/thrift/lib/go/thrift"
	_ "github.com/cespare/xxhash/v2"
	_ "github.com/elastic/go-elasticsearch/v8"
	_ "github.com/emicklei/go-restful/v3"
	_ "github.com/fxamacker/cbor/v2"
	_ "github.com/gin-contrib/sse"
	_ "github.com/go-logr/logr"
	_ "github.com/go-openapi/jsonpointer"
	_ "github.com/go-openapi/jsonreference"
	_ "github.com/go-openapi/swag"
	_ "github.com/go-playground/locales"
	_ "github.com/go-playground/universal-translator"
	_ "github.com/go-playground/validator/v10"
	_ "github.com/go-stack/stack"
	_ "github.com/gocql/gocql"
	_ "github.com/golang/snappy"
	_ "github.com/google/gofuzz"
	_ "github.com/google/uuid"
	_ "github.com/hashicorp/consul/api"
	_ "github.com/hashicorp/vault/api"
	_ "github.com/influxdata/influxdb-client-go/v2"
	_ "github.com/ipfs/kubo"
	_ "github.com/josharian/intern"
	_ "github.com/json-iterator/go"
	_ "github.com/klauspost/compress"
	_ "github.com/leodido/go-urn"
	_ "github.com/lib/pq"
	_ "github.com/mailru/easyjson"
	_ "github.com/mattn/go-isatty"
	_ "github.com/modern-go/concurrent"
	_ "github.com/modern-go/reflect2"
	_ "github.com/munnerz/goautoneg"
	_ "github.com/nats-io/nats.go"
	_ "github.com/olivere/elastic/v7"
	_ "github.com/pkg/errors"
	_ "github.com/prometheus/procfs"
	_ "github.com/uber-go/dosa"
	_ "github.com/ugorji/go/codec"
	_ "github.com/x448/float16"
	_ "github.com/xdg-go/pbkdf2"
	_ "github.com/xdg-go/scram"
	_ "github.com/xdg-go/stringprep"
	_ "github.com/youmark/pkcs8"
	_ "go.elastic.co/apm"
	_ "golang.org/x/oauth2"
	_ "golang.org/x/term"
	_ "golang.org/x/text"
	_ "gopkg.in/evanphx/json-patch.v4"
	_ "gopkg.in/inf.v0"
	_ "gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v3"
	_ "k8s.io/api"
	_ "k8s.io/apimachinery"
	_ "k8s.io/klog/v2"
	_ "sigs.k8s.io/json"
	_ "sigs.k8s.io/yaml"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	r.Run(":8080") // Listen and serve on 0.0.0.0:8080

	// Bring logging for debugging
	log.Println("Service is up and running")

	// To ensure the imports are executed
	fmt.Println("All imports are included")
}

