# Get started with Pyroscope

Pyroscope is an open source continuous profiling platform. It will help you:

- Optimize your application's performance
- Debug hard-to-reproduce issues
- Understand your application's bottlenecks
- Monitor your application's performance over time
- And much more!

## Download and configure Pyroscope

```yaml
services:
  pyroscope:
    image: "grafana/pyroscope:latest"
    ports:
      - "4040:4040"
    command:
      - "server"
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:4040/ready"]
      interval: 30s
      timeout: 10s
      retries: 5
```
