"""
Dependencies of the GDH proejct
"""

load("@openapi_tools_generator_bazel//:defs.bzl", "openapi_tools_generator_bazel_repositories")

def _openapi_deps():
    openapi_tools_generator_bazel_repositories(
        openapi_generator_cli_version = "7.2.0",
        sha256 = "",
        prefix = "openapi_tools_generator_bazel",
    )

def deps():
    _openapi_deps()
