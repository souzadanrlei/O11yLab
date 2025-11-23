# 🟣 OpenTelemetry Operator – Guia de Instalação no Rancher Desktop

Este guia tem como objetivo ajudar você a instalar o OpenTelemetry Operator usando o Helm dentro do Rancher Desktop.

### ✅ Requisitos

Antes de iniciar, certifique-se de ter instalado:

Rancher Desktop
 👉 https://rancherdesktop.io/

### 📘 Documentação Oficial

Durante este laboratório, iremos seguir as instruções presentes na documentação oficial:

OpenTelemetry Operator Helm Chart
👉 https://github.com/open-telemetry/opentelemetry-helm-charts/tree/main/charts/opentelemetry-operator

### ⚠️ Atenção importante
A versão 0.58.0 do Helm Chart introduz uma nova versão do CRD OpenTelemetryCollector (v1beta1).
Leia as instruções de upgrade:
https://github.com/open-telemetry/opentelemetry-operator/blob/main/docs/crd-changelog.md#opentelemetrycollectoropentelemetryiov1beta1

E também o guia de upgrade do Helm Chart:
https://github.com/open-telemetry/opentelemetry-helm-charts/blob/main/charts/opentelemetry-operator/UPGRADING.md#0560-to-0570

### 🔐 Configuração de Certificado TLS para o Webhook

O Operator utiliza webhooks, e o Kubernetes precisa confiar no certificado TLS usado.

Você pode configurar isso:
1. Usar cert-manager (recomendado)
Ativar: admissionWebhooks.certManager.enabled=true
Cert-manager cria automaticamente certificados self-signed
Docs cert-manager:
https://cert-manager.io/docs/
https://cert-manager.io/docs/releases/


### 📦 Adicionar o Repositório Helm
```
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update
```

### 🚀 Instalando o OpenTelemetry Operator
Instalação simples (namespace padrão)
```
helm install opentelemetry-operator open-telemetry/opentelemetry-operator \
  --set "manager.collectorImage.repository=ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-k8s"
```
Instalação em namespace customizado

Use esta opção se planeja organizar melhor seus componentes.
```
helm install opentelemetry-operator open-telemetry/opentelemetry-operator \
  --namespace opentelemetry-operator-system \
  --create-namespace \
  --set "manager.collectorImage.repository=ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-k8s"
```
Instalação com certificado self-signed automático (sem cert-manager)
```
helm install opentelemetry-operator open-telemetry/opentelemetry-operator \
  --set "manager.collectorImage.repository=ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-k8s" \
  --set admissionWebhooks.certManager.enabled=false \
  --set admissionWebhooks.autoGenerateCert.enabled=true
```
### ❌ Desinstalar o Operator

Para remover totalmente:

```
helm uninstall opentelemetry-operator
```
Se estiver em namespace customizado:
```
helm uninstall opentelemetry-operator -n opentelemetry-operator-system
```