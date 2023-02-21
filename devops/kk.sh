#!/bin/bash

podman stop kkfile
podman rm kkfile
podman run -d --name kkfile -p 8012:8012 keking/kkfileview
