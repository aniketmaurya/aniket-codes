# deepspeed mii - https://github.com/microsoft/DeepSpeed-MII
import pandas as pd
import mii
from prompts import questions
import time

pipe = mii.pipeline("meta-llama/Llama-2-7b-hf")
output = pipe("This is me warming up the model", max_new_tokens=200)


results = []
for q in questions:
    t0 = time.perf_counter()
    output = pipe(q, max_new_tokens=200)
    t1 = time.perf_counter()
    results.append(
        {"time": t1 - t0, "tokens_generated": len(output[0].generated_length)}
    )

df = pd.DataFrame(results)
df.to_csv(f"deepspeed-benchmark-{tensor_parallel_size}GPUs.csv", index=False)
