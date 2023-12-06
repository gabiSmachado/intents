# UMA SOLUÇÃO PARA REDES AUTO-ORGANIZADAS BASEADAS EM INTENÇÕES
Este trabalho tem o objetivo de criar redes 5g auto-organização da rede através de Intents, desenvolvido em cima da arquitetura da rede disponível pela O-RAN Alliance em https://docs.o-ran-sc.org/en/latest/architecture/architecture.html.

## Arquitetura

O Intent é definido pelo operador, através do intent interface que lê as solicitações do usuário, que são encaminhas para o Intent Receiver.

O Intent Receiver executa as operações de criação, listagem e visualização da descrição de intents, assim como deletá-los, através de uma interface REST com dados JSON. Quando um intente é criado o Intent Receiver verifica a sanidade e, se sã, é salvo em uma DataBase e publicado no Message Queue para consumo do Intent Broker.

O Intent Broker consome as mensagens e as traduz em em ações de controle que devem ser tomadas pelo near-RT RIC, publicando-as na interface A1 como políticas, que serão distribuídas para as xApps do near-RT RIC através do Mediador A1.
![Alt text](/arqu.png)

## Requisitos 
Linguagem GO;
Mariadb: https://www.digitalocean.com/community/tutorials/how-to-install-mariadb-on-ubuntu-20-04-pt
Apache Kafka: https://kafka.apache.org/quickstart

Instalação do near-RT RIC: https://github.com/gabiSmachado/Near-RT-RIC_deploy
