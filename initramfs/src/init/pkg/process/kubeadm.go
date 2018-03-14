package process

import (
	"io/ioutil"
	"os"

	"github.com/autonomy/dianemo/initramfs/src/init/pkg/process/conditions"
)

const MasterConfiguration = `
kind: MasterConfiguration
apiVersion: kubeadm.k8s.io/v1alpha1
skipTokenPrint: true
networking:
  dnsDomain: cluster.local
  ProcessSubnet: 10.96.0.0/12
  podSubnet: 10.244.0.0/16
featureGates:
  HighAvailability: true
  SelfHosting: true
  StoreCertsInSecrets: true
  DynamicKubeletConfig: true
`

type Kubeadm struct{}

func init() {
	os.MkdirAll("/etc/kubernetes", os.ModeDir)
	if err := ioutil.WriteFile("/etc/kubernetes/kubeadm.yaml", []byte(MasterConfiguration), 0644); err != nil {

	}
}

func (p *Kubeadm) Cmd() (name string, args []string) {
	name = "/bin/kubeadm"
	args = []string{
		"init",
		"--config=/etc/kubernetes/kubeadm.yaml",
	}

	return name, args
}

func (p *Kubeadm) Condition() func() (bool, error) {
	return conditions.WaitForFileExists("/var/run/docker.sock")
}

func (p *Kubeadm) Env() []string { return []string{} }

func (p *Kubeadm) Type() Type { return Once }