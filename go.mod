module github.com/vmware/cloud-provider-for-cloud-director

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/antihax/optional v1.0.0
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/google/uuid v1.2.0
	github.com/kr/text v0.2.0 // indirect
	github.com/peterhellberg/link v1.1.0
	github.com/stretchr/testify v1.7.0
	github.com/vmware/go-vcloud-director/v2 v2.12.0-alpha.4
	go.uber.org/atomic v1.3.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sys v0.0.0-20210314195730-07df6a141424 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/apiserver v0.0.0
	k8s.io/client-go v1.5.1
	k8s.io/cloud-provider v0.0.0
	k8s.io/component-base v0.0.0
	k8s.io/klog v0.3.1
	k8s.io/kubernetes v1.15.5
	k8s.io/utils v0.0.0-20210305010621-2afb4311ab10 // indirect
)

replace k8s.io/api => k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190918200908-1e17798da8c1

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190918202139-0b14c719ca62

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190918200256-06eb1244587a

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190918203125-ae665f80358a

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20190918202959-c340507a5d48

replace k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190612205613-18da4a14b22b

replace k8s.io/component-base => k8s.io/component-base v0.0.0-20190918200425-ed2f0867c778

replace k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190817025403-3ae76f584e79

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20190918203248-97c07dcbb623

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190918201136-c3a845f1fbb2

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20190918202837-c54ce30c680e

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20190918202429-08c8357f8e2d

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20190918202713-c34a54b3ec8e

replace k8s.io/kubelet => k8s.io/kubelet v0.0.0-20190918202550-958285cf3eef

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190918203421-225f0541b3ea

replace k8s.io/metrics => k8s.io/metrics v0.0.0-20190918202012-3c1ca76f5bda

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20190918201353-5cc279503896
