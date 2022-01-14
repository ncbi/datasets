# NCBI Datasets command-line clients

Build all command-line clients
```bash
bazel build src/...
```

Build only datasets command-line client
```bash
bazel build src:datasets
```
The executable should be at bazel-bin/src/datasets_/datasets

Build only dataformat command-line client
```bash
bazel build src:dataformat
```
The executable should be at bazel-bin/src/dataformat_/dataformat
