# namespace plugin
load("ext://namespace", "namespace_inject", "namespace_create")
# trigger_mode(TRIGGER_MODE_MANUAL)

# Set up Minikube Docker environment
local('eval $(minikube docker-env)')

namespace = "is459"
registry = "local"
# registry = "registry.gitlab.com/fyp6033103"

modules = [
    {
        "image_repo": "go-consumer",
        "chart_repo": "go-consumer-charts",
        "values": "values.yaml",
        "branch": "main",
    }
    # {
    #     "image_repo": "media-server",
    #     "chart_repo": "media-server-charts",
    #     "values": "dev.values.yaml",
    #     "branch": "jk-dev",
    # },
    #   {
    #       "image_repo": "platform-api",
    #       "chart_repo": "platform-api-charts",
    #       "values":  "dev.values.yaml" ,
    #   },
]

# create the namespace
namespace_create(namespace)

# # deploy secrets first
# k8s_yaml(namespace_inject(read_file("./secrets.yml"), namespace))
# k8s_yaml(namespace_inject(read_file("./secrets.yml"), "challenge"))

# Deploy Kafka, Kafka UI, Kafka Connect, MongoDB
# k8s_yaml(namespace_inject(helm("./k8s/kafka-charts/", name="kafka"), namespace ), allow_duplicates=False)
# k8s_yaml(namespace_inject(helm("./k8s/kafka-connect-charts/", name="kafka-connect"), namespace ), allow_duplicates=False)
# k8s_yaml(namespace_inject(helm("./k8s/kafka-ui-charts/", name="kafka-ui"), namespace ), allow_duplicates=False)
k8s_yaml(namespace_inject(helm("./k8s/mongodb-charts/", name="mongodb-db"), namespace ), allow_duplicates=False)


# for each module
for m in modules:
    # image_tag = registry + '/' + m["image_repo"] + '/' + m["branch"]
    image_tag = m["image_repo"] + ":latest"
    context = "./" + m["image_repo"]
    dockerfile = "./" + m["image_repo"] + "/docker/Dockerfile"
    chart = "k8s/" + m["chart_repo"] + "/"
    values = chart + m["values"]

    # build it
    docker_build(
        ref=image_tag,
        context=context,
        dockerfile=dockerfile,
        live_update=[sync("./" + m["image_repo"], "/app")],
        extra_tag=["latest"],
    )

    # and deploy it with helm
    k8s_yaml(
        namespace_inject(helm(chart, name=m["image_repo"], values=values), namespace),
        allow_duplicates=False,
    )

# # port forward application-server
# k8s_resource(
#     "application-server",
#     port_forwards=[
#         "8080:8080"  # Port forwarding from local port 8080 to container port 8080
#     ],
# )
