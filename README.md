# UMA SOLUÇÃO PARA REDES AUTO-ORGANIZADAS BASEADAS EM INTENÇÕES
This repository contains the O-RAN Load balancing app. This application takes in intents at the SMO level and delivers policies to the RAN to self-organize the network.

## Arquitetura

O Intent é definido pelo operador, através do intent interface que lê as solicitações do usuário, que são encaminhas para o Intent Receiver.

O Intent Receiver executa as operações de criação, listagem e visualização da descrição de intents, assim como deletá-los, através de uma interface REST com dados JSON. Quando um intente é criado o Intent Receiver verifica a sanidade e, se sã, é salvo em uma DataBase e publicado no Message Queue para consumo do Intent Broker.

O Intent Broker consome as mensagens e as traduz em em ações de controle que devem ser tomadas pelo near-RT RIC, publicando-as na interface A1 como políticas, que serão distribuídas para as xApps do near-RT RIC através do Mediador A1.
![Alt text](/arqu.png)

## Requisitos 
Instalação do near-RT RIC: https://github.com/gabiSmachado/Near-RT-RIC_deploy

