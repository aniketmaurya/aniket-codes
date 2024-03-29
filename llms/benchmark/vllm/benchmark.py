# vLLM - https://github.com/vllm-project/vllm
from vllm import LLM, SamplingParams
import pandas as pd
import time
from rich.progress import track
import os
from prompts import questions

tensor_parallel_size = int(os.environ.get("DEVICES", "1"))
print("tensor_parallel_size: ", tensor_parallel_size)

sampling_params = SamplingParams(max_tokens=200)

llm = LLM("meta-llama/Llama-2-7b-hf", tensor_parallel_size=tensor_parallel_size)
print(llm.generate("This is me warming up the model", sampling_params=sampling_params))

results = []
for q in track(questions):
    t0 = time.perf_counter()
    output = llm.generate(q, sampling_params=sampling_params)[0]
    t1 = time.perf_counter()
    results.append(
        {"time": t1 - t0, "tokens_generated": len(output.outputs[0].token_ids)}
    )

df = pd.DataFrame(results)
df.to_csv(f"vllm-benchmark-{tensor_parallel_size}GPUs.csv", index=False)
