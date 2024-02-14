# coop-dev-ex 
[![App:main](https://github.com/stianSjoli/mean-mug-CI/actions/workflows/main_app.yml/badge.svg)](https://github.com/stianSjoli/mean-mug-CI/actions/workflows/main_app.yml)

## The Case 
The case given is to create a small developer platform that enables developers to deploy their containers with ease. The solution should follow GitOps principles, and by default it should be easy to use. However, complex configuration should be allowed, meaning that it should be allowed to add arbitrary Kubernetes resources. On commit it should run CI, build and publish the image, and update the deployment in kubernetes. Developers should not have direct access to the kubernetes cluster. The interesting part of the assignment is the developer experience, so avoid to much configuration of what is referred to as "supporting systems". The case gives these examples of supporting systems that can be used. 

Examples of supporting systems:
* Kubernetes: GCP GKE, k3s, kind, Minikube  
* Container registry: GCP artifact registry, GitHub Container Registry, Dockerhub
* Git: Github, GitLab  
* CI: Github Actions, GitLab CI, Dagger

It is also stated to consider using Dagger or Mage to demonstrate proficiency with Go. The case is open ended and can be freely interpreted with own assumptions and decisions.   

## Thought process - "start early fail early"
I was completely wrong about dagger from the start of the implementation. I thought it would completely remove Github action, and would then require alot more features (commit trigger, run on a defined vm/container, configuration parser etc.). And I thought about scaling/reuse/generalization way too early. I see now that a combination between mage, dagger and Github action is the way to go.   

### The safe solution 
The architecture that would be safest to implement based on time, and my previous experiences would consist of the following tech stack
* Kubernetes: GKE
* Container registry: GCP Artifact registry (or Dockerhub)
* Git: GitHub
* CI: GCP cloud build (not mentioned as an example)
* CD: GCP cloud deploy (uses skaffold, also not mentioned as an example)    

### The solution 
The architecture will based on the following tech stack
* Kubernetes: minikube on home network with inbound initiated requests denied (so no push from webhooks, only pulls from within my network) 
* Container registry: Dockerhub or minikube docker registry
* Git: GitHub
* CI: Dagger and Github Actions 
* CD: ArgoCD on minikube

#### What is missing 

* lots more testing
* more work on the triggering of CD (I thought it would work to make a k8 yaml parser and automatic commit on the repo to initiate a ArgoCD deployment)
* branching (publish on main (publish) and build on feature branches)  
* I think I would explore "adding arbitrary kubernetes resources" on a separate git repo or a separate folder from resources associated with the App deployment
* my solution assumes only sync from git repo and not from cluster to github (inline with GitOps)   

