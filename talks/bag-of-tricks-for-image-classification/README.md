# Bag of Tricks for Image Classification

```
!pip install -q chitra==0.0.20
```

## import

```python
import numpy as np
import math
import tensorflow as tf
import tensorflow_datasets as tfds
import tensorflow_addons as tfa

from chitra.trainer import Trainer, create_cnn

import matplotlib.pyplot as plt

from tqdm import tqdm
```

## helper functions

```python
# data visualization

def visualize(data, nrow=2, ncol=3):
    fig, axs = plt.subplots(nrow, ncol)
    _data = []
    for img, label in data.take(nrow*ncol):
        _data.append((img.numpy(), str(label.numpy())))
    
    k = 0
    for i in range(nrow):
        for j in range(ncol):
            _data[k][0]
            axs[i][j].imshow(_data[k][0])
            axs[i][j].set_title(_data[k][1])
            k+=1
```

## define constants

```python
AUTOTUNE = tf.data.experimental.AUTOTUNE
BS = 16
```


## Load data with tfds

```python
(ds_train, ds_test), ds_info = tfds.load(
    'cifar10',
    split=['train', 'test'],
    shuffle_files=True,
    as_supervised=True,
    with_info=True,
)

visualize(ds_train)
```

## Preprocessing

Images augmentation is applied to generate more data points from the existing data. It helps in generalization of the model and produce a regularization effect.

```python
@tf.function
def rescale(image, label):
    image = tf.cast(image, tf.float32) / 127.5 - 1.0
    return image, label

@tf.function
def augment(image, label):
    image = tf.cast(image, tf.float32)
    image = tf.image.random_flip_left_right(image)
    image = image + tf.random.normal((32, 32, 3), mean=0.0, stddev=0.1)
    return image, label

train_dl = ds_train.map(augment, AUTOTUNE).map(rescale, AUTOTUNE).cache().batch(BS).prefetch(AUTOTUNE)
test_dl = ds_test.map(rescale).batch(BS).prefetch(AUTOTUNE)
```


## Build model

```python
model = create_cnn('resnet50', num_classes=10, weights=None)
```

## Learning rate finder

```python
def exp_annealing(step, max_steps, min_lr=1e-7, max_lr=10):
    return min_lr * (max_lr / min_lr) ** (step * 1. / max_steps)


def lr_finder(model, dataloader, min_lr=1e-7, max_lr=10., max_steps=100, beta=0.98):
    model = tf.keras.models.clone_model(model)

    avg_loss = 0.
    best_loss = 0.
    best_lr = 0.
    losses = []
    log_lrs = []
    batch_num = 0
    lr = min_lr
    criterion = tf.keras.losses.SparseCategoricalCrossentropy(from_logits=True)
    optimizer = tf.keras.optimizers.Adam(learning_rate=min_lr)

    for i, (images, gt_labels) in enumerate(tqdm(dataloader.take(max_steps), total=max_steps)):
        batch_num+=1
        with tf.GradientTape() as tape:
            logits = model(images)
            loss = criterion(gt_labels, logits)

        grads = tape.gradient(loss, model.trainable_variables)
        optimizer.apply_gradients(zip(grads, model.trainable_variables))

        avg_loss = beta * avg_loss + (1-beta) * loss.numpy()
        # print(f'avg_loss:{avg_loss}')
        smoothed_loss = avg_loss / (1 - beta**batch_num)
        # print(f'smoothed_loss:{smoothed_loss}')

        # Stop if the loss is exploding
        if batch_num > 1 and smoothed_loss > 4 * best_loss:
            return log_lrs, losses, {'best_loss':best_loss, 'best_lr': best_lr}
        # Record the best loss
        if smoothed_loss < best_loss or batch_num == 1:
            best_loss = smoothed_loss
            best_lr = lr

        losses.append(smoothed_loss)
        log_lrs.append(lr)

        lr = exp_annealing(i+1, max_steps, min_lr, max_lr)
        tf.keras.backend.set_value(optimizer.lr, lr)

    return (lrs, losses, {'best_loss':best_loss, 'best_lr': best_lr})

lrs,losses, lr_info = lr_finder(model, train_dl, max_steps=50000//BS)


def plot(lrs, losses, info=None):
    fig, ax = plt.subplots(1, 1)
    ax.set_ylabel('Loss')
    ax.set_xlabel('Learning Rate')
    ax.set_xscale('log')
    ax.xaxis.set_major_formatter(plt.FormatStrFormatter('%.0e'))
    if info:
        ax.plot(info['best_lr'], info['best_loss'], 'r+', markersize=25)
    ax.plot(lrs, losses)


plot(lrs[10:-5], losses[10:-5], lr_info)
plt.show()

lr_info['best_lr']/5
```


## Train the network with this learning rate

```python
init_lr = lr_info['best_lr']/5

model = create_cnn('resnet50', num_classes=10, weights=None)
optimizer = tf.keras.optimizers.Adam(init_lr)
criterion = tf.keras.losses.SparseCategoricalCrossentropy(from_logits=True)

model.compile(optimizer, criterion, metrics=['accuracy'])

callbacks = []
callbacks.append(tf.keras.callbacks.EarlyStopping(patience=7))

model.fit(train_dl, validation_data=test_dl, epochs=100, callbacks=callbacks)
```
