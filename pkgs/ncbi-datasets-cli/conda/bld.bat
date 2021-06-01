FOR %%name in datasets dataformat DO (
    bazel build //src:%name%
    COPY "bazel-bin/src/%name%_/%name%" "%PREFIX%/bin/"
)
bazel clean
bazel shutdown
