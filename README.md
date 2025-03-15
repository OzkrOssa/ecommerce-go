# Learn Microservices with Golang

In this project, I'm focused on distributed microservices using Golang and various tools.

## Technologies

1. **Configuration Servers**  
   [HashiCorp Vault](https://www.vaultproject.io/) as the configuration server to manage and centralize the configuration of our applications.

2. **Discovery Server**  
   [Consul](https://www.consul.io/) to register and discover our services. Instead of Eureka, Consul provides a robust solution for service discovery.

3. **API Gateway**  
   API Gateway with [Kong](https://konghq.com/), a high-performance API gateway that acts as an intermediary between clients and microservices.

4. **Asynchronous Communication with Kafka**  
   For asynchronous communication between microservices, we will use [Apache Kafka](https://kafka.apache.org/) as a high-performance messaging system.

5. **Synchronous Communication with gRPC**  
   [gRPC](https://grpc.io/) for synchronous communication between microservices.

6. **Protocol Buffers with Buf**  
   [Buf](https://buf.build/) to simplify Protocol Buffers development, enforce best practices, and improve gRPC API consistency across microservices.

7. **Distributed Tracing with Jaeger**  
   [Jaeger](https://www.jaegertracing.io/) for distributed tracing, which helps us monitor and trace requests across microservices and detect bottlenecks.

8. **Containerization with Docker and Docker Compose**  
   Set up the entire infrastructure using [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/), making it easy to create, deploy, and manage containers.  

9. **Orchestration with Kubernetes using Red Hat Developer Sandbox**  
   Deploy and manage containerized applications using [Red Hat OpenShift Developer Sandbox](https://developers.redhat.com/developer-sandbox), a free managed Kubernetes environment. This enables seamless scaling, networking, and resource management for microservices.
