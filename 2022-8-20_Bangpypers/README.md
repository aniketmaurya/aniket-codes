# Bangpypers meetup

## Steps to build ML API

### Step 1: Create FastAPI Work

### Step 2: Create Rootflow to run the FastAPI Work
*    Configure layout for swagger UI


## Code Demos used in slides


### Train and Deploy Flow

```python
class Flow(L.LightningFlow):
    def __init__(self):
        super().__init__()
        self.train = ModelTraining(cloud_compute=L.CloudCompute("gpu", disk_size=20))
        self.deploy = ModelDeploy(cloud_compute=L.CloudCompute("cpu"))

    def run(self):
        self.train.run()
        self.deploy.run(checkpoint_dir=self.train.checkpoint_dir)

```


### Run training

```python
import lightning as L

class ModelTraining(L.LightningWork):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        # Convert every filesystem path you want to share with other LightningWorks to by adding lit:// in front of it.
        self.checkpoint_dir = "lit://outputs/checkpoints"

    def run(self):
        download_data("https://MachineLearning-Dataset.zip", "./dataset")
        trainer.fit(model, datamodule)
        torch.save(model, os.path.join(self.checkpoint_dir, "checkpoint.ckpt"))
```

### Model deploy

```python
class ModelDeploy(LightningWork):
    def __init__(self, *args, **kwargs):
        super().__init__()

    def run(self, checkpoint_dir):
        ckpts = os.listdir(checkpoint_dir)
        checkpoint = torch.load(ckpts[0])
        deploy_model(checkpoint)
```

