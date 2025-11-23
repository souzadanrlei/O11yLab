# 📘 Tracing Pipeline Documentation

Este documento descreve a arquitetura e a configuração do pipeline de Tracing utilizado no projeto O11yLab, baseado em OpenTelemetry e integrado ao Grafana Cloud.
A solução foi construída utilizando:

- OpenTelemetry Operator (core)
- Collectors dedicados (traces, metrics, gateway)
- Recebimento via OTLP (gRPC + HTTP)
- Processamento com batch e memory_limiter
- Envio para Grafana Cloud via OTLP/HTTP

O objetivo é oferecer um pipeline escalável, padronizado e compatível com cloud-native para ingestão e análise de traces.

### 🏗️ Arquitetura do Tracing

Componentes:

- Gateway Collector
Recebe os traces das aplicações e distribui via LoadBalancer para o collector de traces.
- Tracing Collector
Processa, filtra e envia traces para o Grafana Cloud.
- Grafana Cloud (Tempo)
Armazena e indexa os spans para visualização no Grafana.

### 📝 Fluxo de Processamento

**1. Receivers**

Recebimento via OTLP (gRPC/HTTP):
Porta 4317 = gRPC
Porta 4318 = HTTP
Referência:
🔗 https://opentelemetry.io/docs/specs/otlp/

**2. Processors**
✔ memory_limiter

Evita estouro de memória no collector.
Docs:
🔗 https://opentelemetry.io/docs/collector/configuration/#memory_limiter
✔ batch
Agrupa spans antes de enviar, reduzindo carga e melhorando performance.
Docs:
🔗 https://opentelemetry.io/docs/collector/configuration/#batch

**3. Exporters**
✔ otlphttp/grafana_cloud
Envia traces para o Grafana Cloud Tempo via HTTP OTLP.
Referência oficial:
🔗 https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/
Referência Tempo:
🔗 https://grafana.com/docs/tempo/latest/

### 🔄 Pipeline de Tracing
1. Aplicações exportam traces via OTLP.
2. Gateway Collector recebe e balanceia entre os coletores.
3. Tracing Collector processa e aplica limites.
4. Spans são enviados ao Grafana Cloud (Tempo).
5. As traces podem ser consultadas no Grafana.

### 📚 Documentação Oficial Referenciada

**OpenTelemetry Collector**
🔗 https://opentelemetry.io/docs/collector/
🔗 https://opentelemetry.io/docs/collector/configuration/
🔗 https://opentelemetry.io/docs/kubernetes/operator/

**OTLP Protocol**
🔗 https://opentelemetry.io/docs/specs/otlp/

**Grafana Cloud — OTLP Ingestão**
🔗 https://grafana.com/docs/grafana-cloud/send-data/