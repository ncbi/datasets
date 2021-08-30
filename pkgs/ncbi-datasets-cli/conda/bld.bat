cd pkgs/ncbi-datasets-cli
FOR %%N in (datasets dataformat) DO (
    bazel build //src:%%N
    COPY "bazel-bin/src/%%N_/%%N" "%%PREFIX%%/bin/"
)
bazel clean
bazel shutdown
