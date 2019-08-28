# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

__config__ = pulumi.Config('docker')

ca_material = __config__.get('caMaterial') or utilities.get_env('DOCKER_CA_MATERIAL')
"""
PEM-encoded content of Docker host CA certificate
"""

cert_material = __config__.get('certMaterial') or utilities.get_env('DOCKER_CERT_MATERIAL')
"""
PEM-encoded content of Docker client certificate
"""

cert_path = __config__.get('certPath') or utilities.get_env('DOCKER_CERT_PATH')
"""
Path to directory with Docker TLS config
"""

host = __config__.get('host') or (utilities.get_env('DOCKER_HOST') or 'unix:///var/run/docker.sock')
"""
The Docker daemon address
"""

key_material = __config__.get('keyMaterial') or utilities.get_env('DOCKER_KEY_MATERIAL')
"""
PEM-encoded content of Docker client private key
"""

registry_auth = __config__.get('registryAuth')

