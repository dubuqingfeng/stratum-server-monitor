## stratum server 监控

### 功能

1. 监听 stratum server，任务更新事件
2. 写入高度到 redis
3. 高度变化事件到 队列 ，或者 mysql