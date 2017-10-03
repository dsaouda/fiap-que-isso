# fiap-queisso

Aplicação que tem como objetivo ser um joguinho para meu filho que adora brincar com logos, letras e números. Um simples joguinho com DEVOPS usando RancherOS, Racher, Jenkins, Docker, Gogs, MongoDB, Golang - Gin, PHP - Zend Expressive e VueJS.

# Arquitetura

![Arquitetura](https://raw.githubusercontent.com/dsaouda/fiap-queisso/master/docs/arquitetura.png)

## Entendendo a arquitetura (simplificado)

 1. rancher gerencia os hosts e containers
 2. Jenkins faz polling no gogs (git server), quando percebe uma munça envia um sinal de atualização ao rancher
 3. rancher remove container antigo e adiciona novo container (tudo foi passado pelo jenkins via API rancher v2)
 4. No host1 estão o rancher server, jenkins, gogs, mongodb e as aplicações frontend do fiap-queisso
 5. No host 2 estão os microservices api do fiap-queisso
 
 # Videos de como funciona
 
 [![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/u8OSpYwDPzQ/0.jpg)](https://www.youtube.com/watch?v=u8OSpYwDPzQ)
