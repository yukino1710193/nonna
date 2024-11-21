# Nonna - Нонна - ノンナ

### (Blizzard Nonna - Bão Tuyết Nonna - ブリザードのノンナ)

[![release](https://img.shields.io/badge/nonna--v0.1-log?style=flat&label=release&color=darkgreen)]()
[![LICENSE](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

[![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white&link=https%3A%2F%2Fkubernetes.io)](https://kubernetes.io/)
[![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)]()
[![Knative](https://img.shields.io/badge/knative-log?style=for-the-badge&logo=knative&logoColor=white&labelColor=%230865AD&color=%230865AD)](https://knative.dev/docs/)
[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Protobuf](https://img.shields.io/badge/Protobuf-log?style=for-the-badge&logo=nani&logoColor=green&labelColor=red&color=darkgreen)](https://protobuf.dev/)

`Nonna` is the Queue Modifier Module of the [ikukantai Fleet](https://github.com/bonavadeur/ikukantai)

![](docs/images/nonna_wp.jpg)

## 1. Motivation

`nonna` supports deploying Queue-Modifying Algorithm in `ikukantai` Fleet without deeping into Knative source code.

`nonna` is actually **Knative Queue-Proxy** underneath. In the Vanilla Knative Serving, The queueing model implemented in Queue-Proxy is FIFO Queue. Although FIFO is a simple queueing model, it is not the most optimal queueing model in some complicated scenarios. By using `nonna`, you can implement your own priority queue that adopts parameters such as: **HTTP method**, **URI path**, **HTTP Header**, **Source IP address** and **Domain name**.

`nonna` also supports piggybacking based Load Balancer in the `ikukantai` Fleet along with [katyusha](https://github.com/bonavadeur/katyusha) by modifying HTTP header of responsed packets.

The name `nonna` is inspired by the character **Nonna** in the anime **Girls und Panzer**. `nonna` and [katyusha](https://github.com/bonavadeur/katyusha) form a complete Load Balancing system for the `ikukantai` Fleet. This Load Balancing system uses piggybacking mechanism to update load status as fast as possible, much like how **Nonna** always carries **Katyusha** on her back in anime **Girls und Panzer**.

## 2. Structure

![](docs/images/nonna-structure.jpg)

## 3. Installation

## 4. Development

## 5. Author

Đào Hiệp - Bonavadeur - ボナちゃん  
The Future Internet Laboratory, Room E711 C7 Building, Hanoi University of Science and Technology, Vietnam.  
未来のインターネット研究室, C7 の E ７１１、ハノイ百科大学、ベトナム。  

![](docs/images/github-wp.png)
