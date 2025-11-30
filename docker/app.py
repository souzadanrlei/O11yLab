from fastapi import FastAPI
from pydantic import BaseModel
import logging

from opentelemetry import trace, metrics

# CORREÇÃO AQUI
tracer = trace.get_tracer("fastapi.tracer")

# Métricas
meter = metrics.get_meter("fastapi.meter")

item_counter = meter.create_counter(
    name="app.items.count",
    description="Contagem de operações de items"
)

app = FastAPI()

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


class Item(BaseModel):
    name: str


@app.get("/")
def root():
    logger.info("Endpoint root chamado")
    return {"message": "API Fast com Opentelemetry"}


@app.post("/items")
def create_item(item: Item):
    with tracer.start_as_current_span("post_create_item") as span:
        span.set_attribute("item.name", item.name)
        item_counter.add(1, attributes={"operation": "create"})
        logger.warning(f"Criado item: {item.name}")
        return {"status": "create", "item": item.name}


@app.get("/item/{item_id}")
def get_item(item_id: int):
    with tracer.start_as_current_span("get_item") as span:
        span.set_attribute("item.id", item_id)
        result = {"id": item_id, "name": f"item{item_id}"}
        item_counter.add(1, attributes={"operation": "read"})
        return result


@app.put("/item/{item_id}")
def update_item(item_id: int, item: Item):
    with tracer.start_as_current_span("update_item") as span:
        span.set_attribute("item.id", item_id)
        item_counter.add(1, attributes={"operation": "update"})
        return {"status": "update", "id": item_id, "name": item.name}


@app.delete("/item/{item_id}")
def delete_item(item_id: int):
    with tracer.start_as_current_span("delete_item") as span:
        span.set_attribute("item.id", item_id)
        item_counter.add(1, attributes={"operation": "delete"})
        return {"status": "delete", "id": item_id}
