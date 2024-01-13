
import torch
import torch.nn as nn
import torch.nn.functional as F
from pytorch_lightning.demos import Transformer, WikiText2
from torch.utils.data import DataLoader
from pytorch_lightning.callbacks import RichProgressBar
from pytorch_lightning.callbacks import DeviceStatsMonitor

import pytorch_lightning as pl

# intialize model, optimizer and defines training step
class LanguageModel(pl.LightningModule):
    def __init__(self, vocab_size):
        super().__init__()
        self.model = Transformer(
            vocab_size=vocab_size,
            nlayers=64,
            nhid=4096,
            ninp=1024,
            nhead=64,
        )

    def training_step(self, batch):
        input, target = batch
        output = self.model(input, target)
        loss = F.nll_loss(output, target.view(-1))
        self.log("train_loss", loss, prog_bar=True)
        return loss

    def configure_optimizers(self):
        return torch.optim.Adam(self.parameters(), lr=0.1)

pl.seed_everything(42)

# Data
dataset = WikiText2()
train_dataloader = DataLoader(dataset, batch_size=64)

# Model
model = LanguageModel(vocab_size=dataset.vocab_size)

# Trainer
trainer = pl.Trainer(accelerator="cuda", devices=4,)

trainer.fit(model, train_dataloader)
trainer.print(f"Memory used: {torch.cuda.max_memory_allocated() / 1e9:.02f} GB")
