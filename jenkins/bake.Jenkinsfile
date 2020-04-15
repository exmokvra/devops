pipeline {
    agent any

    stages {
        stage('Build') {
            environment {
                AWS_ACCESS_KEY = credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_KEY = credentials('AWS_SECRET_KEY')
            }

            steps {
                sh 'rm -rf packer'
                sh 'rm -rf devops'
                sh 'rm -rf packer_1.5.5_linux_amd64.zip'
                
                sh 'git clone https://github.com/guisesterheim/devops/'

                sh 'wget -c https://releases.hashicorp.com/packer/1.5.5/packer_1.5.5_linux_amd64.zip'
                sh 'unzip -o packer_1.5.5_linux_amd64.zip'

                sh 'echo printing pwd'
                sh 'pwd'
                sh 'ls -lsa'
                sh 'cat devops/jenkins/aws-template.json'

                sh './packer build -var \'aws_access_key=$AWS_ACCESS_KEY_PSW\' -var \'aws_secret_key=$AWS_SECRET_KEY_PSW\' devops/jenkins/aws-template.json'

                echo 'Done!'
            }
        }
    }
}