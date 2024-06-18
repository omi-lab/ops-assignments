### Assignment 00

#### Dependencies

For that assignment, you'll need at least the following dependencies:
* [minikube](https://minikube.sigs.k8s.io/docs/), [kind](https://kind.sigs.k8s.io/), or whatever other tool to get a local k8s cluster running
* [taskfile](https://taskfile.dev/)
* [kubectl](https://kubernetes.io/docs/tasks/tools/)
* [tilt](https://docs.tilt.dev/api.html)

#### Browse the repo

The `00` directory consists of a few different subdirectories:
* `deployments` contains the manifests for deploying our app
* `fixtures` stores temporary placeholder files
* `Dockerfile` to build the API as a `Docker` image
* `Taskfile.yml` for local task presets
* `Tiltfile` to orchestrate the local deployment of our app

#### Get a cluster up and running

If you decide to go with the already available `Taskfiles` and you got `kind` installed, you should be able to run at the root of this repo:
```bash
$ task k8s:create
```

Once done, your current context should be set to `kind-kind`.

Otherwise, just start a local `Kubernetes` cluster with the tool of your choice.

#### Start the application

You can now start the local application running:
```bash
$ cd 00
$ task start
```

`Tilt` will be starting, and you can press `<space>` to open the UI and track its deployment.

If you prefer to deploy it without `task` just run `tilt up`.

Now that our app is running, you should be able to see it with `kubectl` running `kubectl get pods -n api`. You should notice nothing too fancy.

When running `tilt up`, a tunnel was opened on port 3000 between your local host and the running pod. To make sure of it, you can run:

```bash
$ lsof -nP -iTCP:3000 -sTCP:LISTEN
COMMAND   PID            USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
tilt    20700 elliotmaincourt   21u  IPv4 0x60cdbe7994fa5ea7      0t0  TCP 127.0.0.1:3000 (LISTEN)
tilt    20700 elliotmaincourt   23u  IPv6 0x60cdbe6b2243e01f      0t0  TCP [::1]:3000 (LISTEN)
```

Or, more simply, you can also run:

```bash
$ curl localhost:3000/v1/healthz -v
*   Trying 127.0.0.1:3000...
* Connected to localhost (127.0.0.1) port 3000 (#0)
> GET /v1/healthz HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/8.1.2
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Mon, 10 Jun 2024 15:17:12 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
```

#### Assignment

Now that we're all setup, run the following command and see what happens:

```bash
$ curl -i -X POST -H "Content-Type: multipart/form-data" -F "file=@fixtures/file_to_upload" http://localhost:3000/v1/files
```

Your first task will be to diagnose what is happening when uploading that file, and fix it. Note that deployment manifests should not be changed.

Please describe below:
* How did you troubleshoot the issue ?
* What does it seems to be ?
* How could you solve it ?

Then eventually implement the fix in *2 different ways*

Then, you can stop the api running the following command, and move on to the next assignment:

```bash
$ task stop
```
