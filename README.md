# Serverless/Cloud-Native watermark-service

A global publishing company that publishes books and journals wants to develop a service to
watermark their documents. Book publications include topics in business, science and media. Journals don’t include any specific topics. A document (books, journals) has a title, author and a watermark property. An empty watermark property indicates that the document has not been watermarked yet.

The watermark service has to be asynchronous. For a given content document the service should return a ticket, which can be used to poll the status of processing (e.g.: Status: Started, Pending, Finished). If the watermarking is finished the document can be retrieved with the ticket. The watermark of a book or a journal is identified by setting the watermark property of the object. For a book the watermark includes the properties content, title, author and topic. The journal watermark includes the content, title and author.

##### Examples for watermarks:
{content:”book”, title:”The Dark Code”, author:”Bruce Wayne”, topic:”Science”}

{content:”book”, title:”How to make money”, author:”Dr. Evil”, topic:”Business”}

{content:”journal”, title:”Journal of human flight routes”, author:”Clark Kent”}

#### Tasks
Implement the Watermark-Service using RESTful APIs in microservice-oriented model, meeting the above conditions.

1. A microservice for watermarking and returning the ticket-id API and user authorization API. Users allowed: SuperAdmin, Default
2. A microservice for retrieving a book by its ticket-id when finished.
3. An API for retrieving the statuses by its ticket-id

#### Note
Provide sufficient Unit-Tests to ensure the functionality of the service by giving enough logging output monitoring various (10 books, 10 Journals)
asynchronous watermark processes identified by a unique ticket-id.

Use Golang/gRPC stack, MongoDB

A test script wrapper in the root directory running the tests, like runWatermarkTests.sh

Setup a local development environment of your choice based on Kubernetes (Minikube, Docker for Desktop, OKD, K3s, microk8s, etc..)

### Nice To Haves:

Create test scenarios for following typical microservice patterns:
- Ambassador API Gateway (Facade Pattern)
- Circuit Breaker
- Mirroring
- Canary Deployments
- Alerting with Grafana

###### References and samples:
See impressive sample for some of these service mesh features here
https://github.com/DerSalvador/demo-mesh-arena

Knative CloudEvents Sample here
https://github.com/DerSalvador/eventing-contrib.git

##### Conribution:
Issue Tracker: https://trello.com/b/tuDO5DYR/watermark-service-application
