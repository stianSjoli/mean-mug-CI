# coop-dev-ex

## The Case 
The case given is to create a small developer platform that enables developers to deploy their containers with ease. The solution should follow GitOps principles, and by default it should be easy to use. However, complex configuration should be allowed, meaning that it should be allowed to add arbitrary Kubernetes resources. On commit it should run CI, build and publish the image, and update the deployment in kubernetes. Developers should not have direct access to the kubernetes cluster. The interesting part of the assignment is the developer experience, so avoid to much configuration of what is referred to as "supporting systems". The case gives these examples of supporting systems that can be used. 

Examples of supporting systems:
* Kubernetes: GCP GKE, k3s, kind, Minikube  
* Container registry: GCP artifact registry, GitHub Container Registry, Dockerhub
* Git: Github, GitLab  
* CI: Github Actions, GitLab CI, Dagger

It is also stated to consider using Dagger or Mage to demonstrate proficiency with Go. The case is open ended and can be freely interpreted with own assumptions and decisions.   

## Thought process 
How the system should follow "GitOps principles", and also what developers associate with "GitOps" vs "DevOps", can be an interesting topic in itself. "GitOps" is now commonly associated with using Git for versioning infrastructure code ("Infrastructure as code") and Git pull requests to manage and deploy infrastructure. Thes terms, "DevOps" and "GitOps" might have been closer in their earlier days, but had a divergent evolution as a consequence of more automation and evolution of platform tools. Thus, I feel "GitOps" would be more in the configuration and deployment of the supporting systems. The case also states that the platform should enable developers to deploy their containers with ease, which is in the "DevOps" domain. In that interpretation, the proposed solution will not follow "GitOps" principles to a large extent as supporting system setup will not be automated. However, a "GitOps"-oriented statement is that complex configuration should be allowed, meaning that it should be allowed to add arbitrary Kubernetes resources. Also, some focus on the fact that a GitOps principle is git as the single truth of change (for the platform), and my suggested solution will follow that principle.

### The safe solution 
The architecture that would be safest to implement based on time, and my previous experiences would consist of the following tech stack
* Kubernetes: GKE
* Container registry: GCP Artifact registry (or Dockerhub)
* Git: GitHub
* CI: GCP cloud build (not mentioned as an example)
* CD: GCP cloud deploy (uses skaffold, also not mentioned as an example)    

### The solution 
The architecture will based on the following tech stack
* Kubernetes: minikube on home network with inbound initiated requests denied (so no push from webhooks, only pulls) 
* Container registry: Dockerhub or minikube docker registry
* Git: GitHub
* CI: Dagger
* CD: ArgoCD on minikube

#### Mage
I come from a C and OS heavy bachelor degree (and was scheduled as a teaching assistant for OS course at UIO). Thus, makefiles are something I only associate with (efficient (dependent on how good the makefile was)) (re-)compiling and deleting software binaries for low level languages like C. Makefiles are hard to read and hard to write (from Mage). Mage could be used with their description of targets, and Mage as a CLI tool - but the origin story seems odd for my association with makefiles. 

#### Skaffold
I might have tested skaffold if the documentation was better. 

