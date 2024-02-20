# mean-mug-CI 
[![App:main](https://github.com/stianSjoli/mean-mug-CI/actions/workflows/main_app.yml/badge.svg)](https://github.com/stianSjoli/mean-mug-CI/actions/workflows/main_app.yml)

## The Case 
The case given is to create a small developer platform that enables developers to deploy their containers with ease. The solution should follow GitOps principles, and by default it should be easy to use. However, complex configuration should be allowed, meaning that it should be allowed to add arbitrary Kubernetes resources. On commit it should run CI, build and publish the image, and update the deployment in kubernetes. Developers should not have direct access to the kubernetes cluster. The interesting part of the assignment is the developer experience, so avoid to much configuration of what is referred to as "supporting systems". The case gives these examples of supporting systems that can be used. 

Examples of supporting systems:
* Kubernetes: GCP GKE, k3s, kind, Minikube  
* Container registry: GCP artifact registry, GitHub Container Registry, Dockerhub
* Git: Github, GitLab  
* CI: Github Actions, GitLab CI, Dagger

It is also stated to consider using Dagger or Mage to demonstrate proficiency with Go. The case is open ended and can be freely interpreted with own assumptions and decisions.   

### The solution 
The architecture will based on the following tech stack
* Kubernetes: minikube
* Container registry: ttl.sh ("easy reference for unathorised pulls, but ok to use now for illustration purposes")
* Git: GitHub
* CI: Dagger, mage and Github Actions 
* CD: ArgoCD on minikube


   
