	
	para acessar o rancher via SSH do putty

ifconfig
	para ver o IP para acesso	

ssh rancher@192.168.99.100
	acessando o rancher

ssh-keygen -t rsa -b 4096 -C "dsaouda@gmail.com"


ssh-keygen -t rsa -b 4096 -C "dsaouda@gmail.com"
Generating public/private rsa key pair.
Enter file in which to save the key (/c/Users/Diego.Saouda/.ssh/id_rsa): rancher_id_rsa
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in rancher_id_rsa.
Your public key has been saved in rancher_id_rsa.pub.
The key fingerprint is:
SHA256:8vgT9HXup1dhb3V8XkeUJdqnSDmfaUTT2FypCgt2HiM dsaouda@gmail.com

cat rancher_id_rsa.pub
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCgDEqdveR1v6cBbSwU1ej69A86QJqbbFC3ZH074w7F6naqDLaF783HPFVOiIzcjtWBG6ohDrpZ6vPr4jgamxOL9kb8ZGRhTqYDsC3t5tivmCknrnYCMADPXLKvhGnfFFbq65WvuuRwT40cmoTQzW5/zd+KgP7yldKDFeFCIPyz3N9sZRyr/YAsluMVXP5BKFOuH1zn2BLAcUPSsHBPaCTcFajB7qAQNMR+gWubXfV8RUr4dp1OI4t33IeAZkEhbHhQieivkz1edFuEgaDDo/qaVJp+g5I2rY+asxbOIfJMSue6C5OnB0/GuxEmiO8e1ktgEmkYdz0N1qiN8ybPBdhwzGZepXhn8kDkIwHwDSdNXvbPJQdy/CrZ2fa6m4bUcsFhsBjw//Ke8yv6TJQv0OTUX4BsoScET+dWM91aoIiUM71K2nULYR4Vg6h0pIsrXT6l0wWXWL7W3uv8WVjHtTeNLBYKiKfhb7rqs+ws3NnNrLd3WO7fBmKd1hiRAoDiF8Nd7seh3baM6ORgBiBVnHutTQWQaUOX3H7PfsLLDk07cLQgDZoMwKNtsyDhCD492Di0RTY8WIrMNIgAsLtez0SEAtCyFPc5eolpuDaPwdMUAO1SYprIny9kCBehhKAhe6QgIeQmC7PvTffH08epp722gmnBfTMcfVNz9rYq3prgGw== dsaouda@gmail.com


http://rancher.com/docs/os/v1.0/en/configuration/#cloud-config
vi cloud-config.yml

http://rancher.com/docs/os/v1.0/en/configuration/#validating-a-configuration-file
sudo ros config validate -i cloud-config.yml


http://rancher.com/docs/os/v1.0/en/running-rancheros/server/install-to-disk/#installing-rancheros-to-disk
sudo ros install -c cloud-config.yml -d /dev/sda

sudo ros config set rancher.network.interfaces.eth1.address "192.168.200.10/24"
sudo system-docker restart network


docker run -d --restart=unless-stopped -p 8080:8080 rancher/server