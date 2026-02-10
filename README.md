# cgroup-monitor-go

一个使用 `Go` 实现的 `cgroup v2` 资源监控工具，用于检测
`Memory OOM`、`CPU throttling` 等常见容器/服务问题。

## 背景
在容器环境中，服务常因 cgroup 限制被 kill，但传统监控难以及时反映。

## 功能
- 监控 memory.events / cpu.stat
- 输出 OOM / throttling 事件