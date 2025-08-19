# Supabase Transaction Pooling + Application Connection Pooling

## Overview

Your application now uses a **dual-layer connection pooling strategy**:
1. **Supabase Transaction Pooling** (External) - Managed by Supabase
2. **Application Connection Pooling** (Internal) - Managed by pgxpool

## Configuration Details

### Supabase Pooler Configuration
```
Host: aws-0-us-east-2.pooler.supabase.com
Port: 6543
Pool Mode: transaction
User: postgres.tadkhffiuamvixlewpwf
```

### Application Pool Configuration (Optimized for Supabase)
```bash
PG_MAX_CONNS=10          # Smaller pool since Supabase handles pooling
PG_MIN_CONNS=2           # Minimal persistent connections
PG_MAX_CONN_LIFETIME=30m # Shorter lifetime for efficient rotation
PG_MAX_CONN_IDLE_TIME=15m # Quick cleanup of idle connections
```

## How It Works

### 1. **Application Level (Your Go App)**
- Your app maintains a small pool of 2-10 connections
- pgxpool manages connection lifecycle and reuse
- Connections are efficiently shared across goroutines

### 2. **Supabase Level (External Pooler)**
- Supabase's transaction pooler handles the actual PostgreSQL connections
- Each "connection" from your app is actually a pooled transaction
- Supabase manages the real database connections efficiently

## Benefits of This Approach

### ✅ **Double Efficiency**
- Application-level pooling reduces connection creation overhead
- Supabase pooling manages actual database connections efficiently

### ✅ **Optimal Resource Usage**
- Small local pool (2-10 connections) reduces memory usage
- Supabase handles the heavy lifting for database connection management

### ✅ **Better Performance**
- Faster transaction handling through Supabase's transaction pooling
- Reduced latency from connection reuse at both levels

### ✅ **Scalability**
- Can handle many concurrent requests with minimal connections
- Supabase pooler can serve multiple applications efficiently

## Connection Flow

```
[Your App] → [pgxpool (2-10 conns)] → [Supabase Pooler] → [PostgreSQL Database]
```

1. **Request comes in** → Your app gets a connection from pgxpool
2. **Transaction starts** → Supabase pooler assigns a database connection
3. **Query executes** → Direct to PostgreSQL
4. **Transaction ends** → Supabase pooler releases database connection
5. **Request completes** → Your app returns connection to pgxpool

## Environment Variables

### Required
```bash
PG_USER=postgres.tadkhffiuamvixlewpwf
PG_PASSWORD=[YOUR-PASSWORD]
PG_HOST=aws-0-us-east-2.pooler.supabase.com
PG_PORT=6543
PG_DATABASE=postgres
PG_SSLMODE=require
```

### Optional (Optimized for Supabase)
```bash
PG_MAX_CONNS=10          # Small pool for Supabase
PG_MIN_CONNS=2           # Minimal persistent connections
PG_MAX_CONN_LIFETIME=30m # Shorter rotation
PG_MAX_CONN_IDLE_TIME=15m # Quick cleanup
```

## Monitoring

### Application Pool Stats
Access `GET /api/v1/stats` to see:
- Current active connections
- Pool utilization
- Connection acquisition metrics

### Supabase Dashboard
Monitor Supabase pooler performance:
- Pool utilization
- Transaction throughput
- Connection wait times

## Best Practices

### 1. **Keep Local Pool Small**
- Since Supabase handles pooling, keep `PG_MAX_CONNS` low (5-15)
- Reduces memory usage and connection overhead

### 2. **Short Connection Lifetimes**
- Use shorter `PG_MAX_CONN_LIFETIME` (15-30 minutes)
- Allows better load distribution across Supabase pooler

### 3. **Monitor Both Levels**
- Watch application pool stats via `/api/v1/stats`
- Monitor Supabase dashboard for pooler performance

### 4. **Handle Connection Errors Gracefully**
- Implement retry logic for connection failures
- Use circuit breakers for database unavailability

## Troubleshooting

### Issue: "Too many connections"
**Solution**: Reduce `PG_MAX_CONNS` in your application

### Issue: "Slow query performance"
**Solution**: 
- Check Supabase pooler utilization
- Consider increasing `PG_MIN_CONNS` slightly

### Issue: "Connection timeouts"
**Solution**:
- Reduce `PG_MAX_CONN_IDLE_TIME`
- Check network connectivity to Supabase

## Performance Tips

1. **Use Transactions Efficiently**: Supabase transaction pooling works best with short transactions
2. **Batch Operations**: Group multiple queries in single transactions when possible
3. **Monitor Pool Stats**: Regularly check pool utilization via the stats endpoint
4. **Adjust Based on Load**: Fine-tune pool settings based on actual application usage patterns
