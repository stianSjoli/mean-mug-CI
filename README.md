# mean-mug-CI 
[![App:main](https://github.com/stianSjoli/mean-mug-CI/actions/workflows/main_app.yml/badge.svg)](https://github.com/stianSjoli/mean-mug-CI/actions/workflows/main_app.yml)

## The Case 
This repo alows me to play around with go (and mage, dagger), ttl.sh, ArgoCD and aspects around CI. The ArgoCD server runs on a minikube-cluster. It was turned into a template so I can rapidly reuse elements for future repositories.   
   
### Remote Testing on repository
A commit in /App will initiate the CI, and if test and build step are green, it will lead to a deploy to a minikube cluster.

### Local Testing of CI (assumes mage install and docker deamon running) 
How to run test step: 
```
cd CI/builder
mage test 
```
How to run build step: 
```
cd CI/builder
mage build 
```
How to run deploy step: 
```
cd CI/builder
mage deploy "ArgoCD/deployment.yml"  https://github.com/stianSjoli/mean-mug-CI.git [GITHUB TOKEN with write permission here]  
```




   
