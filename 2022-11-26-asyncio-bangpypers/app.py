from fastapi import FastAPI
import time
import asyncio
import uvicorn

app = FastAPI()

@app.get("/test1")
async def hello_world():
    await asyncio.sleep(4)
    return "ok"

uvicorn.run(app)
