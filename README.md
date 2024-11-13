## Objetivo do projeto
Este projeto tem como objetivo desenvolver um sistema de persistência e transmissão de dados de localização em tempo real, garantindo a disponibilidade, escalabilidade e eficiência no armazenamento e na transmissão desses dados. A solução utiliza uma arquitetura baseada em microserviços, com Kafka como broker de mensagens para garantir a entrega confiável e o processamento distribuído dos dados de localização, e um servidor em Go para gerenciar a transmissão desses dados via WebSocket.

### Tecnologias utilizadas
* **Go:** Para a implementação das APIs de Publisher e Consumer.
* **Kafka:** Para o gerenciamento do fluxo de mensagens em tempo real.
* **MySQL:** Para o armazenamento dos dados de localização mais recentes.
* **MongoDB:** Para o armazenamento do histórico dos dados de localização.
* **WebSocket:** Para transmissão em tempo real de dados de localização para os clientes.

#### Arquitetura do sistema
![Arquitetura](./assets/architecture.png "Arquitetura")
* **Publisher API:** Recebe os dados de localização das embarcações ou serviços de localização. Esses dados são publicados no Kafka, permitindo o processamento e armazenamento distribuído.
* **Kafka Broker:** Atua como um intermediário confiável para o fluxo de mensagens, armazenando os dados de localização recebidos do Publisher API e disponibilizando-os para consumo..
* **Consumer API:** Consome os dados de localização do Kafka e os armazena na bases de dados (MongoDB, MySQL). Além disso, a API serve os dados para interfaces de usuário.
* **MySQL:** Focado na transmissão em tempo real dos dados de localização para os clientes, o MySQL organiza e disponibiliza os dados mais recentes em uma estrutura relacional.
* **MongoDB:** Atua como a base de dados principal para o armazenamento a longo prazo dos dados de localização, contendo o histórico completo dos dados de localização recebidos das embarcações.

