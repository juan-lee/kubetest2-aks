package deployer

import (
	"flag"

	"github.com/octago/sflags/gen/gpflag"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"

	"github.com/juan-lee/kubetest2-aks/deployer/options"
	"sigs.k8s.io/kubetest2/pkg/types"
)

// Name is the name of the kubetest2 deployer
const Name = "aks"

var (
	GitTag string
	_      types.NewDeployer            = New
	_      types.Deployer               = &aksDeployer{}
	_      types.DeployerWithKubeconfig = &aksDeployer{}
)

// New implements deployer.New for aks
func New(opts types.Options) (types.Deployer, *pflag.FlagSet) {
	co := options.ClusterOptions{}
	fs := bindFlags(&co)
	klog.InitFlags(nil)
	fs.AddGoFlagSet(flag.CommandLine)
	return newDeployer(opts, &co), fs
}

func newDeployer(opts types.Options, co *options.ClusterOptions) *aksDeployer {
	return &aksDeployer{
		Options:        opts,
		ClusterOptions: co,
	}
}

type aksDeployer struct {
	types.Options
	*options.ClusterOptions

	clusterResourceID string
	kubeconfig        string
}

// DumpClusterLogs should export logs from the cluster. It may be called
// multiple times. Options for this should come from New(...)
func (ad *aksDeployer) DumpClusterLogs() error {
	panic("not implemented") // TODO: Implement
}

// Build should build kubernetes and package it in whatever format
// the deployer consumes
func (ad *aksDeployer) Build() error {
	panic("not implemented") // TODO: Implement
}

func bindFlags(co *options.ClusterOptions) *pflag.FlagSet {
	flags, err := gpflag.Parse(co)
	if err != nil {
		klog.Fatalf("unable to generate flags from deployer")
		return nil
	}

	return flags
}
