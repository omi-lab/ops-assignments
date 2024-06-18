### Assignment 00

#### Dependencies

* [k6](https://k6.io/docs/get-started/installation/)

#### Browse the repo

Same as assignment 00.

#### Get a cluster up and running

Same as assignment 00.

#### Start the application

Same as assignment 00.

#### Assignment

You're asked to load test the api to test its horizontal scalability capabilities.

In order to do so, you'll write a `k6` script. You will notice that a `HorizontalPodAutoscaler` was added to the `deployments` directory. However, the autoscaling seems to be broken. Fix it.

Please describe below:
* How did you troubleshoot the issue ?
* What does it seems to be ? (there are actually 2)
* How did you solve it ?

Then, you can stop the api running the following command, and move on to the next assignment:

```bash
$ task stop
```
