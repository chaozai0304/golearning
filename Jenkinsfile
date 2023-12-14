pipeline {
  agent any
  stages {
    stage('one') {
      steps {
        sh 'ls'
        sh '''ls -l
'''
      }
    }

    stage('two') {
      steps {
        sh 'uname -a'
      }
    }

  }
}