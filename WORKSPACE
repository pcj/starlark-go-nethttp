load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# ===============================================================
# rules_proto
# ===============================================================

http_archive(
    name = "build_stack_rules_proto",
    urls = ["https://github.com/stackb/rules_proto/archive/4c2226458203a9653ae722245cc27e8b07c383f7.tar.gz"],
    sha256 = "0be90d609fcefae9cc5e404540b9b23176fb609c9d62f4f9f68528f66a6839bf",
    strip_prefix = "rules_proto-4c2226458203a9653ae722245cc27e8b07c383f7",
)

# ===============================================================
# rules_go
# ===============================================================

load("@build_stack_rules_proto//go:deps.bzl", "go_grpc_library")

go_grpc_library()

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

# ===============================================================
# bazel_gazelle
# ===============================================================

load("@build_stack_rules_proto//:deps.bzl", "bazel_gazelle")

bazel_gazelle()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# gazelle:repo bazel_gazelle

# ===============================================================
# starlark-go
# ===============================================================

go_repository(
    name = "net_starlark_go",
    commit = "5eb2bffb512a476426ca2447144f0b670efa1313",
    importpath = "go.starlark.net",
)

go_repository(
    name = "com_github_chzyer_readline",
    commit = "2972be24d48e78746da79ba8e24e8b488c9880de",
    importpath = "github.com/chzyer/readline",
)
