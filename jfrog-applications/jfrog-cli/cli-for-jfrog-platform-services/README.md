# CLI for JFrog Platform Services

This page describes how to use JFrog CLI with JFrog Platform Services.

Read more about JFrog CLI [here](https://jfrog-external.fluidtopics.net/r/help/JFrog-CLI/JFrog-CLI).

**Managing JFrog Workers**

Workers is a JFrog Platform service that allows you to extend and control your execution flows. It provides a serverless execution environment. You can create workers to enhance the platform's functionality. Workers are triggered automatically by events within the JFrog Platform, giving you the flexibility to address specific use cases. For on-demand tasks, configure Http-triggered workers.

You can read more about JFrog Workers [here](https://jfrog.com/help/r/jfrog-platform-administration-documentation/workers).


With the JFrog CLI for JFrog Worker it is possible to:

- [Initialize a JFrog worker](./platform-workers-init.md)
- [Test-Run a JFrog worker](./platform-workers-dry-run.md)
- [Deploy a JFrog worker](./platform-workers-deploy.md)
- [Add secrets to a JFrog worker](./platform-workers-add-secret.md)
- [Undeploy a JFrog worker](./platform-workers-undeploy.md)
- [Trigger an HTTP-Triggered worker](./platform-workers-execute.md)
- [Retrieve the list of available events](./platform-workers-list-event.md)
- [Retrieve the list of existing JFrog workers](./platform-workers-list.md)
- [Retrieve a JFrog worker's execution history](./platform-workers-execution-history.md)
- [Edit an SCHEDULE_EVENT worker schedule](./platform-workers-edit-schedule.md)