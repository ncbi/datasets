for name in datasets dataformat; do
    bazel build --build_event_text_file=$name.bep //src:$name
    bin_name=$(sed -n '/^completed {/,/^}/{s,.*uri: "file://\(.*\)",\1,p}' $name.bep)
    cp "$bin_name" "$PREFIX/bin/"
done
bazel clean
bazel shutdown
