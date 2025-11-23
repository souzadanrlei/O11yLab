# 📘 OpenTelemetry Gateway Collector — Deployment Mode

Este documento descreve como criar um Gateway Collector utilizando o OpenTelemetry Operator em Deployment Mode.
Esse modo é recomendado quando você precisa operar o Collector como uma aplicação independente, com controle de escalabilidade, versionamento e gerenciamento mais flexível.


### 🚀 Visão Geral

O Deployment Mode é ideal quando você deseja:
Escalar horizontalmente o Collector
Ter controle independente sobre a aplicação
Realizar rollbacks de versão
Parar / iniciar a instância do Collector de forma isolada
Tratar o Collector como um componente autônomo da arquitetura
Este modo é amplamente utilizado para criar um Gateway Collector, que atua como um ponto centralizador para ingestão de métricas, logs e traces antes de enviá-los para destinos externos ou para coletores do tipo agent.

### 📦 Pré-requisitos
Antes de começar, certifique-se de ter:
Kubernetes 1.21+
Helm instalado
OpenTelemetry Operator instalado no cluster

### 🏗️ Criando o Gateway Collector (Deployment Mode)

O exemplo abaixo cria um Collector do tipo Gateway, que recebe dados via OTLP (gRPC e HTTP), processa com batch, e exporta usando um exporter loadbalancing, configurado para detectar automaticamente instâncias de Collectors usando DNS.

📄 Arquivo: otel-gateway.yaml
```
apiVersion: opentelemetry.io/v1beta1
kind: OpenTelemetryCollector
metadata:
  name: otel-gateway
spec:
  mode: deployment
  replicas: 1

  config:
    receivers:
      otlp:
        protocols:
          grpc:
            endpoint: "0.0.0.0:4317"
          http:
            endpoint: "0.0.0.0:4318"

    processors:
      batch:

    exporters:
      debug: {}

      loadbalancing:
        protocol:
          otlp:
            tls:
              insecure: true
        resolver:
          dns:
            hostname: otel-metrics-collector-headless.opentelemetry-operator-system.svc.cluster.local
            port: 4317

    service:
      pipelines:
        metrics:
          receivers: [otlp]
          processors: [batch]
          exporters: [loadbalancing]

        logs:
          receivers: [otlp]
          processors: [batch]
          exporters: []

        traces:
          receivers: [otlp]
          processors: [batch]
          exporters: []
```

### ▶️ Aplicando no cluster

Execute:

```
kubectl apply -f otel-gateway.yaml
```
Verifique o Deployment:
```
kubectl get pods -l app.kubernetes.io/name=otel-gateway
```

### ⚙️ Explicação dos Componentes
**Receivers**
O Collector aceita dados via:
- OTLP gRPC (porta 4317)
- OTLP HTTP (porta 4318)
**Processors**
batch: otimiza processamento e envio dos dados.

**Exporters**
debug: útil para validação local
loadbalancing: distribui traces/metrics para collectors agentes usando DNS

**Pipelines**
São definidas três pipelines:
| Pipeline | Receivers | Processors | Exporters     |
| -------- | --------- | ---------- | ------------- |
| Metrics  | otlp      | batch      | loadbalancing |
| Logs     | otlp      | batch      | —             |
| Traces   | otlp      | batch      | —             |

### 📚 Documentação oficial

Para mais detalhes sobre o Deployment Mode:
👉 [OpenTelemetry Helm Charts – Operator](https://github.com/open-telemetry/opentelemetry-helm-charts/tree/main/charts/opentelemetry-operator)
(Seção: [Deployment Mode](https://github.com/open-telemetry/opentelemetry-helm-charts/tree/main/charts/opentelemetry-operator#daemonset-mode))