# Github workflows - architecture and implementation instructions.


## Project structure

ai-agent-management-platform
|   agent-manager-service
    |   Dockerfile
    |   ...
|   console
    |   Dockerfile
    |   ...
|   deployment
    | helm-charts
        | agent-management-platform
        | build-ci
        | observability-dataprepper
| quick-start
    | install.sh
    | ...
|   traces-observer-service
    | Dockerfile
    |   ...

## Goal

- When a new github release is created, 
    - build and push all the docker images to the ghcr (oci://ghcr.io/wso2) tagged with the release tag 
    - Update the helm charts with new docker image version, package and push to the ghcr
    - Update the helm chart version of the respective components in the install-helpers.sh 
    
    ```sh
    AMP_CHART_VERSION="${AMP_CHART_VERSION:-0.1.0}"
    BUILD_CI_CHART_VERSION="${BUILD_CI_CHART_VERSION:-0.1.0}"
    OBSERVABILITY_CHART_VERSION="${OBSERVABILITY_CHART_VERSION:-0.1.1}"
    ```
    - Make an archive of the quick-start directory and attach it to release artifacts with checksum.

## what should be done

Assuming that the project structure is there, create the required github workflow files to achieve the goal.

*IMPORTANT*

1. Do not make the scripts too verbose, make it simple
2. Log all essential steps and actions
3. This is a public repo so all the credentials should be taken from github secrets.
5. Follow github action, workflow best practices.


## UPDATE (NEW APPROACH)

Workflow inputs for release 

- branch (default main)
- Target-release
- Last release tag

*Release tag Naming convention*

- amp/v{target-release}. ex: amp/v0.0.1

### Images to be pushed

| Component                     | Image Name                            | Tag           |
|-------------------------------|---------------------------------------|---------------|
| console                       | amp-console                           | target-release|
| agent-management-service      | amp-api                               | target-release|
| traces-observer-service       | amp-trace-observer                    | target-release|
| python-instrumentation-provider                       | amp-python-instrumentation-provider   | target-release|
| quick-start                   | amp-quick-start                       | target-release|

### Helm charts

Update the target-release as the version, package and push

### Github process

1. Create a new branch from the input branch with the target release as the name (amp-{target-release})
2. Update the helm versions and image versions and commit
3. Create a tag with same approach as branch
4. Trigger the image build, push and helm chart publish
5. Generate the change-log from the target-release and the last release tag
6. Publish the release


*IMPORTANT*
The previous workflows are in backup directory. Refer to them if required.
Should be a single, user triggered workflow with clear steps and logs 
NO excessive logging


