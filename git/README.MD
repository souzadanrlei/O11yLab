# 📌 Como criar e subir novas branches no projeto

Este guia descreve o fluxo recomendado para criar uma nova branch, adicionar mudanças e enviar o conteúdo para o repositório remoto no GitHub.

### 🚀 1. Atualize o repositório local

Antes de criar uma branch nova, sempre garanta que você está com o código atualizado:
```
git pull
```
### 🌱 2. Crie uma nova branch

Crie e já mude automaticamente para ela:
```
git checkout -b development
```

Substitua development pelo nome da branch desejada.

### 📄 3. Verifique o status dos arquivos

Veja quais arquivos foram alterados, adicionados ou removidos:
```
git status
```
### ➕ 4. Adicione os arquivos ao stage

Adicione todos os arquivos modificados:
```
git add .
```

Verifique novamente:
```
git status
```
### 📝 5. Faça o commit

Com uma mensagem clara indicando a mudança:
```
git commit -m "feat: informações sobre o commit"
```

Exemplos de mensagens:
feat: adiciona gateway OpenTelemetry
fix: corrige configuração do collector
docs: adiciona instruções de instalação

### 📤 6. Envie a branch para o repositório remoto

Envie sua nova branch para o GitHub:
```
git push -u origin development
```
### ✔️ Pronto!

Sua branch agora está disponível no GitHub e pode ser utilizada para Pull Requests, revisão de código ou integração.