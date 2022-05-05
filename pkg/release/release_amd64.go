/*
Copyright © 2021 Microshift Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package release

// For the amd64 architecture we use the existing and tested and
// published OKD or other component upstream images

func init() {
	Image = map[string]string{
		"cli":                           "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:5cb8eed4d9713a7a3da331026e21773c76d52feef9a591921807ccf217f66757",
		"coredns":                       "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ceb0d1d2015b87e9daf3e57b93f5464f15a1386a6bcab5442b7dba594b058b24",
		"haproxy_router":                "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a18d11de075fc9b11409f4bc4864bc73023d1510b0bb62b78f156d9a8819386e",
		"kube_flannel":                  "quay.io/coreos/flannel:v0.14.0",
		"kube_flannel_cni":              "quay.io/microshift/flannel-cni:" + Base,
		"kube_rbac_proxy":               "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:08e8b4004edaeeb125ced09ab2c4cd6d690afaf3a86309c91a994dec8e3ccbf3",
		"kubevirt_hostpath_provisioner": "quay.io/kubevirt/hostpath-provisioner:v0.8.0",
		"pause":                         "k8s.gcr.io/pause:3.2",
		"service_ca_operator":           "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:794adb97803879fc56331519893da0d8d17f31e058c0bca4972bc9937a253ef4",
	}
}
