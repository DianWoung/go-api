创建dashboard
~~~~
kubectl apply -f recommended.yaml
~~~~
创建账户
~~~~~
kubectl apply -f dashboard-adminuser.yaml
~~~~~

赋予权限
~~~~
kubectl apply -f cluster-role.yaml
~~~~

获取用户token

~~~~
kubectl -n kubernetes-dashboard create token admin-user
~~~~