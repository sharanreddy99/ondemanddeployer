<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">Adhoc Project Deployer</h1>
</p>
<p align="center">
    <em>An API which launches personal dockerized projects on github on-demand in a single server</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/commit-activity/m/sharanreddy99/ondemanddeployer" alt="last-commit">
	<img src="https://img.shields.io/github/created-at/sharanreddy99/ondemanddeployer" alt="created_at">
   <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/sharanreddy99/ondemanddeployer">
   <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/sharanreddy99/ondemanddeployer">
   <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/sharanreddy99/ondemanddeployer">

</p>

<p align="center">
	<!-- default option, no dependency badges. -->
</p>

<br><!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary><br>

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ Getting Started](#-getting-started)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
</details>
<hr>

##  Overview

The ondemanddeployer project automates deployments by orchestrating AWS resources and GitHub interactions through key components like AWS SDK and Beego framework. It streamlines deployment processes by executing bash scripts, handling SNS communications, and managing GitHub repositories. Leveraging configuration settings and constants for consistency, ondemanddeployer facilitates efficient deployment workflows, ensuring swift deployment setups and streamlined project management within its architecture.

---

##  Features

|    |   Feature         | Description |
|----|-------------------|---------------------------------------------------------------|
| ⚙️  | **Architecture**  | The project uses Beego web framework for web server functionalities, interacts with AWS resources for SNS subscriptions, and automates deployments through bash scripts. It maintains essential configuration settings and constants for deployment processes. |
| 🔩 | **Code Quality**  | The codebase follows a structured approach, uses modular design patterns, and adheres to consistent coding style practices. It ensures readability and maintainability with clear separation of concerns. |
| 📄 | **Documentation** | The project includes detailed documentation in code files such as main.go, github.go, conf/app.conf, and utils/utils.go. It covers configuration settings, controller operations, and utility functions. Additional documentation can enhance understanding and onboarding. |
| 🔌 | **Integrations**  | Key integrations include AWS SDK for SNS interactions, GitHub API for user data retrieval, and Beego framework for web server functionalities. External dependencies like crypto, oauth2, and mapstructure are utilized for various operations. |
| 🧩 | **Modularity**    | The codebase exhibits modularity through separate packages for AWS, GitHub, controllers, and utilities. This modular design enhances reusability, scalability, and ease of maintenance. Each module focuses on specific functionality, promoting code clarity and organization. |
| 🧪 | **Testing**       | The testing framework and tools used are not explicitly mentioned in the repository contents. Implementing unit tests, integration tests, or end-to-end tests can ensure code reliability and robustness. |
| ⚡️  | **Performance**   | The project focuses on efficiency by using caching mechanisms for GitHub data, concurrent operations for logging, and synchronized locks for file handling. Performance improvements can be made by optimizing resource usage and enhancing execution speed. |
| 🛡️ | **Security**      | Security measures like HTTPS configuration, access control settings in app.conf, and AWS authentication mechanisms are in place for data protection. Continuous monitoring, encryption standards, and vulnerability assessments can further strengthen security aspects. |
| 📦 | **Dependencies**  | Key external libraries and dependencies include Beego, AWS SDK, crypto, oauth2, mapstructure, and githubv4. These dependencies facilitate seamless integration, enhance functionality, and ensure compatibility with external services. |
| 🚀 | **Scalability**   | The project's architecture and modularity support scalability by enabling the addition of new features, handling increased traffic, and accommodating growing user demands. Implementing load balancing, database sharding, and microservices architecture can enhance scalability further. |

---

##  Repository Structure

```sh
└── ondemanddeployer/
    ├── components
    │   ├── aws
    │   ├── bashscript
    │   └── github
    ├── conf
    │   └── app.conf
    ├── constants
    │   └── constants.go
    ├── controllers
    │   ├── aws.go
    │   ├── base.go
    │   └── github.go
    ├── data
    │   └── github
    ├── github.go
    ├── go.mod
    ├── go.sum
    ├── lastupdate.tmp
    ├── main.go
    ├── routers
    │   └── router.go
    ├── scripts
    │   └── scripts.sh
    ├── templates.txt
    └── utils
        ├── http.go
        ├── logger.go
        └── utils.go
```

---

##  Modules

<details closed><summary>.</summary>

| File                                                                                             | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ---                                                                                              | ---                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| [go.sum](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/go.sum)               | The code file in this repository plays a crucial role in automating on-demand deployments by orchestrating AWS resources, executing bash scripts, and interacting with GitHub repositories. It leverages configuration settings from the app.conf file and maintains essential constants for consistency. This code file is integral to the on-demand deployers architecture, facilitating efficient and streamlined deployment processes. |
| [main.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/main.go)             | Initiates AWS SNS subscription with specified endpoint using the AWS SDK. Configures Beego web framework for CORS handling. Schedules subscription call after a delay or serves Swagger UI in development mode.                                                                                                                                                                                                                            |
| [go.mod](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/go.mod)               | Defines project dependencies and versions for `ondemanddeployer` app. Ensures compatibility with external libraries like Beego and AWS SDK. Facilitates smooth integration and stable performance across the applications modules.                                                                                                                                                                                                         |
| [github.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/github.go)         | Handles CRUD operations for User entities in GitHubController.-Fetches user data, lists all users, updates and deletes users, logs users in/out.-Integrates with Beego framework for web operations from parent repositorys architecture.                                                                                                                                                                                                  |
| [templates.txt](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/templates.txt) | Defines deployment templates for project ceta, specifying actions and parameters. Facilitates automated project setup, building, and management within the ondemanddeployer architecture.                                                                                                                                                                                                                                                  |

</details>

<details closed><summary>routers</summary>

| File                                                                                             | Summary                                                                                                                                             |
| ---                                                                                              | ---                                                                                                                                                 |
| [router.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/routers/router.go) | Defines API routes for AWS, GitHub, and Base controllers in Beego, facilitating interaction with distinct services in the ondemanddeployer project. |

</details>

<details closed><summary>constants</summary>

| File                                                                                                     | Summary                                                                                                                                                                                                                                                                                                                          |
| ---                                                                                                      | ---                                                                                                                                                                                                                                                                                                                              |
| [constants.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/constants/constants.go) | Defines essential configuration constants for HTTP, AWS, and GitHub integration, including ports, URLs, allowed repositories, data paths, cache expiry time, AWS region, SNS details, and instance metadata endpoint. Impactful for orchestrating on-demand deployment functionality within the parent repositorys architecture. |

</details>

<details closed><summary>conf</summary>

| File                                                                                        | Summary                                                                                                                                                                                                                                          |
| ---                                                                                         | ---                                                                                                                                                                                                                                              |
| [app.conf](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/conf/app.conf) | Implements configuration settings for the on-demand deployment application. Manages app details, HTTP settings, and file paths for secure communication. Key settings include app name, port, run mode, HTTPS configuration, and SQL connection. |

</details>

<details closed><summary>controllers</summary>

| File                                                                                                 | Summary                                                                                                                                                                                                                 |
| ---                                                                                                  | ---                                                                                                                                                                                                                     |
| [aws.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/controllers/aws.go)       | PublishSNS sends messages to SNS Topic, SubscribeSNS receives published messages. Utilizes AWS and utility functions. Key features serve as endpoints to manage SNS communications within the repository architecture.  |
| [github.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/controllers/github.go) | Retrieves repositories list, languages, and metadata from user profile. Key features include fetching data via GitHub API and serving JSON responses. Contributes to repositorys architecture for managing deployments. |
| [base.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/controllers/base.go)     | Implements health check endpoint for validating app status in the projects Beego-based web server. A part of the controllers package, it serves a JSON response indicating server health.                               |

</details>

<details closed><summary>utils</summary>

| File                                                                                           | Summary                                                                                                                                                                                                                         |
| ---                                                                                            | ---                                                                                                                                                                                                                             |
| [logger.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/utils/logger.go) | Enables logging to a file with concurrency control. Provides a reusable mechanism for writing log messages.                                                                                                                     |
| [http.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/utils/http.go)     | Executes HTTP requests to external APIs, reading and parsing response body. Handles errors and returns the response map. Essential for integrating external services in the ondemanddeployer architecture.                      |
| [utils.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/utils/utils.go)   | Manages file read/write operations and caching with synchronized locks. Implements file operations and cache handling using Beego cache client, ensuring data consistency and persistence in the ondemanddeployer architecture. |

</details>

<details closed><summary>scripts</summary>

| File                                                                                               | Summary                                                                                                                                                                                                       |
| ---                                                                                                | ---                                                                                                                                                                                                           |
| [scripts.sh](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/scripts/scripts.sh) | Manages project setup, building, and deployment. Updates environment URL, clones/pulls Git repositories, copies files from S3, builds Docker projects, and starts services. Handles cleaning up all projects. |

</details>

<details closed><summary>components.github</summary>

| File                                                                                                                                   | Summary                                                                                                                                                                                                                                     |
| ---                                                                                                                                    | ---                                                                                                                                                                                                                                         |
| [github_graphql_types.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/components/github/github_graphql_types.go) | Defines GraphQL types for fetching users GitHub statistics—repositories contributed to, pinned items, issues, and pull requests—through the GitHub API. Centralizes data retrieval for user analytics in the ondemanddeployer architecture. |
| [github.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/components/github/github.go)                             | Fetches GitHub repository and language data, ensuring freshness through caching. Utilizes GitHub API and GraphQL to provide up-to-date statistics and metadata. Maintains active repository status.                                         |

</details>

<details closed><summary>components.aws</summary>

| File                                                                                              | Summary                                                                                                                                                                                                                |
| ---                                                                                               | ---                                                                                                                                                                                                                    |
| [sns.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/components/aws/sns.go) | Handles AWS SNS requests, confirms subscriptions, and processes incoming notifications. Publishes messages and executes corresponding actions based on message type. Integrated with AWS SDK for seamless interaction. |

</details>

<details closed><summary>components.bashscript</summary>

| File                                                                                                                   | Summary                                                                                                                                                                                        |
| ---                                                                                                                    | ---                                                                                                                                                                                            |
| [bashscript.go](https://github.com/sharanreddy99/ondemanddeployer.git/blob/master/components/bashscript/bashscript.go) | Manages a queue of Bash script tasks to execute in sequence. Logs task execution, creates test files, and redirects script output. Utilizes a time-based ticker to continuously process tasks. |

</details>

---

##  Getting Started

###  Installation

<h4>From <code>source</code></h4>

> 1. Clone the ondemanddeployer repository:
>
> ```console
> $ git clone https://github.com/sharanreddy99/ondemanddeployer.git
> ```
>
> 2. Change to the project directory:
> ```console
> $ cd ondemanddeployer
> ```
>
> 3. Run the appliacation using
> ```console
> $ go run main.go
> ```

###  Usage


> Access the api at localhost:9452

---


##  Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Report Issues](https://github.com/sharanreddy99/ondemanddeployer.git/issues)**: Submit bugs found or log feature requests for the `ondemanddeployer` project.
- **[Submit Pull Requests](https://github.com/sharanreddy99/ondemanddeployer.git/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/sharanreddy99/ondemanddeployer.git/discussions)**: Share your insights, provide feedback, or ask questions.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/sharanreddy99/ondemanddeployer.git
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="center">
   <a href="https://github.com/sharanreddy99/ondemanddeployer.git/graphs/contributors">
      <img src="https://contrib.rocks/image?repo=sharanreddy99/ondemanddeployer">
   </a>
</p>
</details>

---
