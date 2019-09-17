## ansible-nginx-lb-example

[Ansibleでnginxを入れてLoad Balancingする - sambaiz-net](https://www.sambaiz.net/article/239/)

```
$ vi conf/app/nginx.conf
$ make build
$ pip install --user ansible
$ ansible-galaxy install nginxinc.nginx
$ vi hosts
$ ansible all -i ./hosts -m ping 
$ ansible-playbook -i ./hosts playbook.yml
```
