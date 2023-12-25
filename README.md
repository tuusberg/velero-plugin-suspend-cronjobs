# velero-plugin-suspend-cronjobs

Automatically suspend Kubernetes CronJobs prior to restoring them into a cluster

![Build Status][1]

## Building the plugin

To build the plugin, run

```bash
$ make
```

To build the image, run

```bash
$ make container
```

This builds an image tagged as `velero/velero-plugin-example:main`. If you want to specify a different name or version/tag, run:

```bash
$ IMAGE=your-repo/your-name VERSION=your-version-tag make container 
```

## Deploying the plugins

To deploy your plugin image to an Velero server:

1. Make sure your image is pushed to a registry that is accessible to your cluster's nodes.
2. Run `velero plugin add <registry/image:version>`. Example with a dockerhub image: `velero plugin add velero/velero-plugin-example`.


## Creating your own plugin project

1. Create a new directory in your `$GOPATH`, e.g. `$GOPATH/src/github.com/someuser/velero-plugins`
2. Copy everything from this project into your new project

```bash
$ cp -a $GOPATH/src/github.com/vmware-tanzu/velero-plugin-example/* $GOPATH/src/github.com/someuser/velero-plugins/.
```

3. Remove the git history

```bash
$ cd $GOPATH/src/github.com/someuser/velero-plugins
$ rm -rf .git
```

4. Adjust the existing plugin directories and source code as needed.

The `Makefile` is configured to automatically build all directories starting with the prefix `velero-`.
You most likely won't need to edit this file, as long as you follow this convention.

If you need to pull in additional dependencies to your vendor directory, just run

```bash
$ make modules
```

[1]: https://github.com/vmware-tanzu/velero-plugin-example/workflows/Continuous%20Integration/badge.svg

