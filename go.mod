module github.com/netlify/gotrue

require (
	github.com/GoogleCloudPlatform/cloudsql-proxy v1.27.0
	github.com/badoux/checkmail v0.0.0-20170203135005-d0a759655d62
	github.com/beevik/etree v1.1.1-0.20200718192613-4a2f8b9d084c
	github.com/didip/tollbooth/v5 v5.1.1
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gobuffalo/pop/v5 v5.3.4
	github.com/gobuffalo/uuid v2.0.5+incompatible
	github.com/golang-jwt/jwt/v4 v4.1.0
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/imdario/mergo v0.0.0-20160216103600-3e95a51e0639
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/microcosm-cc/bluemonday v1.0.18 // indirect
	github.com/netlify/mailme v1.1.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.6.0
	github.com/russellhaering/gosaml2 v0.6.0
	github.com/russellhaering/goxmldsig v1.1.0
	github.com/sebest/xff v0.0.0-20160910043805-6c115e0ffa35
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1
	golang.org/x/time v0.0.0-20220609170525-579cf78fd858 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.12.1
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/beevik/etree => github.com/beevik/etree v1.1.1-0.20200718192613-4a2f8b9d084c
	github.com/gobuffalo/github_flavored_markdown => github.com/gobuffalo/github_flavored_markdown v1.1.1
	github.com/russellhaering/goxmldsig => github.com/russellhaering/goxmldsig v1.1.1
)

go 1.15
