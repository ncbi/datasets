# Building Datasets Client
## Setup
The build system depends on having `bazel` available.  You should use [bazelisk](https://github.com/bazelbuild/bazelisk).  

Second, you must have the file `workspace_status.sh` available in your PATH.

## Building
From this directory, run:

```
bazel build apps/public/Datasets/...
```



The executable will be available at `bazel-bin/apps/public/Datasets/v2/cmd/datasets/datasets_/datasets`.

## Alternative installations
`datasets` is also available through [conda](https://anaconda.org/conda-forge/ncbi-datasets-cli).
