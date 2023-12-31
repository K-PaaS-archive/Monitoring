package gardener

import (
	"errors"

	specs "github.com/opencontainers/runtime-spec/specs-go"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/garden-shed/rootfs_spec"
	"code.cloudfoundry.org/lager"
)

type NoopVolumeCreator struct{}

var ErrGraphDisabled = errors.New("volume graph is disabled")

func (NoopVolumeCreator) Create(lager.Logger, string, rootfs_spec.Spec) (specs.Spec, error) {
	return specs.Spec{}, ErrGraphDisabled
}

func (NoopVolumeCreator) Destroy(lager.Logger, string) error {
	return nil
}

func (NoopVolumeCreator) Metrics(lager.Logger, string, bool) (garden.ContainerDiskStat, error) {
	return garden.ContainerDiskStat{}, nil
}

func (NoopVolumeCreator) GC(lager.Logger) error {
	return nil
}
