# HasGeek Workshop

## Training Optimization with FSDP

```

```

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
