/*

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

package cinderscheduler

const (
	// KollaConfig -
	KollaConfig = "/var/lib/config-data/merged/cinder-scheduler-config.json"
	// ServiceCommand -
	ServiceCommand = "/usr/local/bin/kolla_set_configs && /usr/local/bin/kolla_start"
	// ProbeCommand run with 'bash -c'.  Requires -E to preserve env variables
	ProbeCommand = "sudo -E /usr/local/bin/kolla_set_configs && /usr/local/bin/container-scripts/healthcheck.py scheduler"
	// ProbeDebug run with 'bash -c'.  Requires -E to preserve env variables
	ProbeDebug = "sudo -E /usr/local/bin/kolla_set_configs && /bin/sleep infinity"
)
