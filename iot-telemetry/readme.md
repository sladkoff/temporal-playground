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