# terraform-vote-provider

## Task: [Extending Terraform]

1. How Terraform Works
    > Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently.

    - **Terraform Core**
    
        The primary responsibilities of Terraform Core are:
    
        - Infrastructure as code: reading and interpolating configuration files and modules
        - Resource state management
        - Construction of the Resource Graph
        - Plan execution
        - Communication with plugins over RPC
    
    - **Terraform Plugins**
        - Provider, such as AWS
        - Provisioner, such as bash

        The primary responsibilities of Provider Plugins are:
            
        - Initialization of any included libraries used to make API calls.
        - Authentication with the infrastructure Provider
        - Define Resources that map to specific Services

        The primary responsibilites of Provisioner Plugins are:

        - Executing commands or scripts on the designated Resource after creation, or on destruction.

1. Terraform Plugin Types

- Providers

    >Providers define **Resources** and are responsible for managing their life cycles.

    Each Resource implements `CREATE`, `READ`, `UPDATE`, and `DELETE` (CRUD) methods to manage itself, while Terraform Core manages a **Resouces Graph** of all the resources declared in the configuration as well as their current state.

    Providers can be fetched from:

    - [releases.hashicorp.com](https://releases.hashicorp.com/)
    - [Provider Installation](https://www.terraform.io/docs/cli/config/config-file.html#provider-installation)

        1. Explicit Installation Method Configuration
        2. Implied Local Mirror Directories

## Develop Customized/In-house Providers

Terraform support developers to explore their customized or in-house providers on their local machine, referring to office doc: [In-house Provider] and [Implied Local Mirror Directories]

For simlicity, run `make install` in the project root path.

More details can be found in the `install` target of [Makefile](./Makefile)


[In-house Provider]:https://www.terraform.io/docs/language/providers/requirements.html#in-house-providers
[Implied Local Mirror Directories]:https://www.terraform.io/docs/cli/config/config-file.html#implied-local-mirror-directories


[Extending Terraform]:https://www.terraform.io/docs/extend/index.html