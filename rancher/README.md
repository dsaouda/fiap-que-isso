# instalando o rancher os

Seguir guia de instalação em http://rancher.com/docs/os/v1.1/en/running-rancheros/server/install-to-disk/
O rancher os pode ser baixado no link https://github.com/rancher/os/releases/download/v1.1.0/rancheros.iso

## Passo a passo simplificado

1. iniciar o virtual box com a iso do rancher e uma placa de rede com acesso a internet e outra com acesso para ssh
2. depois de inciado altere a senha do rancher para acessar utilizando putty, para isso digite `sudo passwd` e escolha a senha
3. acesse o rancher via ssh pelo putty ou outra ferramenta de preferência
4. criar um arquivo chamado cloud-config.yml (existe um arquivo de exemplo aqui https://github.com/dsaouda/fiap-queisso/blob/master/rancher/cloud-config.yml e aqui mais informações http://rancher.com/docs/os/v1.0/en/configuration/#cloud-config)
5. gere um ssh-keygen e adicione a chave pública na sessão ssh_authorized_keys do cloud-config.yml e a chave privada será usada para acessar o sistema
6. valide o arquivo gerado `sudo ros config validate -i cloud-config.yml` http://rancher.com/docs/os/v1.0/en/configuration/#validating-a-configuration-file
7. instalando no disco `sudo ros install -c cloud-config.yml -d /dev/sda` http://rancher.com/docs/os/v1.0/en/running-rancheros/server/install-to-disk/#installing-rancheros-to-disk

## Instalando o rancher gerenciador de container

1. depois de instalando o rancher os, acesse via SSH apontando para sua chave privada (`ssh -i <chave> <usuario>@<host>`)
2. execute `docker run -d --restart=unless-stopped -p 8080:8080 rancher/server`

## gerando chave pública e privada com ssh-keygen

1. `ssh-keygen -t rsa -b 4096 -C "<um-email@teste.com.br>"`
2. `Enter file in which to save the key (/c/Users/Diego.Saouda/.ssh/id_rsa): rancher_id_rsa`
3. `Enter passphrase (empty for no passphrase):`
4. `Enter same passphrase again:`
5. Copiar chave rsa para colocar no arquivo cloud-config.yml `cat rancher_id_rsa.pub`
