pipeline {
    agent {docker 'public.ecr.aws/docker/library/golang:latest'}
    environment { GOCACHE = "/tmp" }
    stages {
        stage('Source') {
            steps {
                sh 'go version'
                sh 'git --version'
                git branch: 'main',
                    url: 'https://github.com/acorallo/test-golang-ci.git'
            }
        }
        stage('Clean') {
            steps {
                    sh 'pwd'
                    sh 'go run main.go'
                }
            }
        }
    }
}
