package constants

const (
	ClusterDomainName              = "cluster-domain"
	DefaultHTTPOpenPort            = "80"
	DefaultHTTPSOpenPort           = "443"
	DefaultServiceVersion          = "v0"
	GatewaySecretName              = "rio-certs"
	IstioGatewayDeploy             = "istio-gateway"
	IstioGateway                   = "istio-gateway"
	IstioMeshConfigKey             = "meshConfig"
	IstionConfigMapName            = "mesh"
	IstioSidecarTemplateName       = "sidecarTemplate"
	IstioStackName                 = "istio"
	IstioTelemetry                 = "istio-telemetry"
	LetsEncryptProductionServerURL = "https://acme-v02.api.letsencrypt.org/directory"
	LetsEncryptStagingServerURL    = "https://acme-staging-v02.api.letsencrypt.org/directory"
	ProductionIssuerName           = "letsencrypt-production-issuer"
	ProductionType                 = "production"
	Prometheus                     = "prometheus"
	PublicDomainType               = "RIO_PUBLICDOMAIN_CERT_TYPE"
	RDNSURL                        = "https://api.on-rio.io/v1"
	RioGateway                     = "rio-gateway"
	RioWildcardType                = "RIO_WILDCARD_CERT_TYPE"
	SelfSignedIssuerName           = "selfsigned-issuer"
	SelfSignedType                 = "selfsigned"
	StagingIssuerName              = "letsencrypt-staging-issuer"
	StagingType                    = "staging"
)
