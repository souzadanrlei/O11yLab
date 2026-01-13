# 🌐 O11yLab

Laboratório de Observabilidade com OpenTelemetry e Ferramentas Open Source

O O11yLab é um projeto colaborativo criado para servir como um laboratório unificado de observabilidade moderna, utilizando OpenTelemetry, Prometheus, Loki, Tempo, Grafana, além de ferramentas complementares para monitoramento, instrumentação, rastreamento distribuído e métricas em larga escala.

Nosso objetivo é construir um ambiente educacional e experimental que facilite o aprendizado, a implementação e a comparação de tecnologias de observabilidade de mercado, seguindo boas práticas inspiradas em padrões da indústria e referências de documentação como o Google Cloud Documentation Style Guide.

### 🎯 Objetivo do Projeto

O projeto O11yLab foi criado para:

- Demonstrar arquiteturas reais de observabilidade modernas.
- Oferecer um ambiente modular, extensível e fácil de replicar.
- Mostrar na prática como usar o OpenTelemetry Collector, instrumentação automática, pipelines de métricas, logs e traces.
- Integrar diversas ferramentas de monitoramento para fornecer um ambiente completo.
- Servir como laboratório, material de estudo, referência técnica e repositório de experimentos.
- Criar uma base sólida de conteúdo para documentação, tutoriais e vídeos educacionais.

### 🏗️ Visão Técnica Inicial

A primeira versão do O11yLab possui:

📡 OpenTelemetry Operator — responsável por gerenciar coletores e instrumentações.
🛰️ Gateway Collector — centraliza a entrada de métricas, logs e traces.
📊 Métricas — coletadas via OTLP e distribuídas por um loadbalancer interno.
📁 Logs — coletados e enviados para análise e indexação.
🔍 Traces — suportando pipelines completos para rastreamento distribuído.
🧩 Arquitetura modular — diretórios separados para cada pipeline (metrics, logs, traces).

### 📦 Helm — templates, operadores e componentes instaláveis.

Conforme o laboratório evoluir, serão adicionados:

Grafana 
VictoriaMetrics
Prometheus
Loki
Tempo
Blackbox Exporter
k6 para testes
Ferramentas []
Dashboards e guias de instrumentação

### 📁 Estrutura do Repositório (versão inicial)

![arquitetura](src/img/
2026-01-13_14-45.png)


O11yLab/
├── helm/
│   └── README.md
├── opentelemetry/
│   ├── README.md
│   ├── otel-gateway.yaml
│   ├── metrics/
│   │   └── otel-metrics.yaml
│   ├── logs/
│   │   └── otel-logs.yaml
│   └── traces/
│       └── otel-traces.yaml
└── README.md  ← documentação principal

### 👥 Colaboradores e Responsabilidades

**[Danrlei Souza — Software Engineer Jr. / Project Collaborator](https://github.com/souzadanrlei)**

Idealização do projeto O11yLab
Suporte na estruturação do laboratório
Configuração do OpenTelemetry Operator
Integração com ferramentas externas
Criação do Gateway Collector
Documentação técnica e conteúdos educacionais
Integrações (Grafana, Zabbix, Datadog, Rew Relic)

**[Ezequiel — Software Engineer  Sr. - Tech Lead / Project Collaborator](https://github.com/ezeqme)**

Idealização do projeto O11yLab
Arquitetura de observabilidade
Configuração dos ambientes
Suporte na estruturação do laboratório
Configuração do OpenTelemetry Operator
Integração com ferramentas externas
Contribuições técnicas em pipelines de logs, métricas e traces

Este projeto é colaborativo e aberto a contribuições futuras conforme novos módulos forem desenvolvidos.

### 📘 Padrões de Documentação

Este projeto segue diretrizes inspiradas no Google Cloud Documentation Style Guide, priorizando:

- Clareza
- Concisão
- Estrutura consistente
- Exemplos práticos
- Fluxos orientados a tarefas
- Mantenabilidade e expansão futura

### 🚀 Status do Projeto

Fase atual:
✔️ Setup inicial do OpenTelemetry
✔️ Operator instalado
✔️ Gateway Collector configurado
✔️ Primeiros pipelines de métricas, logs e traces criados
⬜ Documentação avançada
⬜ Deploy de ferramentas adicionais
⬜ Diagramas de arquitetura
⬜ Scripts de automação
⬜ CI/CD para validação YAML

### 🤝 Como contribuir

Em breve adicionaremos um CONTRIBUTING.md com:
- Nomeação de branches
- Padrões de commit
- Estrutura para PRs
- Boas práticas para novos módulos
