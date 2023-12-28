# velero-plugin-suspend-cronjobs

Automatically suspend Kubernetes CronJobs prior to restoring them into a cluster.

## Building the plugin

To build the plugin, run

```bash
$ make
```

To build the Docker image, run

```bash
$ make container
```

This builds an image tagged as `github.com/tuusberg/velero-plugin-suspend-cronjobs:latest`. If you want to specify a different name or version/tag, run:

```bash
$ IMAGE=your-repo/your-name VERSION=your-version-tag make container 
```

To push the image to a Docker repository, run
```bash
$ make push
```

## Deploying the plugin

To deploy your plugin image to an Velero server:

### Using Velero CLI
1. Make sure your image is pushed to a registry that is accessible to your cluster's nodes.
2. Run `velero plugin add <registry/image:version>`. Example with a dockerhub image: `velero plugin add velero/velero-plugin-example`.

### Using Helm
1. Make sure your image is pushed to a registry that is accessible to your cluster's nodes.
2. Add the plugin to your Velero Helm chart's `values.yaml`:

    ```yaml
    velero:
      initContainers:
        - name: velero-plugin-suspend-cronjobs
              image: tuusberg/velero-plugin-suspend-cronjobs:latest
              imagePullPolicy: Always
              volumeMounts:
                - name: plugins
                  mountPath: /target
    ```
