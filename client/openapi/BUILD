# gazelle:ignore

load("@aspect_bazel_lib//lib:copy_to_directory.bzl", "copy_to_directory")
load("@rules_go//go:def.bzl", "go_library")
load("@openapi_tools_generator_bazel//:defs.bzl", "openapi_generator")

package(default_visibility = ["//visibility:public"])

## Golang
filegroup(
    name = "golang_template",
    srcs = glob([
        "templates/go/*.mustache",
    ]),
)

filegroup(
    name = "openapi3_v2_yaml",
    srcs = ["openapi3_v2.yaml"],
)

[
    openapi_generator(
        name = "openapi.genlib.{ver}.go".format(ver = version),
        config = "configs/openapigen.go.config".format(ver = version),
        generator = "go",
        spec = ":openapi3_{ver}_yaml".format(ver = version),
        template_dir = "templates/go",
    )
    for version in [
        "v2",
    ]
]

## Create go Library (two step process)
# See issue https://github.com/OpenAPITools/openapi-generator-bazel/issues/22
# for use of copy_to_directory
# TODO: Clean up and put into DT/toolshed
[
    copy_to_directory(
        name = "gofiles.{ver}.go".format(ver = version),
        srcs = [":openapi.genlib.{ver}.go".format(ver = version)],
        exclude_srcs_patterns = [
            "**/main.go",
            "**/README.md",
            "**/git_push.sh",
            "**/*_test.go",
            "**/.*ignore",
            "**/.travis.yml",
            "**/docs/*",
            "**/api/*",
            "**/.openapi-generator/*",
        ],
        include_srcs_patterns = ["**/*.go"],
    )
    for version in [
        "v2",
    ]
]

[
    go_library(
        name = "golib.{ver}".format(ver = version),
        srcs = [":gofiles.{ver}.go".format(ver = version)],
        importpath = "datasets/openapi/{ver}".format(ver = version),
        visibility = [
            "//visibility:public",
        ],
    )
    for version in [
        "v2",
    ]
]
