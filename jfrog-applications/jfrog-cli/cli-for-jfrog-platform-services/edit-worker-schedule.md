# Edit Worker Schedule

### Overview



Edit the manifest of a SCHEDULED\_EVENT worker to udpate the schedule criteria.

The worker should be deploy afterward with `jf worker deploy` for the change to be applied to the server.

|                      |                                                                                                                         |
| -------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| Command name         | worker edit-schedule                                                                                                    |
| Abbreviation         | worker es                                                                                                               |
| **Command options:** |                                                                                                                         |
| `--cron`             | \[Mandatory] A standard cron expression with minutes resolution. Seconds resolution is not supported by Worker service. |
| `--timezone`         | \[Default: UTC] The timezone to use for scheduling.                                                                     |

### Example



Edit a worker manifest so that it is executed every minute.

```
jf worker edit-schedule --cron "* * * * *"
```
