## How to Make an AWS 3-Tiered Web App

1. Set up your database on AWS:

* Go to AWS
* Services
* Database
* RDS
* Create Database
* (Engine Types) MySQL
* Free Tier
* (Settings | DB instance identifier) "mydb"
* (Settings | Master Username) "awsuser"
* (Settings | Master Password) "mypassword"
* (Storage autoscaling | Allocated Storage) 5
* (Storage autoscaling) Make sure Enable storage autoscaling is unchecked.
* (Connectivity | Additional connectivity configuration | Public Accessibility) "Yes"
* (Connectivity | Additional connectivity configuration | VPC Security Group) "Choose Existing"
* (Additional Configuration | Backup) Make sure Enable Automatic Backups is unchecked.
* Create Database

2. Connect MySQL Workbench to MySQL Server on AWS (this will no longer work after step 15):

* Go to AWS
* Services
* Database
* RDS
* Databases
* mydbinstance
* (Connectivity & Security | Endpoint & Port) Copy Endpoint.
* Open MySQL Workbench
* Click plus sign next to MySQL Connections.
* (Connection Name) "Test01"
* (Connection Type) "Standard (TCP/IP)"
* (Hostname) paste endpoint
* (Username) "awsuser"
* (Password | Store in keychain ...) "mypassword"
* Test Connection
* OK

3. Create security groups for different pieces of 3-tiered app:

a. First security group is loadbalancer-sg

* Type: HTTP, Protocol: TCP, Port: 80, Source: Anywhere

b. Next security group is webtier-sg

* Type: SSH, Protocol: TCP,Port: 22, Source: Anywhere
* Type: HTTP, Protocol: TCP, Port: 80, Source: <loadbalancer-sg group ID>
* Type: MYSQL/Aurora, Protocol: TCP, Port: 3306, Source: <webtier-sg group ID>

4. Create load balancer:

* Go to AWS
* Services
* EC2
* Load Balancers
* Create Load Balancer
* (Name) "standard-lb"
* Add two subnets.
* Next (Configure Security Settings)
* Next (ignore the warning)
* Select loadbalancer-sg from previous instruction
* Next
* (Target group) New target group
* (Name) standard-tg
* (Path) /ping
* Next
* Next
* Create

5. Modify AWS MySQL RDB security group

* Go to AWS
* Services
* Database
* RDS
* Databases
* Select "mydb"
* Click Modify
* (Network & Security | Security Group) Add webtier-sg and delete any others.
* Continue
* Modify DB Instance

6. Create EC2 instance to add to target group

* Go to AWS
* Services
* EC2
* Instances
* Launch Instance
* Select Ubuntu Server
* Next (until Step 6)
* Select an existing security group (select webtier)
* Review and Launch
* Launch
* Create a new key pair
* (Key pair name) "kp-golang-webdev"
* Download Key Pair
* Launch Instance

7. Move key pair to ssh folder.

* Open terminal.
* $ mv ~/Downloads/kp-golang.webdev.pem ~/.ssh/
* $ chmod 400 ~/.ssh/kp-golang-webdev.pem

8. Make Amazon Machine Image

* Go to AWS
* Services
* EC2
* Instances
* Right click on your instance.
* Image
* Create Image
* (Image name) ami-golang-webdev
* Create Image

9. Create a second EC2 instance based on AMI

* Go to AWS
* Services
* EC2
* Instances
* Launch instance
* My AMIs
* Select ami-golang-webdev
* Next (until step 6)
* Select an existing security group (select webtier)
* Review and Launch
* Launch
* Choose an existing key pair
* Select kp-golang-webdev
* Launch Instance

10. Register instances under target groups.

* Go to AWS
* Services
* EC2
* Target groups
* Click standard-tg
* Targets
* Register Targets
* Select both.
* Include as pending below
* Register pending Targets

11. Compile golang file into binary.

* Navigate to this folder in the Terminal.
* $ GOOS=linux GOARCH=amd64 go build -o mybinary

12. Secure copy binary to EC2 instance (both EC2 instances).

* scp -i ~/.ssh/kp-golang-webdev.pem mybinary ubuntu@<EC2 instance public DNS, like ec2-54-227-222-133.compute-1.amazonaws.com>:

13. SSH to the instance

* ssh -i ~/.ssh/kp-golang-webdev.pem ubuntu@<EC2 instance public DNS, like ec2-54-227-222-133.compute-1.amazonaws.com>

14. Run binary on instance.

* sudo chmod 777 mybinary
* sudo ./mybinary

15. Make it persistent (w/in instance).

* $ cd /etc/systemd/system/
* $ sudo nano first.service

[Unit]
Description=Go Server

[Service]
ExecStart=/home/ubuntu/mybinary
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target

* $ sudo systemctl enable first.service
* $ sudo systemctl start first.service
* $ sudo systemctl status first.service
* $ sudo systemctl stop first.service

15. Connect to MySQL RDB over SSH.

* Turn of public accessibility to DB.
* Open MySQL Workbench
* Click plus sign next to MySQL Connections.
* (Connection Name) "Test01"
* (Connection Type) "Standard TCP/IP Over SSH"
* (SSH Hostname) public DNS of EC2 instance
* (SSH Username) "ubuntu"
* (SSH Keyfile) Find it in files.
* (MYSQL Hostname) paste endpoint
* (MYSQL Username) "awsuser"
* (MYSQL Password | Store in keychain ...) "password"
* Test Connection
* OK


## Tips:

- MAKE SURE YOUR LOAD BALANCER SUPPORTS THE REGIONS YOUR EC2 INSTANCES ARE IN.

- If you need to make changes, use the following commands, then continue as normal.

$ sudo systemctl stop first.service

$ sudo systemctl daemon-reload

- $ sudo systemctl status first.service will show all prints (so it can act as a log file).

### Navigate to public DNS of load balancer to view!!
