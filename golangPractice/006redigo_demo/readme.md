## keng

```
### String or Strings
log.Println(redis.String(c.Do("info")))
log.Println(redis.Strings(c.Do("config", "get", "*")))

### slowlog get
slowlogs, _ := redis.SlowLogs(c.Do("slowlog", "get"))
log.Println(slowlogs)
```