# HasGeek Workshop

## Training Optimization with FSDP

**Step 1:**
Run `fsdp/01.py` and observe the GPU memory usage.


**Step 2:**
Distribute the model to multiple GPU. Run `fsdp/02.py`

**Step 3:**
Use activation checkpointing and observe the GPU memory usage. Run `fsdp/03.py`


## Install Lit-GPT

```
git clone https://github.com/Lightning-AI/lit-gpt.git
cd lit-gpt
pip install -r requirements-all.txt
```

## Finetuning TinyLlama

```shell
python scripts/download.py --repo_id TinyLlama/TinyLlama-1.1B-intermediate-step-955k-token-2T

python scripts/convert_hf_checkpoint.py \
    --checkpoint_dir checkpoints/TinyLlama/TinyLlama-1.1B-intermediate-step-955k-token-2T
```

**Test**
```
python chat/base.py --checkpoint_dir checkpoints/TinyLlama/TinyLlama-1.1B-intermediate-step-955k-token-2T
```
