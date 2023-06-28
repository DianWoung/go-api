最简的nginx应用

创建namespace
~~~~
kubectl apply -f namespace.yaml
~~~~
创建service
~~~~
kubectl apply -f service.yaml
~~~~
创建default的configmap (实验性质，配置nginx目前还是用定制dockerfile镜像比较合适)
~~~
kubectl create configmap nginx-config --from-file=conf/default.conf -n ns-test
~~~
创建deployment
~~~~
kubectl apply -f deployment.yaml
~~~~