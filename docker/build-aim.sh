#!/bin/bash

set -e
set -x

mkdir -p build/aim/dist

pushd build/aim

# Build base image and dependencies
git clone -b k8s_universe https://github.com/noironetworks/aci-integration-module \
    --depth 1

cp ../../docker/Dockerfile-aim-build aci-integration-module
docker build -t noiro/aci-integration-module-build \
       -f aci-integration-module/Dockerfile-aim-build \
       aci-integration-module
cat <<EOF | docker run -i noiro/aci-integration-module-build tar -X - -z -c -C /usr \
       lib/python2.7 bin/aim-aid bin/aimctl bin/aimdebug \
    > dist/aim-dist.tar.gz
lib/python2.7/site-packages/flask*
lib/python2.7/site-packages/Flask*
lib/python2.7/site-packages/pip*
lib/python2.7/site-packages/setuptools*
lib/python2.7/site-packages/git*
lib/python2.7/site-packages/pymysql*
EOF

# Build the minimal AIM container
cp ../../docker/Dockerfile-aim dist
docker build -t noiro/aci-integration-module -f dist/Dockerfile-aim dist

popd
