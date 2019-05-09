# go-ssm
A single binary for fetching AWS SSM parameters

# Install

```
glide install
```

# Build

```bash
make build
```

# Usage

First, make sure your AWS region is already set:

```bash
export AWS_REGION=us-west-2
```

Then export the list of parameters that you need:

```bash
export SSM_KEYS=/PATH/KEY_1,/PATH/KEY_2,/PATH/KEY_3
```

Now you can fetch the parameters and pass them to the next command:

```bash
./bin/go_ssm node main.js
```
(In this example, all the parameters will be available for the node.js app in `process.env`)