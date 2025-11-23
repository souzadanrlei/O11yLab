# Otel-metrics Collector para qualquer vendor

### 1. Visão geral
Esta documentação contém a configuração do **Opentelemetry Collector**(otel-metrics) para enviar métricas para o qualquer vendor. A configuração é aplicada por meiuo de um CRD OpentelemetryCollector, usando o Opentelemetry Operator do Kubernetes.

A base dessa configuração foi inspirada pelo [cookbook do Juracir krohling](https://github.com/jpkrohling/otelcol-cookbook/blob/main/grafana-cloud/otelcol.yaml) para grafana cloud e nas melhores praticas da comunidade Opentelemetry Operator.

### 2. Pré-requisitos

- Um cluster Kubernetes (versão compatível com o OpenTelemetry Operator)
- OpenTelemetry Operator instalado no cluster. Você pode instalar por Helm:
```
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update
helm install opentelemetry-operator open-telemetry/opentelemetry-operator \
  --set admissionWebhooks.certManager.create=true
```

Conta no Grafana Cloud e credenciais (ID de stack + API token) para autenticação via basic_auth.
- Para criar as credenciais abra o seu grafana cloud, clique em connections -> collector -> instal e gere seu token.

3. Arquitetura da Configuração

Esta instância do Collector (otel-metrics):
- Recebe métricas via protocolo OTLP (grpc e http).
- Utiliza autenticação básica (BasicAuth) para exportar para o Grafana Cloud.
- Exporta métricas para dois destinos:
    - otlphttp/grafana_cloud (Grafana Cloud)
    - debug (para debug local ou logs do Collector)

A pipeline configurada é apenas para métricas, sem traces ou logs (adaptável conforme necessidade, mas nesse caso usaremos apenas para métricas).

#### 4. Estrutura do CRD OpenTelemetryCollector

Aqui está a explicação das seções principais:
```
apiVersion: opentelemetry.io/v1beta1  
kind: OpenTelemetryCollector  
metadata:  
  name: otel-metrics  
spec:  
  replicas: 1                    # Número de réplicas do Collector
  mode: deployment               # Modo de deployment (pode ser daemonset, sidecar, statefulset) :contentReference[oaicite:2]{index=2}
  config:                         # Configuração do próprio Collector
    receivers:
      otlp:
        protocols:
          grpc:
            endpoint: "0.0.0.0:4317"   # Porta para OTLP gRPC
          http:
            endpoint: "0.0.0.0:4318"   # Porta para OTLP HTTP

       prometheus/internal:                            # Monitoramento do collector com o prometheus
        config:
          scrape_config:
            - job_name: otel-collector-internal
              static_config:
                - targets: ["0.0.0.0:8888"]

    processors: {}                # Nenhum processador adicional configurado

    extensions:
      basicauth/grafana_cloud:    # Extensão para autenticação BASIC
        client_auth:
          username: "$<user_token>"   # Substituir pelo seu ID de stack (usuário)
          password: "$<token>"        # Substituir pelo API token

    exporters:
      debug: {}                    # Exportador de debug (útil para logs do Collector)
      otlphttp/grafana_cloud:
        endpoint: https://otlp-gateway-prod-eu-west-2.grafana.net/otlp  # Endpoint do Grafana Cloud
        auth:
          authenticator: basicauth/grafana_cloud
        tls:
          insecure: false
          insecure_skip_verify: true   # Pode ser ajustado conforme necessidade

    service:
      extensions: [basicauth/grafana_cloud]   # Habilita a extensão no pipeline
      pipelines:
        metrics:
          receivers: [otlp]
          processors: []
          exporters: [otlphttp/grafana_cloud, debug]  # Exporta para Grafana + debug
```
### 5. Autenticação para Grafana Cloud

username: deve ser o ID da stack do Grafana Cloud (um número ou string fornecida por Grafana).
password: o API token gerado no Grafana Cloud, que deve ter permissão para ingestão OTLP. Em muitos exemplos oficiais, o token começa com glc_.
Essa autenticação é feita usando a extensão basicauth do OpenTelemetry Collector.

### 6. TLS e Segurança

insecure_skip_verify: true está ativado para ignorar verificação do certificado TLS, isso pode facilitar no início, mas não é recomendado para ambientes de produção, pois pode expor você a ataques de man-in-the-middle.

Se possível, prefira desativar o insecure_skip_verify e garantir que o Collector confie nos certificados do Grafana Cloud.

### 7. Escalabilidade e Resiliência

- Réplicas: Está definido replicas: 1. Se precisar de alta disponibilidade, você pode aumentar esse número.
- Modo: deployment foi escolhido para facilitar escalonamento, rollback e gestão como uma aplicação normal. Para cenários onde cada nó do Kubernetes deve ter um Collector, considere ty-elastic.github.io
- Extensibilidade: Você pode adicionar processadores (por exemplo, batch, memory_limiter) conforme requerido por volume ou performance.

### 8. Observabilidade do Collector

O debug exporter permite ver métricas e logs internos do Collector para troubleshooting.
Se quiser monitorar o próprio Collector via Prometheus, pode configurar um prometheus receiver e pipeline separados. A comunidade OpenTelemetry Collector Helm Chart tem exemplos disso.
O Operator pode criar Service / ServiceMonitor para o Collector automaticamente, se ele detectar ports nos receivers.

### 9. Exemplos de Uso

9.1 Deploy com kubectl apply
```
kubectl apply -f path/to/otel-metrics.yaml
```

### 10. Boas Práticas e Recomendações

- Segurança: Guarde o token do Grafana Cloud em um Kubernetes Secret e monte como variáveis de ambiente no Collector, ao invés de deixar hard-coded no YAML.
- Escalabilidade: Ajuste replicas conforme a carga esperada de métricas.
- Processadores: Avalie adicionar batch ou memory_limiter para otimizar o envio e uso de memória.
- Observabilidade do Collector: Habilite métricas de auto-observabilidade para monitorar o desempenho do Collector (latência, uso de memória, erros).
- Validação de config: Antes de aplicar em produção, valide seu YAML de Collector localmente com a imagem do Collector (para pegar erros de sintaxe ou estrutura).
- CI/CD: Use GitOps para gerenciar essa configuração (por exemplo, armazene o manifest no Git e aplique via pipeline).
