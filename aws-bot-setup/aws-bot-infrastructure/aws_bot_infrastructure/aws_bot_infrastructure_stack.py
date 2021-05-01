from aws_cdk import core as cdk

# For consistency with other languages, `cdk` is the preferred import name for
# the CDK's core module.  The following line also imports it as `core` for use
# with examples from the CDK Developer's Guide, which are in the process of
# being updated to use `cdk`.  You may delete this import if you don't need it.
from aws_cdk import (
    aws_ec2 as ec2,
    aws_iam as iam,
    aws_s3 as s3,
    core
)
from cdk_ec2_key_pair import KeyPair


class AwsBotInfrastructureStack(cdk.Stack):

    def __init__(self, scope: cdk.Construct, construct_id: str, **kwargs) -> None:
        super().__init__(scope, construct_id, **kwargs)

        # The code that defines your stack goes here

        bucket = s3.Bucket(self, "AWS Bot Bucket",
                           versioned=True,
                           removal_policy=core.RemovalPolicy.DESTROY,
                           auto_delete_objects=True)

        # Key Pair
        key = KeyPair(self, "AWS Sample EC2 KeyPair",
                      name="Sandbox Key",
                      description="Generic key for AWS Sandbox - or any purpose really",
                      store_public_key=True)

        # VPC - open port 22 somewhere to allow for SSH, should be done from the subgroup
        vpc = ec2.Vpc(self, "VPC",
                      nat_gateways=0,
                      subnet_configuration=[ec2.SubnetConfiguration(name="public",subnet_type=ec2.SubnetType.PUBLIC)]
                      )

        # AMI
        amzn_linux = ec2.MachineImage.latest_amazon_linux(
            generation=ec2.AmazonLinuxGeneration.AMAZON_LINUX_2,
            edition=ec2.AmazonLinuxEdition.STANDARD,
            virtualization=ec2.AmazonLinuxVirt.HVM,
            storage=ec2.AmazonLinuxStorage.GENERAL_PURPOSE
        )

        # Instance Role and SSM Managed Policy
        role = iam.Role(self, "InstanceSSM", assumed_by=iam.ServicePrincipal("ec2.amazonaws.com"))

        role.add_managed_policy(iam.ManagedPolicy.from_aws_managed_policy_name("service-role/AmazonEC2RoleforSSM"))

        # Instance
        instance = ec2.Instance(self, "Instance",
                                instance_type=ec2.InstanceType("t2.micro"), # t2.micro is free tier, bigger can be used if needed
                                machine_image=amzn_linux,
                                vpc = vpc,
                                role = role,
                                key_name = key.name # EC2 instance will use the keypair we created earlier, it can be grabbed using the CLI
                                # key_name = "discord-bot" # existing key that exists in my account
                                )
