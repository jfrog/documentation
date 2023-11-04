# Installing Frogbot on JFrog Pipelines

**Important Notice**: For Scanning Pull Requests, it is advisable to refrain from setting up Frogbot using JFrog Pipelines for open source projects. For further details, please refer to the [👮 Security Note for Pull Requests Scanning](./scan-pull-requests.md#security-note-for-pull-requests-scanning).

* Make sure you have the connection details of your JFrog Platform.
* Inside JFrog Pipelines, save the JFrog connection details as a [JFrog Platform Access Token Integration](https://www.jfrog.com/confluence/display/JFROG/JFrog+Platform+Access+Token+Integration) named **jfrogPlatform**.
* Inside JFrog Pipelines, save your Git access token in a [Git Server Integration](https://jfrog.com/help/r/jfrog-pipelines-documentation/pipelines-integrations) named **gitIntegration**.
* Create a **pipelines.yml** file using one of the [available templates](https://github.com/jfrog/frogbot/tree/master/docs/templates/jfrog-pipelines) and push the file into one of your Git repositories, under a directory named `.jfrog-pipelines`.
* In the **pipelines.yml**, make sure to set values for all the mandatory variables.
* In the **pipelines.yml**, if you're using a Windows agent, modify the code inside the onExecute sections as described in the template comments.
* Ensure that the JFrog Pipelines agent has the necessary package managers installed. For example, if the project utilizes npm, it is crucial to have the npm client installed on the agent.
