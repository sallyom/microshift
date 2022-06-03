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
// published OCP or other component upstream images

func init() {
	Image = map[string]string{
		"cli":                           "quay.io/openshift/okd-content@sha256:4be491a686dc8ec268a81721bf40010415ea124c926be496546d66b129245539",
		"coredns":                       "quay.io/openshift/okd-content@sha256:82723e1d41ab68c0fde2e2f8bfa22ea470a11dd3101d7d214cdf9dd63171788d",
		"haproxy_router":                "quay.io/openshift/okd-content@sha256:131cf281cc34e5de3b7ca2d63690484a2c54d7f2e8803c909b0a16d1ae781f51",
		"kube_flannel":                  "quay.io/coreos/flannel:v0.14.0",
		"kube_flannel_cni":              "quay.io/microshift/flannel-cni:v0.14.0",
		"kube_rbac_proxy":               "quay.io/openshift/okd-content@sha256:baedb268ac66456018fb30af395bb3d69af5fff3252ff5d549f0231b1ebb6901",
		"kubevirt_hostpath_provisioner": "quay.io/kubevirt/hostpath-provisioner:v0.8.0",
		"pause":                         "k8s.gcr.io/pause:3.6",
		"service_ca_operator":           "quay.io/openshift/okd-content@sha256:0692de34bfd9f40455a5afc69fb47be56b1f386c55183b569d9d71d2969ff2e6",
	}
}
