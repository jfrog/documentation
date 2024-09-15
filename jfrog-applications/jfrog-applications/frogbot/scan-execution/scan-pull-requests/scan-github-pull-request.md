# Scan GitHub Pull Request

After you create a new pull request, the maintainer of the Git repository can trigger Frogbot to scan the pull request from the pull request UI.

_**NOTE:**_ The scan output will include only new vulnerabilities added by the pull request. Vulnerabilities that aren't new, and existed in the code before the pull request was created, will not be included in the report. In order to include all the vulnerabilities in the report, including older ones that weren't added by this PR, use the includeAllVulnerabilities parameter in the frogbot-config.yml file.

The Frogbot GitHub scan workflow is:

1. The developer opens a pull request.
2. The Frogbot workflow automatically gets triggered and a [GitHub environment](https://docs.github.com/en/actions/deployment/targeting-different-environments/using-environments-for-deployment#creating-an-environment) named `frogbot` becomes pending for the maintainer's approval.

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/github-pending-deployment.png" alt="" data-size="original">

3. The maintainer of the repository reviews the pull request and approves the scan:

<img src="https://raw.githubusercontent.com/jfrog/frogbot/master/images/github-deployment.gif" alt="" data-size="original">

4. Frogbot can be triggered again following new commits, by repeating steps 2 and 3.
