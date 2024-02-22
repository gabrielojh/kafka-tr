# Prerequisities
This project can be deployed on a local kubernetes cluster. Before you begin, ensure you have the following software installed on your system:

## 1. Docker

Docker is used as Minikube's container runtime. Ensure you have Docker installed and running on your system.

- [Install Docker](https://docs.docker.com/get-docker/)

## 2. kubectl

`kubectl` is a command-line tool that allows you to run commands against Kubernetes clusters.

- [Install kubectl](https://kubernetes.io/docs/tasks/tools/)

## 3. Minikube

Minikube is a utility that runs a single-node Kubernetes cluster locally on your machine.

- [Install Minikube](https://minikube.sigs.k8s.io/docs/start/)

## 4. Tilt

Tilt automates the development process in Kubernetes, offering real-time feedback and updates for a smoother and more efficient workflow. It's designed to enhance productivity by automatically applying changes to your environment as you develop.

- [Install Tilt](https://docs.tilt.dev/install.html)

# Setup

## Kubernetes Cluster Setup

1. **Start Minikube**: Initialize your local Kubernetes cluster with Minikube.

   ```bash
   minikube start
   ```

2. **Run Tilt**: Deploy kubernetes in minikube with Tilt

    ```bash
    tilt up
    ```