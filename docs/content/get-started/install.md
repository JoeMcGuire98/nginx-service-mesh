---
title: "Install NGINX Service Mesh using nginx-meshctl"
date: 2020-02-20T19:43:59Z
draft: false
toc: true
description: "This topic explains how to download, install, and deploy NGINX Service Mesh."
weight: 300
categories: ["tasks"]
docs: "DOCS-681"
---

## Overview

This guide contains instructions for downloading and installing NGINX Service Mesh using the `nginx-meshctl` command line tool.

For Helm users, see how to [Install NGINX Service Mesh using Helm]( {{< ref "/get-started/install-with-helm.md" >}} ).

### Prerequisites

{{< important >}} Before installing NGINX Service Mesh, make sure you've completed the following steps. {{< /important >}}

- You have a working and [supported]({{< ref "/about/tech-specs.md#supported-versions" >}}) Kubernetes cluster.
- You followed the [Kubernetes]( {{< ref "/get-started/kubernetes-platform/_index.md" >}} ) or [OpenShift]( {{< ref "/get-started/openshift-platform/_index.md" >}} ) Platform Setup guide to **prepare your cluster** to work with NGINX Service Mesh.
- You have the Kubernetes `kubectl` command-line utility configured on the machine where you want to install NGINX Service Mesh.
- You reviewed the [Configuration Options for NGINX Service Mesh]( {{< ref "/get-started/configuration.md" >}} ).

### Download nginx-meshctl

The NGINX Service Mesh command-line tool -- `nginx-meshctl` -- allows you to deploy, remove, and interact with the NGINX Service Mesh control plane.

To install NGINX Service Mesh, you need to download the `nginx-meshctl` binary for your architecture. The latest version of `nginx-meshctl` is available on our [Github releases](https://github.com/nginxinc/nginx-service-mesh/releases/latest) page.

### Install the CLI

The following sections describe how to install the CLI on Linux, macOS, and Windows.

#### Install on Linux

1. Download the appropriate binary for your architecture, either `nginx-meshctl_<version>_linux_amd64.tar.gz` or `nginx-meshctl_<version>_linux_arm64.tar.gz`.

1. Unzip the binary.

    ```bash
    tar -xvf nginx-meshctl_<version>_linux_amd64.tar.gz nginx-meshctl
    ```

1. Move the `nginx-meshctl` executable in to your PATH.

    ```bash
    sudo mv nginx-meshctl /usr/local/bin/nginx-meshctl
    ```

1. Ensure the `nginx-meshctl` is executable.

    ```bash
    sudo chmod +x /usr/local/bin/nginx-meshctl
    ```

1. Test the installation.

    ```bash
    nginx-meshctl
    ```

#### Install on macOS

1. Download the appropriate binary for your architecture, either `nginx-meshctl_<version>_darwin_arm64.tar.gz` for M1 Macs or `nginx-meshctl_<version>_darwin_amd64.tar.gz` for Intel based Macs.

1. Unzip the binary.

    ```bash
    tar -xvf nginx-meshctl_<version>_darwin_amd64.tar.gz nginx-meshctl
    ```

1. Move the `nginx-meshctl` executable in to your PATH.

    ```bash
    sudo mv nginx-meshctl /usr/local/bin/nginx-meshctl
    ```

1. Ensure the `nginx-meshctl` is executable.

    ```bash
    sudo chmod +x /usr/local/bin/nginx-meshctl
    ```

1. Test the installation.

    ```bash
    nginx-meshctl
    ```

#### Install on Windows

1. Download the appropriate binary for your architecture, either `nginx-meshctl_<version>_windows_amd64.zip` or `nginx-meshctl_<version>_windows_arm64.zip`.

1. Extract the binary, `nginx-meshctl.exe`, from the zip file.

1. Add the binary to your PATH.

1. Test the installation.

    ```bash
    nginx-meshctl
    ```

### Air Gap Environment

NGINX Service Mesh will pull multiple required images into your Kubernetes cluster in order to function, some of which are from publicly-accessible third parties. For a full list refer to the [Technical Specifications]( {{< ref "/about/tech-specs.md#images" >}} ). If you are using a private registry, see our [private registry guide]( {{< ref "/guides/private-registry.md" >}} ).

## Install the NGINX Service Mesh Control Plane

{{< see-also >}}
Check out [Getting Started with NGINX Service Mesh]({{< ref "/get-started/configuration.md" >}}) to learn about the deployment options before proceeding.  
You can find the full list of options in the [`nginx-meshctl` Reference]( {{< ref "nginx-meshctl.md" >}} ).
{{< /see-also >}}

{{< important >}}
`nginx-meshctl` creates the namespace for the NGINX Service Mesh control plane.  
This namespace is dedicated to the NGINX Service Mesh control plane and **should not be used for anything else**.  
If desired, you can specify any name for the namespace via the `--namespace` argument, but do not create this namespace yourself.
{{< /important >}}

Take the steps below to install the NGINX Service Mesh control plane.

1. Run the `nginx-meshctl deploy` command using the desired [options]({{< ref "nginx-meshctl.md#deploy" >}}).

   For example, the following command will deploy NGINX Service Mesh using all of the default settings for the latest release:

    ```bash
    nginx-meshctl deploy
    ```

2. Verify the pods are running.

    ```bash
    $ kubectl get pods -n nginx-mesh
    NAME                                   READY   STATUS    RESTARTS   AGE
    nats-server-84f8b6f669-xszkc           1/1     Running   0          14m
    nginx-mesh-controller-954467945-sc7qh  1/1     Running   0          14m
    nginx-mesh-metrics-57464df46d-qskd2    1/1     Running   0          14m
    spire-agent-92ktv                      1/1     Running   0          15m
    spire-agent-9dbn6                      1/1     Running   0          15m
    spire-agent-z5cq6                      1/1     Running   0          15m
    spire-server-0                         2/2     Running   0          15m
    ```

    {{< note >}} If running in OpenShift, you will see two pods per Spire Agent container. {{< /note >}}

## UDP MTU Sizing

UDP traffic proxying is turned off by default. You can activate it at deploy time using the `--enable-udp` flag. Linux kernel 4.18 or greater is required.

NGINX Service Mesh automatically detects and adjusts the `eth0` interface to support the 32 bytes of space required for PROXY Protocol V2.
See the [UDP and eBPF architecture]({{< ref "architecture.md#udp-and-ebpf" >}}) section for more information.

NGINX Service Mesh does not detect changes made to the MTU in the pod at runtime.
If adding a CNI changes the MTU of the `eth0` interface of running pods, you should re-roll the affected pods to ensure those changes take place.

## Uninstall

{{< important >}}
OpenShift users: Before uninstalling, read through the [OpenShift considerations]({{< ref "/get-started/openshift-platform/considerations.md#remove" >}}) guide to make sure you understand the implications.
{{< /important >}}

Uninstalling does the following:

1. Removes the control plane and its contents from Kubernetes.
2. Deletes all NGINX Service Mesh traffic policies.

The `nginx-meshctl` command-line utility prints a list of resources that contain the sidecar proxies when the uninstall completes. You must re-roll the Deployments in Kubernetes to remove the sidecars. Until you re-roll the resources, the sidecar proxies still exist, but they don't apply any rules to the traffic.

### Uninstall the Control Plane

To uninstall the Service Mesh control plane using the `nginx-meshctl` command-line utility, run the command shown below.

```bash
nginx-meshctl remove
```

When prompted for confirmation, specify `y` or `n`.
If you want to skip the confirmation prompt, add the `-y` flag as shown in the example below.

```bash
nginx-meshctl remove -y
```

{{< note >}}
If the removal process gets stuck or fails to clean up all resources, you can manually view all NGINX Service Mesh resources that were left over using the following command:

```bash
kubectl api-resources --verbs=list -o name | xargs -n 1 kubectl get --show-kind --ignore-not-found -l app.kubernetes.io/part-of=nginx-service-mesh -A
```

These resources can be manually removed if necessary.
{{< /note >}}

### Remove the Sidecar Proxy from Deployments

If your resources support Rolling Updates (Deployments, DaemonSets, and StatefulSets), run the following `kubectl` command for each resource to complete the uninstall.

```bash
kubectl rollout restart <resource type>/<resource name>
```

For example:

```bash
kubectl rollout restart deployment/frontend
```

{{< note >}}
If you want to redeploy NGINX Service Mesh after removing it, you need to re-roll the resources after the new NGINX Service Mesh is installed. Sidecars from an earlier NGINX Service Mesh installation won't work with a new installation.
{{< /note >}}
