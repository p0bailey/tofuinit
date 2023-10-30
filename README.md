#tofuinit

Go Utility for Generating Fully Compliant IaC Modules Scaffolding Code

I created this utility because I was frustrated with having to manually set up the module structure each time I needed to write a new Infrastructure as Code (IaC) module.

This Golang CLI tool automatically generates the scaffolding for fully compliant IaC modules. It can be used to generate a new IAC module from scratch or to add new components to an existing module.


To use tofuinit, simply run the following command:

```
tofuinit <module_name>
```

Features

Creates a basic module structure, including files for the module's inputs, outputs, resources, and providers.
    
    
# Installation

```
brew install  p0bailey/homebrew-tools/tofuinit
brew install tofuinit
```


# Usage


```
Usage: tofuinit <module_name>

tofuinit is a tool for initializing a new IaC module with the specified name.

Options:
  --help     Display this help menu and exit

Arguments:
  module_name    The name of the module to be created
```

```
mymodule
├── README.md
├── main.tf
├── outputs.tf
├── variables.tf
├── versions.tf
├── docs
│   └── README.md
├── examples
│   └── README.md
└── tests
    └── main.tftest.hcl
```    
    
    
#Contributions 
    
I accept contributions and pull requests! I am always looking for ways to improve, and I am grateful for the help of the community.

If you have a contribution or pull request, please feel free to open one on GitHub. I will review your contribution or pull request as soon as possible, and I will merge it if it meets the following criteria:

The contribution or pull request is well-written and easy to understand.
The contribution or pull request adds value to the project.
The contribution or pull request is well-tested.

I look forward to your contributions!    