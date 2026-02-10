# cgroup-monitor-go

一个使用 `Go` 实现的 `cgroup v2` 资源监控工具，用于检测
`Memory OOM`、`CPU throttling` 等常见容器/服务问题。

## 背景
在容器环境中，服务常因 cgroup 限制被 kill，但传统监控难以及时反映。

## 功能
- 监控 memory.events / cpu.stat
- 输出 OOM / throttling 事件

## 运行使用
```bash
go run main.go -path /sys/fs/cgroup/test -interval 1s
```

# 测试方法
```bash
mkdir /sys/fs/cgroup/test
# $$ 代表当前Shell, cgroup.procs存放当前组PID
echo $$ > /sys/fs/cgroup/test/cgroup.procs
echo 100M > /sys/fs/cgroup/test/memory.max
# 确保当前Shell进程在test组
cat /proc/self/cgroup
0::/test

# 在 test 组的 Shell 下运行
python3 -c "a = b'0' * 1024 * 1024 * 150"  # 尝试申请 150MB 的内存并持有它
```