# IOT Telemetry with Temporal

## Use case

1. Message comes in
2. Parse Message and check schema
3. Split message into different types
    1. Telemetry
        1. Split telemetry into different types
        2. Parse and validate
        3. Store in database
    2. Errors
        1. Parse and validate
        2. Store in database
        3. Healthcheck with aggregated errors per asset

## Running the workflow

```bash
go run start/main.go "$(<resources/sample-message.json)"
```

## Ship it day presentation

1. What's Temporal?
   1. Durable execution framework (like AWS Step Functions or Azure Durable Functions)
   2. Open source self-hosted server or Temporal Cloud
   3. SDKs for different languages
2. Demo
   1. Start server
   2. Start worker (with bug)
   3. Start workflow
   4. Show Temporal Web UI
   5. Find bug in Web UI
   6. Fix bug live
   7. Restart workflow from Web UI
   8. 🎉
3. Conclusion
   1. The concept seems promising but requires a mindset shift
   2. Lots of constraints that need to be embraced
   3. Currently Temporal is not really viable (Temporal Cloud is 200 bucks a month, self hosting is not explored and probably doesn't scale to zero)
   4. Fun to play around with
   5. Transferable concepts to other systems
