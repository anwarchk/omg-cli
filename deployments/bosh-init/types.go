package boshinit

import (
	"github.com/bosh-ops/bosh-install/deployments/bosh-init/enaml-gen/aws_cpi"
	"github.com/bosh-ops/bosh-install/deployments/bosh-init/enaml-gen/director"
	"github.com/bosh-ops/bosh-install/deployments/bosh-init/enaml-gen/postgres"
	"github.com/bosh-ops/bosh-install/deployments/bosh-init/enaml-gen/registry"
)

type BoshInitConfig struct {
	Name                  string
	BoshReleaseVersion    string
	BoshReleaseSHA        string
	BoshPrivateIP         string
	BoshCPIReleaseVersion string
	BoshCPIReleaseSHA     string
	GoAgentVersion        string
	GoAgentSHA            string
	BoshAvailabilityZone  string
	BoshInstanceSize      string
	BoshDirectorName      string
	AWSSubnet             string
	AWSElasticIP          string
	AWSPEMFilePath        string
	AWSAccessKeyID        string
	AWSSecretKey          string
	AWSRegion             string
}

type Rr registry.Registry
type Ar aws_cpi.Registry

type RegistryProperty struct {
	Rr      `yaml:",inline"`
	Ar      `yaml:",inline"`
	Address string `yaml:"address"`
}
type user struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

type DirectorProperty struct {
	director.Director `yaml:",inline"`
	Address           string
}

type PgSql struct {
	User     string
	Host     string
	Password string
	Database string
	Adapter  string
}

type Postgres interface {
	GetDirectorDB() *director.Db
	GetRegistryDB() *registry.Db
	GetPostgresDB() postgres.Postgres
}