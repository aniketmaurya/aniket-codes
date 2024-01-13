# HasGeek Workshop

## Training Optimization with FSDP
Resource: [Blog](https://lightning.ai/blog/scaling-large-language-models-with-pytorch-lightning/)

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

**Download weights**
```shell
python scripts/download.py --repo_id TinyLlama/TinyLlama-1.1B-intermediate-step-955k-token-2T

python scripts/convert_hf_checkpoint.py \
    --checkpoint_dir checkpoints/TinyLlama/TinyLlama-1.1B-intermediate-step-955k-token-2T
```

**Test**
```
python chat/base.py --checkpoint_dir checkpoints/TinyLlama/TinyLlama-1.1B-intermediate-step-955k-token-2T
```

**Prepare Data**


**Finetune**
```
python finetune/lora.py \
    --checkpoint_dir checkpoints/TinyLlama/TinyLlama-1.1B-intermediate-step-955k-token-2T \
    --data_dir data/lima \
    --out_dir out/lima \
    --precision bf16-true \
    --quantize bnb.nf4
```

**Merge the LoRA weights**
```
python scripts/merge_lora.py \
  --checkpoint_dir "checkpoints/TinyLlama/TinyLlama-1.1B-intermediate-step-955k-token-2T" \
  --lora_path "out/lora/lima/lit_model_lora_finetuned.pth" \
  --out_dir "out/lora/lima/"
```

**Copy the tokenizer files**
```
cp checkpoints/TinyLlama/TinyLlama-1.1B-intermediate-step-955k-token-2T/*.json \
out/lora/lima/
```

**Evaluation**
```
python eval/lm_eval_harness.py \
  --checkpoint_dir "out/lora/lima/" \
  --precision "bf16-true" \
  --save_filepath "results.json"
```
