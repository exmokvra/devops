pipeline {
    agent any

    stages {
        stage('Build') {
            environment {
                AWS_ACCESS_KEY = credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_KEY = credentials('AWS_SECRET_KEY')
            }

            steps {
                echo 'Cloning... https://github.com/guisesterheim/devops/'
                sh 'rm -rf devops/'
                sh 'git clone https://github.com/guisesterheim/devops/'

                echo 'Printing credentials...'
                sh 'echo Access key $AWS_ACCESS_KEY_PSW'
                sh 'echo Secret key $AWS_SECRET_KEY_PSW'

                echo 'Running packer build to AWS...'
                sh 'packer build -var \"aws_access_key=$AWS_ACCESS_KEY_PSW\" -var \"aws_secret_key=$AWS_SECRET_KEY_PSW\" aws-template.json\''

                echo 'Done!'
            }
        }
    }
}