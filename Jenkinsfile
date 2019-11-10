pipeline {
    agent {
        docker {
            image 'golang:1.13.4-alpine'
            args '-p 3030:3030'
        }
    }
    stages { 
        stage('Build') {
            steps {
                sh 'go mod download'
            }
        }
        stage('Test') {
            steps {
                sh 'go test'
            }
        }
    }
}