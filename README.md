# dd-trace-editor

Running this program will:

 - Read a file `traces` containing a list of traces
 - Create new IDs for these traces (traces with same IDs don't get displayed in the UI)
 - Shift the timestamps to ~now (so the traces appear in the live view)
 - Make any other change you want to the traces once you modify the code
 - Write the new traces to `new_traces`
 
 This new file can then be sent to the agent with something like:
 ```
 curl --location --request PUT 'http://localhost:8126/v0.4/traces' \
--header 'Content-Type: application/msgpack' \
--header 'Datadog-Meta-Tracer-Version: 3.4.0' \
--header 'X-Datadog-Trace-Count: 1' \
--header 'Datadog-Meta-Lang: nodejs' \
--header 'Datadog-Meta-Lang-Version: v18.11.0' \
--header 'Datadog-Meta-Lang-Interpreter: v8' \
--data-binary '@new_traces'
 ```
 
 Walkthrough for sending custom traces to the agent: https://datadoghq.atlassian.net/l/cp/x050Er32
