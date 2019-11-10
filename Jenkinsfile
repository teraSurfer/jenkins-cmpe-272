ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
   withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
   }
}
stage('Pre Test'){
    echo 'Pulling Dependencies'

    sh 'go version'
    sh 'cd $GOPATH/github.com/terasurfer/jenkins-cmpe-272 && go mod download'
}

stage('Test'){
     echo 'Testing'

     sh """cd $GOPATH && go test"""
}