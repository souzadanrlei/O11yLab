FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD [ "opentelemetry-instrument", "--service_name", "fastapi-service", "uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8080"]
