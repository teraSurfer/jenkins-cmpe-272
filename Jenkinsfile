pipeline {
    agent {
        docker {
            image 'golang:1.13-alpine'
            args '-p 3030:3030'
        }
    }
  stages {
    stage('Build') {
      steps {
        sh 'apk add libc-dev'
        sh 'go mod download'
      }
    }

    stage('Test') {
      steps {
        sh ' CGO_ENABLED=0 go test'
      }
    }

  }
}