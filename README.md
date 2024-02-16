# mackerel-plugin-gunfish

Mackerel plugin for [Gunfish](https://github.com/mackee/gunfish)

## Install

Use [mkr](https://github.com/mackerelio/mkr).

```console
# mkr plugin install kayac/mackerel-plugin-gunfish@v0.1.0
```

## Synopsis

```shell
Usage of /opt/mackerel-agent/plugins/bin/mackerel-plugin-gunfish:
  -host string
    	Hostname (default "localhost")
  -metric-key-prefix string
    	Metric key prefix (default "gunfish")
  -port string
    	Port (default "8003")
  -tempfile string
    	Temp file name
```

## Example of mackerel-agent.conf

```
[plugin.metrics.gunfish]
command = "/opt/mackerel-agent/plugins/bin/mackerel-plugin-gunfish/mackerel-plugin-gunfish -port=8203"

[plugin.metrics.gunfish-gostats]
command = "mackerel-plugin-gostats -port=8203 -path=/stats/profile -metric-key-prefix=gunfish"
```

Gunfish has a [golang-stats-api-handler](https://github.com/fukata/golang-stats-api-handler) endpoint on `/stats/profile`. You can get Go's metrics by mackerel-plugin-gostats.

## How to release

[goxc](https://github.com/laher/goxc) and [ghr](https://github.com/tcnksm/ghr) are used to release.

### Release by manually

1. Install goxc and ghr by `make setup`
2. Edit CHANGELOG.md, git commit, git push
3. `git tag vx.y.z`
4. GITHUB_TOKEN=... make release
5. See https://github.com/mackerelio/mackerel-plugin-gunfish/releases

## Author

KAYAC Inc.

## License

Copyright 2014 Hatena Co., Ltd.

Copyright 2017 KAYAC Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
