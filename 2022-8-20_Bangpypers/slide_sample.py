import lightning as L
import torch


class ModelTraining(L.LightningWork):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        # Convert every filesystem path you want to share with other LightningWorks to by adding lit:// in front of it.
        self.checkpoint_dir = "lit://outputs/checkpoints"

    def run(self):
        trainer.fit(model, datamodule)
        torch.save(model, os.path.join(self.checkpoint_dir, "checkpoint.ckpt"))


class ModelDeploy(L.LightningWork):
    def __init__(self, *args, **kwargs):
        super().__init__()

    def run(self, checkpoint_dir):
        ckpts = os.listdir(checkpoint_dir)
        checkpoint = torch.load(ckpts[0])
        deploy_model(checkpoint)


class Flow(L.LightningFlow):
    def __init__(self):
        super().__init__()
        self.train = ModelTraining(cloud_compute=L.CloudCompute("gpu", disk_size=20))
        self.deploy = ModelDeploy(cloud_compute=L.CloudCompute("cpu"))

    def run(self):
        self.train.run()
        self.deploy.run(checkpoint_dir=self.train.checkpoint_dir)

    def configure_layout(self):
        return {"name": "Machine Learning API", "content": self.deploy.url}
