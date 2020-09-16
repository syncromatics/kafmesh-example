# kafmesh-example

Sometimes an example is worth a thousand words. This repository supplements
the official [Kafmesh] documentation with complete example code that solves a
real-world problem.

## What Does it Do?

kafmesh-example provides a gRPC service for IoT devices to report telemetry,
including device health. Health data is output into kafka and correlated with
device configurations stored in a server-side database. Then we sink health
information out to reporting tools and a data warehouse.

## Major Components

kafmesh-example's functionality is implemented with three main components:

## Assignments

* Reads assignment details of which device is assigned to which customer from the database.
* Reads information configured for each customer.

## Details

* Emits device configuration details from database
* Whenever there is a change in device details or which customer a device is assigned to, puts an update further down the chain

## Heartbeats

* Continuously receives status information from devices and emits into kafka.
* Enriches health heartbeats with configuration details.
* Sinks health information to an external service.

[Kafmesh]: https://github.com/syncromatics/kafmesh
