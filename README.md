# mean-mug-CI 
[![App:main](https://github.com/stianSjoli/mean-mug-CI/actions/workflows/main_app.yml/badge.svg)](https://github.com/stianSjoli/mean-mug-CI/actions/workflows/main_app.yml)

This repo allows me to play around with go (and mage, dagger), ttl.sh, ArgoCD and aspects around CI. The ArgoCD server runs on a minikube-cluster. It was turned into a template so I can rapidly reuse elements for future repositories.   
   
### Remote Testing on repository
A commit in /App will initiate the CI, and if test and build step are green, it will lead to a deploy to a minikube cluster.

### Local Testing of CI (assumes mage install and docker deamon running) 
How to run CI step: 
```
cd CI/builder
mage ci 
```
How to run deploy step: 
```
cd CI/builder
mage cd [GITHUB TOKEN with write permission here]  
```




   
