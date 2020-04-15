pipeline {
    agent any

    stages {
        stage('Build') {
            environment {
                AWS_ACCESS_KEY = credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_KEY = credentials('AWS_SECRET_KEY')
            }

            steps {
                // Some updates and cleaning old files
                sh 'rm -rf packer'
                sh 'rm -rf devops'
                sh 'rm -rf packer_1.5.5_linux_amd64.zip'
                
                // Get the packer json build file
                sh 'git clone https://github.com/guisesterheim/devops/'

                // Install Packer
                sh 'wget -c https://releases.hashicorp.com/packer/1.5.5/packer_1.5.5_linux_amd64.zip'
                sh 'unzip -o packer_1.5.5_linux_amd64.zip'

                // Packer Build
                sh './packer build -machine-readable -var aws_access_key=$AWS_ACCESS_KEY_PSW -var aws_secret_key=$AWS_SECRET_KEY_PSW packer_content/aws-template.json | tee build.log'

                // Get generated AMI name
                sh 'grep \'artifact,0,id\' build.log cut -d, -f6 | cut -d: -f2' // More at >> https://gist.github.com/danrigsby/11354917

                // Store AMI image name on variable
                // TODO:
            }
        }
        stage('Launch'){
            // Launch a new machine with the AMI generated on the first step
        }
    }
}