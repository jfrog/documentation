# Scan Gitlab Pull Request

After you create a new merge request, the maintainer of the Git repository can trigger Frogbot to scan the merge request from the merge request UI.

_**NOTE:**_ The scan output will include only new vulnerabilities added by the merge request. Vulnerabilities that aren't new, and existed in the code before the merge request was created, will not be included in the report. In order to include all of the vulnerabilities in the report, including older ones that weren't added by this merge request, use the includeAllVulnerabilities parameter in the frogbot-config.yml file.

The Frogbot GitLab flow is as follows:

1. The developer opens a merge request.
2. The maintainer of the repository reviews the merge request and approves the scan by triggering the manual _frogbot-scan_ job.
3. Frogbot is then triggered by the job, it scans the merge request and adds a comment with the scan results.
4. Frogbot can be triggered again following new commits, by triggering the _frogbot-scan_ job again.

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/gitlab-run-button.png" alt="" data-size="original">
