#docker build -t kwick-kube . && docker run -it --rm --entrypoint bash -v $PWD:/code kwick-kube
docker build -t kwick-kube . && docker run -it --rm --entrypoint bash -v $PWD/kwick-kube:/kwick-kube kwick-kube

