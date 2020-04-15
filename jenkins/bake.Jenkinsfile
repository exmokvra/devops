pipeline {
    agent any

    stages {
        stage('Build') {
            environment {
                AWS_ACCESS_KEY = credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_KEY = credentials('AWS_SECRET_KEY')
            }

            steps {
                sh 'rm -rf devops/'
                sh 'git clone https://github.com/guisesterheim/devops/'

                sh 'wget -c https://releases.hashicorp.com/packer/1.5.5/packer_1.5.5_linux_amd64.zip'
                sh 'unzip -o packer_1.5.5_linux_amd64.zip.zip'

                sh 'echo Printing packer version'
                sh './packer build -machine-readable -var \'aws_access_key=$AWS_ACCESS_KEY_PSW\' -var \'aws_secret_key=$AWS_SECRET_KEY_PSW\' devops/jenkins/aws-template.json'

                echo 'Done!'
            }
        }
    }
}