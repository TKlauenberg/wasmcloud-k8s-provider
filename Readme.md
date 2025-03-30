# wasmCloud Kubernetes Provider

The `wasmCloud Kubernetes Provider` is a Kubernetes operator designed to bridge the gap between Kubernetes resources and wasmCloud components. This provider enables wasmCloud actors to dynamically interact with Kubernetes resources by watching, creating, updating, and deleting them. It provides a seamless way for wasmCloud components to integrate with Kubernetes, leveraging the power of both ecosystems.

## Overview

The wasmCloud Kubernetes Provider acts as a Kubernetes operator that listens for requests from wasmCloud components. These requests specify the Kubernetes resources that the components want to manage or observe. The provider handles these requests and performs the necessary operations on the Kubernetes cluster.

### Key Features

1. **Dynamic Resource Watching**: wasmCloud components can register specific Kubernetes resources they want to watch. The provider sets up dynamic watchers for these resources and forwards events (e.g., creation, updates, deletions) to the requesting wasmCloud components.

2. **Resource Management**: wasmCloud components can request the creation, update, or deletion of Kubernetes resources such as ConfigMaps, Deployments, Services, and more.

3. **Seamless Integration**: The provider leverages Kubernetes' controller-runtime library to implement operator patterns, ensuring robust and scalable interactions with the cluster.

4. **Event Forwarding**: Events from Kubernetes resources are forwarded to wasmCloud actors, enabling real-time processing and decision-making.

## Architecture

The wasmCloud Kubernetes Provider is built as a Kubernetes operator using Go. It utilizes the following components:

- **Kubernetes Controller**: Implements the operator pattern to manage Kubernetes resources dynamically.
- **wasmCloud Provider SDK**: Handles communication between the provider and wasmCloud components.
- **Dynamic Watchers**: Allows runtime registration of Kubernetes resources to watch for changes.
- **Resource Handlers**: Implements logic for creating, updating, and deleting Kubernetes resources based on wasmCloud requests.

### Workflow

1. **Registration**: A wasmCloud component sends a request to the provider to register a resource it wants to watch or manage.
2. **Dynamic Watcher Setup**: The provider sets up a dynamic watcher for the specified resource.
3. **Event Handling**: When changes occur in the resource, the provider forwards the events to the requesting wasmCloud component.
4. **Resource Management**: The provider processes requests to create, update, or delete resources in the Kubernetes cluster.

## Example Use Cases

1. **Real-Time Monitoring**: A wasmCloud actor registers to watch a specific ConfigMap. The provider forwards any changes to the ConfigMap to the actor for processing.
2. **Dynamic Resource Creation**: A wasmCloud actor requests the creation of a Deployment. The provider creates the Deployment in the Kubernetes cluster and returns the status to the actor.
3. **Event-Driven Workflows**: Kubernetes resource events trigger workflows in wasmCloud actors, enabling event-driven architectures.

## Getting Started

### Prerequisites

- A Kubernetes cluster
- wasmCloud runtime
- Go (1.24 or later)

### Deployment

1. Build the provider:
   ```bash
   go build -o wasmcloud-k8s-provider main.go