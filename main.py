from fastapi import FastAPI
from app.controller.user_controller import router as user_router
import uvicorn

IP = "127.0.0.1"
PORT = 8080

app: FastAPI = FastAPI(
    title="FastAPI部分的Swagger文档集成",
    description="这是demo项目的FastAPI部分的Swagger文档集成",
    version="1.0.0"
)

app.include_router(user_router)

if __name__ == "__main__":
    uvicorn.run("main:app", host=IP, port=PORT, reload=True)
