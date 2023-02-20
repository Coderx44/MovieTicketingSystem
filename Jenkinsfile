pipeline {
    agent any
    tools {
        go 'Go'
        dockerTool 'docker'
    }
    
    environment {
        def DATETIME = sh(script: 'date "+%Y-%m-%d_%H-%M-%S"', returnStdout: true).trim()
        
    }
    stages {
        stage('Fetch Code from Github') {
            steps {
                checkout scmGit(branches: [[name: '*/unit-tests']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/Coderx44/MovieTicketingSystem/']])
                sh 'go mod tidy'
                sh 'go build'
            }
        }
        stage('Build Docker Image'){
            steps{
                script{
                    sh 'docker build -t movieticket-go .'
                }
            }
        }
        
        stage('Login & Push to AWS-ECR'){
             steps{
                 script{
                     withCredentials([[$class: 'AmazonWebServicesCredentialsBinding',accessKeyVariable: 'AWS_ACCESS_KEY_ID',secretKeyVariable: 'AWS_SECRET_ACCESS_KEY',credentialsId: 'aws-cred']]) {
                        sh 'aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin 637069043585.dkr.ecr.ap-south-1.amazonaws.com/movieticket-go'
    
                        sh 'docker tag movieticket-go:latest 637069043585.dkr.ecr.ap-south-1.amazonaws.com/movieticket-go:$DATETIME'
    
                        sh 'docker push 637069043585.dkr.ecr.ap-south-1.amazonaws.com/movieticket-go:$DATETIME'
                    }
                 }
             }
         }
         stage('Pull & Run Image'){
             agent { node { label 'ec2-node'}}
             steps{
                 withCredentials([[$class: 'AmazonWebServicesCredentialsBinding',accessKeyVariable: 'AWS_ACCESS_KEY_ID',secretKeyVariable: 'AWS_SECRET_ACCESS_KEY',credentialsId: 'aws-cred']]) {
                    sh 'aws ecr get-login-password --region ap-south-1 | docker login --username AWS --password-stdin 637069043585.dkr.ecr.ap-south-1.amazonaws.com/movieticket-go'
    
                    sh 'docker pull 637069043585.dkr.ecr.ap-south-1.amazonaws.com/movieticket-go:$DATETIME'
                    sh """docker stop deploy || true"""
                    sh 'docker run --name=deploy --rm -d -v /home/ubuntu/secrets:/akshat -p 80:3000 637069043585.dkr.ecr.ap-south-1.amazonaws.com/movieticket-go:$DATETIME'
                    sh 'docker ps'
                    }
                }
            }
    }
    post {

        always {

            mail bcc: '', body: " Hi Team \n I have forwarded the build status of $JOB_NAME  \n Build : $BUILD_NUMBER  ${currentBuild.currentResult}.\n \n \n Check the console output at ${env.BUILD_URL} to view results\n Thanks and Regards ", cc: 'akshat.dubey@joshsoftware.com', from: '', replyTo: '', subject: 'Test Email From Jenkins Job', to: 'dubeyakshat88@joshsoftware.com'

        }

    }
    
}
