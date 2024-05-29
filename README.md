# demo-kubebuilder

## Description
Let’s do a little exercise: we’ll build a simple foo operator which has no real use, except to demonstrate the capabilities of an operator.
Our operator will manage the CRD named PodFriend. We fetch the PodFriend resource that triggered the reconciliation request to get the name. Then, we list all the pods that have the same name as PodFriend. If we find one (or more), we update PodFriend's happy status to true, else we set it to false.

## Getting Started

### Prerequisites
- go version v1.21.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### Create an operator
`kubebuilder init --domain my.company --repo my.company/demo`

The `kubebuilder init --domain my.company --repo my.company/demo` command is used to initialize a new controller (operator) project in Kubernetes using Kubebuilder, which is a tool for building Kubernetes controllers and operators.

The meaning of each part of the command is explained in detail below:

`kubebuilder` startup
This is the main command to initialize a new project with Kubebuilder. Create the basic structure of the project, including the directories and files necessary to begin developing a controller.

`--domain my.company`
The --domain option specifies the domain to use in the API Group of the Custom Resource Definitions (CRDs) that you will create with this operator. This domain is used to construct the full name of the API group, for example, if you have a MyResource resource in the apps group, the full group could be myresource.apps.my.operator.

`--repo my.domain/tutorial`

The --repo option defines the Go module (the Go module path) to use for the project. This is important because Go uses the module name to manage project dependencies. In this case, my.domain/tutorial will be the name of the Go module. This also sets the base import path for your project, which is important for Go to find packages correctly.


#### Practical example
Let's break down what happens when you run this command:

Project Initialization: A directory and file structure is created that Kubebuilder uses to manage the development of the operator. This includes files like `go.mod`, `main.go` and directories like api and drivers.

Domain: The my.operator domain will be used to generate the full name of the API groups of your CRDs. For example, if you define a resource in the apps group, the group's distinguished name will be `apps.my.company`.

Repository/Go Module: The Go module my.domain/tutorial is set in the project's go.mod file, which sets up the Go development environment for your operator. This is crucial for dependency management and for Go tools to know how to import and compile your code.

#### Create a API

`kubebuilder create api --group operator --version v1 --kind PodFriend`

### Structure

![image](https://github.com/falconcr/kubebuilder-demo/assets/143828508/e772296f-5b53-47ec-aa59-4100e79729d1)

- api/: Directory where you define your CRDs and API versions.

- controllers/: Directory where you define the controllers that manage your resources.

- config/: Contains configuration files for deploying your operator to Kubernetes.

- go.mod: Go modules file that contains the module name my.company/demo and project dependencies.

- main.go: Entry point of the application, where the controller manager is configured and started.

### To test
- make manifests

- make install

- make run

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/demo-kubebuilder:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/demo-kubebuilder:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/demo-kubebuilder:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/demo-kubebuilder/<tag or branch>/dist/install.yaml
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

