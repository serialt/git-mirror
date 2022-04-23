## Git-Mirror

同步github项目到gitee

默认读取的配置文件为`~/.git-mirror.yaml`,需要先配置gitee的访问token

配置文件模板
```yaml
gitLog:
  logLevel: debug
workspace: /tmp/mac_mirror
gitee:
  user: xxxxx
  token: xxxxxxxxxxx
githubRepo:
  - melbahja/goph
```

### build
```shell
cd cmd && go build -o git-mirror
```
