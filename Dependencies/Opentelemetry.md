### What is OpenTelemetry?
- OpenTelemetry is:
    - An observability framework and toolkit designed to facilitate the
        - Generation
        - Export
        - Collection of telemetry data such as traces, metrics, and logs.
    - Open source, as well as vendor- and tool-agnostic, meaning that it can be used with a broad variety of observability backends, including open source tools like Jaeger and Prometheus, as well as commercial offerings. 

### What is observability?
Observability is the ability to understand the internal state of a system by examining its outputs. In the context of software, this means being able to understand the internal state of a system by examining its telemetry data, which includes traces, metrics, and logs.
To make a system observable, it must be instrumented. That is, the code must emit traces, metrics, or logs. The instrumented data must then be sent to an observability backend
### Reliability and metrics
Telemetry refers to data emitted from a system and its behavior. The data can come in the form of traces, metrics, and logs.

Reliability answers the question: “Is the service doing what users expect it to be doing?” A system could be up 100% of the time, but if, when a user clicks “Add to Cart” to add a black pair of shoes to their shopping cart, the system doesn’t always add black shoes, then the system could be unreliable.

Metrics are aggregations over a period of time of numeric data about your infrastructure or application. Examples include: system error rate, CPU utilization, and request rate for a given service. For more on metrics and how they relate to OpenTelemetry, see Metrics.

SLI, or Service Level Indicator, represents a measurement of a service’s behavior. A good SLI measures your service from the perspective of your users. An example SLI can be the speed at which a web page loads.

SLO, or Service Level Objective, represents the means by which reliability is communicated to an organization/other teams. This is accomplished by attaching one or more SLIs to business value.
### Understanding distributed tracing
Distributed tracing lets you observe requests as they propagate through complex, distributed systems. Distributed tracing improves the visibility of your application or system’s health and lets you debug behavior that is difficult to reproduce locally. It is essential for distributed systems, which commonly have nondeterministic problems or are too complicated to reproduce locally.

To understand distributed tracing, you need to understand the role of each of its components: logs, spans, and traces