package redis

import (
	"github.com/codegangsta/cli"
	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-cli/pluginlib/product"
	"github.com/enaml-ops/omg-cli/pluginlib/util"
	"github.com/xchapter7x/lo"
)

const (
	BoshReleaseURL = "https://bosh.io/d/github.com/cloudfoundry-community/redis-boshrelease"
	BoshReleaseVer = "12"
	BoshReleaseSHA = "324910eaf68e8803ad2317d5a2f5f6a06edc0a40"
	Master         = iota
	Slave
	Errand
	Pool
)

type jobBucket struct {
	JobName   string
	JobType   int
	Instances int
}
type Plugin struct{}

func (s *Plugin) GetFlags() (flags []cli.Flag) {
	return []cli.Flag{
		cli.StringSliceFlag{Name: "leader-ip", Usage: "multiple static ips for each redis leader vm"},
		cli.IntFlag{Name: "leader-instances", Value: 1, Usage: "the number of leader instances to provision"},
		cli.StringFlag{Name: "redis-pass", Value: "red1s", Usage: "the password to use for connecting redis nodes"},
		cli.IntFlag{Name: "pool-instances", Value: 2, Usage: "number of instances in the redis cluster"},
		cli.StringFlag{Name: "disk-size", Value: "4096", Usage: "size of disk on VMs"},
		cli.IntFlag{Name: "slave-instances", Value: 1, Usage: "number of slave VMs"},
		cli.IntFlag{Name: "errand-instances", Value: 1, Usage: "number of errand VMs"},
		cli.StringSliceFlag{Name: "slave-ip", Usage: "list of slave VM Ips"},
		cli.StringFlag{Name: "network-name", Usage: "name of your target network"},
		cli.StringFlag{Name: "vm-size", Usage: "name of your desired vm size"},
		cli.StringFlag{Name: "stemcell-url", Usage: "the url of the stemcell you wish to use"},
		cli.StringFlag{Name: "stemcell-ver", Usage: "the version number of the stemcell you wish to use"},
		cli.StringFlag{Name: "stemcell-sha", Usage: "the sha of the stemcell you will use"},
		cli.StringFlag{Name: "stemcell-name", Value: "trusty", Usage: "the name of the stemcell you will use"},
	}
}

func (s *Plugin) GetMeta() product.Meta {
	return product.Meta{
		Name: "redis",
	}
}

func (s *Plugin) GetProduct(args []string, cloudConfig []byte) (b []byte) {
	c := pluginutil.NewContext(args, s.GetFlags())
	lo.G.Debug("context", c)
	var dm = new(enaml.DeploymentManifest)
	dm.SetName("enaml-redis")
	dm.Update = enaml.Update{
		Canaries:        1,
		CanaryWatchTime: "1000-100000",
		MaxInFlight:     50,
		UpdateWatchTime: "1000-100000",
	}
	dm.Properties = enaml.Properties{
		"redis": struct{}{},
	}
	dm.AddRemoteRelease("redis", BoshReleaseVer, BoshReleaseURL, BoshReleaseSHA)
	dm.AddRemoteStemcell(c.String("stemcell-name"), c.String("stemcell-name"), c.String("stemcell-ver"), c.String("stemcell-url"), c.String("stemcell-sha"))

	for _, bkt := range []jobBucket{
		jobBucket{JobName: "redis_leader_z1", JobType: Master, Instances: c.Int("leader-instances")},
		jobBucket{JobName: "redis_z1", JobType: Pool, Instances: c.Int("pool-instances")},
		jobBucket{JobName: "redis_test_slave_z1", JobType: Slave, Instances: c.Int("slave-instances")},
		jobBucket{JobName: "acceptance-tests", JobType: Errand, Instances: c.Int("errand-instances")},
	} {
		dm.AddJob(NewRedisJob(
			bkt.JobName,
			c.String("network-name"),
			c.String("redis-pass"),
			c.String("disk-size"),
			c.String("vm-size"),
			c.StringSlice("leader-ip"),
			c.StringSlice("slave-ip"),
			bkt.Instances,
			bkt.JobType,
		))
	}
	return dm.Bytes()
}

func NewRedisJob(name, networkName, pass, disk, vmSize string, masterIPs, slaveIPs []string, instances int, jobType int) (job enaml.Job) {
	var lifecycle string
	network := enaml.Network{Name: networkName}
	properties := enaml.Properties{
		"network": networkName,
		"redis": map[string]interface{}{
			"password": pass,
		},
	}
	template := enaml.Template{Name: "redis", Release: "redis"}

	switch jobType {
	case Master:
		network.StaticIPs = masterIPs

	case Slave:
		network.StaticIPs = slaveIPs
		properties["redis"].(map[string]interface{})["master"] = masterIPs[0]
		properties["redis"].(map[string]interface{})["slave"] = slaveIPs[0]

	case Errand:
		lifecycle = "errand"
		properties["redis"].(map[string]interface{})["master"] = masterIPs[0]
		properties["redis"].(map[string]interface{})["slave"] = slaveIPs[0]
		template = enaml.Template{Name: "acceptance-tests", Release: "redis"}

	default:
		properties["redis"].(map[string]interface{})["master"] = masterIPs[0]
	}

	job = enaml.Job{
		Name:       name,
		Lifecycle:  lifecycle,
		Properties: properties,
		Instances:  instances,
		Networks: []enaml.Network{
			network,
		},
		Templates:      []enaml.Template{template},
		PersistentDisk: disk,
		ResourcePool:   vmSize,
		Update: enaml.Update{
			Canaries: 10,
		},
	}
	lo.G.Debug("job", job)
	return
}
