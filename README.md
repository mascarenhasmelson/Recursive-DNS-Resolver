# Recursive DNS Resolver


This is a recursive DNS resolver written in Go. It queries domain names step by step, starting from root name servers, and follows the DNS chain until it resolves the requested domain.

### Features
* Recursive DNS resolution (starting from root servers)
* Queries authoritative DNS servers step by step
* Returns detailed query logs showing each resolution step
* Supports streaming output for real-time query progress
* Simple API to query domains via HTTP

### How It Works

* Selects a random root name server from the list.
* Queries the root server for the NS record of the target domain.
* Recursively follows the chain until it gets the final A or CNAME record.
* Outputs each step in real time using Flush(), so clients see updates instantly.

### Installation

1. Clone the repository

```
https://github.com/mascarenhasmelson/Recursive-DNS-Resolver.git
cd Recursive-DNS-Resolver
```
2. Install dependencies

```
go mod tidy
```
3. Build the project:

```
go build -o Recursive-DNS-Resolver
```
4. Run the server:

```
./dns-resolver
```
### Query a Domain

we can either choose curl or browser
```
curl http://localhost:8080/query?domain=example.com
```
or
```
http://localhost:8080/query?domain=example.com
```

When querying example.com, the tool prints the resolution steps:

```
Starting DNS Query for: example.com
Query 193.0.14.129 about example.com
Query 192.5.6.30 about example.com
Query 193.0.14.129 about a.iana-servers.net.
Query 192.5.6.30 about a.iana-servers.net.
Query 199.43.135.53 about a.iana-servers.net.
Query 199.43.135.53 about example.com
Final Resolved Records:
Received Answer: example.com.	300	IN	A	23.192.228.80
Received Answer: example.com.	300	IN	A	23.192.228.84
Received Answer: example.com.	300	IN	A	23.215.0.136
Received Answer: example.com.	300	IN	A	23.215.0.138
Received Answer: example.com.	300	IN	A	96.7.128.175
Received Answer: example.com.	300	IN	A	96.7.128.198
Query Completed.
```
