
Go Utility for Creating IAC Modules

This utility is a Golang CLI tool that creates the scaffolding for a IAC modules. It can be used to create modules for Infrastructure as Code (IAC) tools.


Features

Creates a basic module structure, including files for the module's inputs, outputs, resources, and providers.
    
    
Generates documentation for the module, including a description, usage examples, and input and output descriptions.


```
brew install  p0bailey/homebrew-tools/tofuinit
```


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
    