load("@gazelle//:def.bzl", "gazelle")

# gazelle:resolve go datasets/openapi/v2 //openapi:golib.v2
gazelle(name = "gazelle")

genrule(
    name = "version",
    srcs = [".git"],
    outs = ["version.txt"],
    cmd_bash = "git --git-dir=$(location :.git) describe --always --dirty |tee $@",
)
