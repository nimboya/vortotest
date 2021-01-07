## Deploying the App to Kubernetes

1. Setup a Kubernetes Cluster (EKS, AKS or GKE or Local)

2. Ensure kubectl is installed.

3. Ensure that the kubeconfig file is located in the `.kube/config` path for authentication and authorization of every `kubectl` request.

4. Clone this repository with the command `git clone https://github.com/nimboya/vortotest`

5. Build the Docker Image for the Go application with the command `docker build -t nimboya/vortotest`

6. Push the image with the `docker push nimboya/vortotest` command 
N.B. Replace `nimboya` with your own Dockerhub login credentials

7. Ensure the Kubernetes image value in the `kube.yml` is the image that has been pushed, for this scenario, we shall use our image name `nimboya/vortotest`.

8. The `kube.yml` contains both the `Deployment` and `Service` objects, running the deployment with create the application and the service as a Load Balancer type

8. To deploy, run the kubectl command `kubectl apply kube.yml`.

