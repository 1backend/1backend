# Log Accumulator

This package provides an in-memory object for collecting logs, designed to trigger a flush either after a specified time interval (ensuring frequent updates for the consumer) or once a defined number of logs have been gathered from a producer (identified by the ProducerID field).
