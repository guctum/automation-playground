# Security

In the EC2 instance I use the key_name field with the name of an existing key/pair I have in my AWS account.
To generate one, there is a resource that can be provisioned and retrieved via the awscli

# Welcome to your CDK Python project!

This is a blank project for Python development with CDK.

The `cdk.json` file tells the CDK Toolkit how to execute your app.

This project is set up like a standard Python project.  The initialization
process also creates a virtualenv within this project, stored under the `.venv`
directory.  To create the virtualenv it assumes that there is a `python3`
(or `python` for Windows) executable in your path with access to the `venv`
package. If for any reason the automatic creation of the virtualenv fails,
you can create the virtualenv manually.

To manually create a virtualenv on MacOS and Linux:

```
$ python3 -m venv .venv
```

After the init process completes and the virtualenv is created, you can use the following
step to activate your virtualenv.

```
$ source .venv/bin/activate
```

If you are a Windows platform, you would activate the virtualenv like this:

```
% .venv\Scripts\activate.bat
```

Once the virtualenv is activated, you can install the required dependencies.

```
$ pip install -r requirements.txt
```

At this point you can now synthesize the CloudFormation template for this code.

```
$ cdk synth
```

To add additional dependencies, for example other CDK libraries, just add
them to your `setup.py` file and rerun the `pip install -r requirements.txt`
command.

## Useful commands

 * `cdk ls`          list all stacks in the app
 * `cdk synth`       emits the synthesized CloudFormation template
 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk docs`        open CDK documentation

Enjoy!

# Important Things To Take Note Of
* The keypair generated will have to be retrieved via AWS CLI
    - With EC2 keypairs, when they are generated you can download them in the console but you only get the one chance to do so
    - But it can still be obtained via CLI, otherwise it can be created before using CDK and the name just plugged into the code that creates the EC2 instance
    - Unless changed, script will create an EC2 instance of the type t2.micro which is small with limited resources but is free to use
    - System needs to be setup to communicate with AWS
        - This can either take the form of adding AWS Account ID to app.py or using aws configure on terminal to create aws creds
