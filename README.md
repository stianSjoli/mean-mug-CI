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

### Folders 

* /platform: "owner" of this folder should be platform and yaml for more complex kubernetes resources should be commited here 
* /ArgoCD: "owner" is the team or preferable reserved for automatic commits. Some commits will be done by team on services and other kubernetes resources connected to the application 
* /CI: the dagger CI scripts together with modules for updating the kubernetes manifest with the image ref and ability to commit programmatically on the repo
* /App: a simple application to work on to show that the pipelines forfills image build, publish and deploy to minikube cluster

### Remote Testing on repository
A commit in /App will initiate the CI, and if test and build step are green, it will lead to a deploy to a minikube cluster.

### Local Testing of CI (assumes mage install and docker deamon running) 
How to run test step: 
```
cd CI/builder
mage testApp 
```
How to run build step: 
```
cd CI/builder
mage buildApp 
```
How to run deploy step: 
```
cd CI/builder
mage deployApp "ArgoCD/deployment.yml"  https://github.com/stianSjoli/mean-mug-CI.git [GITHUB TOKEN with write permission here]  
```




   
