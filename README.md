Resources Check
=================

Ensure every Deployment, Statefulset and Daemonset on Kubernetes has request and limits set

# Description

This project can run on Kubernetes itself. 

At every specified interval, it runs through every Deployment, Statefulset and Daemonset on the cluster. 

If a not ignored controller doesn't have limits and/or request specified, an event is sent to any monitoring service (in this case, it is implemented to send to Datadog).

# Development

`make start`

# Test

`make test`
