import lightning as L
import uvicorn
from application.server.main import app as api


class FastAPIWork(L.LightningWork):
    """
    1. a building block for long-running jobs
    2. Runs on an independent process/machine
    3. Use custom docker image
    """
    def __init__(self, parallel: bool = False, **kwargs):
        super().__init__(parallel=parallel, **kwargs)
    
    def run(self):
        uvicorn.run(api, host=self.host, port=self.port)


class RootFlow(L.LightningFlow):
    def __init__(self):
        super().__init__()
        self.fastapi_work = FastAPIWork()
    
    def run(self, *args, **kwargs) -> None:
        self.fastapi_work.run()
    
    def configure_layout(self):
        return {"name":"Swagger", "content":self.fastapi_work.url}

if __name__=="__main__":
    app = L.LightningApp(RootFlow())
